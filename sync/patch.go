package sync

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/math"
	viperUtil "github.com/Conflux-Chain/go-conflux-util/viper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrNoFileInfoToPatch      = errors.New("No file info to patch")
	ErrNoMinerAttemptsToPatch = errors.New("No miner attempts to patch")
	ErrMinPosNotFinalized     = errors.New("MinPos not finalized")
	ErrFinalizedPosNotSynced  = errors.New("FinalizedPos not synced")
	BatchSubmitsOnPatch       = 100
	BatchBlocksOnPatch        = 100
)

type PatchSyncer struct {
	conf            *SyncConfig
	sdk             *web3go.Client
	db              *store.MysqlStore
	mineAddr        string
	minerSubmitSig  string
	currentSubmitId uint64
	currentBn       uint64
}

func MustNewPatchSyncer(sdk *web3go.Client, db *store.MysqlStore, conf SyncConfig) *PatchSyncer {
	syncer := PatchSyncer{
		conf: &conf,
		sdk:  sdk,
		db:   db,
	}

	syncer.mustInit()

	return &syncer
}

func (ps *PatchSyncer) mustInit() {
	ps.mustInitContract()
	ps.mustLoadLastSubmitId()
	ps.mustLoadLastBn()
}

func (ps *PatchSyncer) mustInitContract() {
	var cfg contractConfig
	viperUtil.MustUnmarshalKey("contract", &cfg)

	if cfg.Flow.Address == "" {
		logrus.Fatal("No flow contract address configured")
	}
	ps.mustInitAddresses(cfg.Flow.Address)

	ps.minerSubmitSig = cfg.Mine.SubmitMethodSignature
}

func (ps *PatchSyncer) mustInitAddresses(flowAddr string) {
	_, mineAddr := MustGetContractAddresses(ps.sdk, flowAddr)
	ps.mineAddr = mineAddr
}

func (ps *PatchSyncer) mustLoadLastSubmitId() {
	loaded, err := ps.loadLastSubmitId()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load submit id from db")
	}

	// Submit id is set 0 value if not loaded.
	if !loaded {
		ps.currentSubmitId = 0
	}
}

func (ps *PatchSyncer) loadLastSubmitId() (loaded bool, err error) {
	submitId, ok, err := ps.db.ConfigStore.GetUint64(store.SyncPatchL1TxSubmitId)
	if err != nil {
		return false, errors.WithMessagef(err, "Failed to get submit id")
	}

	if ok {
		ps.currentSubmitId = submitId + 1
	}

	return ok, nil
}

func (ps *PatchSyncer) mustLoadLastBn() {
	loaded, err := ps.loadLastBn()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load bn from db")
	}

	// Block number is set config value if not loaded.
	if !loaded {
		ps.currentBn = ps.conf.BlockWhenFlowCreated
	}
}

func (ps *PatchSyncer) loadLastBn() (loaded bool, err error) {
	bn, ok, err := ps.db.ConfigStore.GetUint64(store.SyncPatchMinerAttemptsBn)
	if err != nil {
		return false, errors.WithMessagef(err, "Fail to get bn")
	}

	if ok {
		ps.currentBn = bn + 1
	}

	return ok, nil
}

func (ps *PatchSyncer) Sync(ctx context.Context, f func(ctx context.Context, ticker *time.Ticker)) {
	ticker := time.NewTicker(intervalNormal)
	defer ticker.Stop()

	logrus.Info("Patch syncer starting to sync data.")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			f(ctx, ticker)
		}
	}
}

func (ps *PatchSyncer) L1Txs(ctx context.Context, ticker *time.Ticker) {
	if err := ps.syncL1Txs(ctx); err != nil {
		if !errors.Is(err, ErrNoFileInfoToPatch) {
			ticker.Reset(intervalException)
			logrus.WithError(err).Error("Failed to patch l1 tx info")
			return
		}
	}
	ticker.Reset(intervalNormal)
}

func (ps *PatchSyncer) syncL1Txs(ctx context.Context) error {
	submits, err := ps.db.SubmitStore.QueryAscWithCursor(&ps.currentSubmitId, BatchSubmitsOnPatch)
	if err != nil {
		return err
	}
	if len(submits) == 0 {
		return ErrNoFileInfoToPatch
	}

	for _, submit := range submits {
		if interrupted(ctx) {
			return nil
		}

		hash := common.HexToHash(submit.TxHash)
		tx, err := ps.sdk.Eth.TransactionByHash(hash)
		if err != nil {
			return errors.WithMessage(err, "Failed to get tx")
		}
		rcpt, err := ps.sdk.Eth.TransactionReceipt(hash)
		if err != nil {
			return errors.WithMessage(err, "Failed to get rcpt")
		}
		if tx == nil || rcpt == nil { // Not patch pruned tx
			continue
		}

		var extra store.SubmitExtra
		if err = json.Unmarshal(submit.Extra, &extra); err != nil {
			return errors.WithMessage(err, "Failed to unmarshal extra")
		}

		extra.GasPrice = tx.GasPrice.Uint64()
		extra.GasLimit = tx.Gas
		extra.GasUsed = rcpt.GasUsed
		updateExtra, err := json.Marshal(extra)
		if err != nil {
			return errors.WithMessage(err, "Failed to marshal extra")
		}

		update := store.Submit{
			SubmissionIndex: submit.SubmissionIndex,
			Extra:           updateExtra,
		}
		if err := ps.db.SubmitStore.UpdateByPrimaryKey(nil, &update); err != nil {
			return err
		}

		if err := ps.db.ConfigStore.Upsert(nil, store.SyncPatchL1TxSubmitId,
			strconv.FormatUint(submit.SubmissionIndex, 10)); err != nil {
			return err
		}
	}

	lastSubmitId := submits[len(submits)-1].SubmissionIndex
	ps.currentSubmitId = lastSubmitId + 1
	return nil
}

func (ps *PatchSyncer) MinerAttempts(ctx context.Context, ticker *time.Ticker) {
	if err := ps.syncMinerAttempts(ctx); err != nil {
		if !errors.Is(err, ErrNoMinerAttemptsToPatch) &&
			!errors.Is(err, ErrFinalizedPosNotSynced) &&
			!errors.Is(err, ErrMinPosNotFinalized) {
			ticker.Reset(intervalException)
			logrus.WithError(err).Error("Failed to patch miner attempts")
			return
		}
	}
	ticker.Reset(intervalNormal)
}

func (ps *PatchSyncer) syncMinerAttempts(ctx context.Context) error {
	minPos, maxPos, err := ps.nextStatRange()
	if err != nil {
		return err
	}

	var bns []types.BlockNumber
	for bn := minPos; bn <= maxPos; bn++ {
		bns = append(bns, types.BlockNumber(bn))
	}
	bn2blk, err := rpc.BatchGetBlocks(ctx, ps.sdk, bns, true, ps.conf.BatchBlocksOnBatchCall)
	if err != nil {
		return err
	}

	miner2Attempts := make(map[uint64]uint64)
	for _, blk := range bn2blk {
		for _, tx := range blk.Transactions.Transactions() {
			input := tx.Input.String()
			if tx.To != nil && tx.To.String() == ps.mineAddr &&
				len(input) >= 266 && tx.Input.String()[0:10] == ps.minerSubmitSig {
				minerID := "0x" + input[202:266]
				register, ok, err := ps.db.GetByMinerID(minerID, blk.Number.Uint64())
				if err != nil {
					return err
				}
				if !ok {
					return errors.Errorf("Register miner id not found %s", minerID)
				}
				miner2Attempts[register.AddressID]++
			}
		}
	}

	minersUpdate := make([]store.Miner, 0)
	for minerID, attempts := range miner2Attempts {
		minersUpdate = append(minersUpdate, store.Miner{
			ID:             minerID,
			MiningAttempts: attempts,
		})
	}

	lastBn := uint64(bns[len(bns)-1].Int64())
	if err := ps.db.DB.Transaction(func(dbTx *gorm.DB) error {
		if len(minersUpdate) > 0 {
			err := ps.db.MinerStore.BatchDeltaUpsertMiningAttempts(dbTx, minersUpdate)
			if err != nil {
				return err
			}
		}
		if err := ps.db.ConfigStore.Upsert(nil, store.SyncPatchMinerAttemptsBn,
			strconv.FormatUint(lastBn, 10)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	ps.currentBn = lastBn + 1

	return nil
}

func (ps *PatchSyncer) nextStatRange() (uint64, uint64, error) {
	minPos := ps.currentBn
	maxPos := ps.currentBn + uint64(BatchBlocksOnPatch) - 1

	block, err := rpc.GetBlockByNumber(ps.sdk, types.FinalizedBlockNumber, false)
	if err != nil {
		return 0, 0, err
	}

	finalizedPos, exists, err := ps.db.BlockStore.MaxBlockFinalized(block.Number.Uint64())
	if err != nil {
		return 0, 0, err
	}
	if !exists {
		return 0, 0, ErrFinalizedPosNotSynced
	}
	if finalizedPos < minPos {
		return 0, 0, ErrMinPosNotFinalized
	}

	maxPos = math.MinUint64(maxPos, finalizedPos)

	return minPos, maxPos, nil
}
