package sync

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/0glabs/0g-storage-scan/rpc"
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/health"
	"github.com/openweb3/web3go"
	"github.com/sirupsen/logrus"
)

var (
	BatchGetSubmitsLatest = 100
	intervalNormal        = time.Second
	intervalException     = time.Second * 10
)

type StorageSyncer struct {
	db               *store.MysqlStore
	storageConfig    rpc.StorageConfig
	alertChannel     string
	healthReport     health.TimedCounterConfig
	storageRpcHealth health.TimedCounter
	blockchainClient *web3go.Client
}

func MustNewStorageSyncer(db *store.MysqlStore, storageConfig rpc.StorageConfig, alertChannel string,
	healthReport health.TimedCounterConfig, blockchainClient *web3go.Client) *StorageSyncer {
	return &StorageSyncer{
		db:               db,
		storageConfig:    storageConfig,
		alertChannel:     alertChannel,
		healthReport:     healthReport,
		storageRpcHealth: health.TimedCounter{},
		blockchainClient: blockchainClient,
	}
}

func (ss *StorageSyncer) Sync(ctx context.Context, f func(ctx context.Context, ticker *time.Ticker)) {
	ticker := time.NewTicker(intervalNormal)
	defer ticker.Stop()

	logrus.Info("Storage syncer starting to sync data.")
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			f(ctx, ticker)
		}
	}
}

func (ss *StorageSyncer) LatestFiles(ctx context.Context, ticker *time.Ticker) {
	if interrupted(ctx) {
		return
	}

	submits, err := ss.db.SubmitStore.QueryDesc(BatchGetSubmitsLatest)
	if err != nil {
		ticker.Reset(intervalException)
		return
	}

	if len(submits) == 0 {
		return
	}

	unfinalized := make([]store.Submit, 0)
	for _, submit := range submits {
		if submit.Status < uint8(rpc.Pruned) {
			unfinalized = append(unfinalized, submit)
		}
	}

	if len(unfinalized) == 0 {
		return
	}

	if _, err := ss.db.UpdateFileInfos(ctx, unfinalized, ss.storageConfig); err != nil {
		ticker.Reset(intervalException)
	}
}

func (ss *StorageSyncer) NodeSyncHeight(ctx context.Context, ticker *time.Ticker) {
	var err error
	nodeStatus, err := rpc.GetNodeStatus(ss.storageConfig)

	if err == nil {
		height := nodeStatus.LogSyncHeight
		err = ss.db.ConfigStore.Upsert(nil, store.SyncHeightNode, strconv.FormatUint(height, 10))
		if err != nil {
			logrus.WithError(err).Error("Failed to upsert storage node sync height")
		} else {
			// Check sync height gaps and use gap error for alerting if present
			err = ss.checkSyncHeightGaps(height)
		}
	}

	if ss.alertChannel != "" {
		e := rpc.AlertErr(ctx, "StorageIndexerRPCError", ss.alertChannel, err, ss.healthReport,
			&ss.storageRpcHealth, ss.storageConfig.Indexer)

		if e != nil {
			ticker.Reset(intervalException)
			logrus.WithError(err).Error("Failed to alert storage node status")
		} else {
			ticker.Reset(intervalNormal)
		}
	}
}

// checkSyncHeightGaps monitors sync height differences and returns error if gaps exceed 1000 blocks
func (ss *StorageSyncer) checkSyncHeightGaps(nodeSyncHeight uint64) error {
	if ss.blockchainClient == nil {
		return nil
	}

	// Get current blockchain height
	currentBlock, err := ss.blockchainClient.Eth.BlockNumber()
	if err != nil {
		logrus.WithError(err).Error("Failed to get current block height for sync monitoring")
		return nil // Don't return blockchain RPC errors as sync gap errors
	}

	currentHeight := currentBlock.Uint64()
	const alertThreshold = 1000

	// Check layer1-logsyncheight gap (node sync height vs blockchain height)
	if currentHeight > nodeSyncHeight && currentHeight-nodeSyncHeight > alertThreshold {
		gap := currentHeight - nodeSyncHeight
		return fmt.Errorf("Layer1LogSyncHeight sync gap: %d blocks behind (sync: %d, current: %d)", gap, nodeSyncHeight, currentHeight)
	}

	// Get scanner sync height from database
	scannerSyncHeight, ok, err := ss.db.BlockStore.MaxBlock()
	if err != nil {
		logrus.WithError(err).Error("Failed to get scanner sync height for monitoring")
		return nil // Don't return DB errors as sync gap errors
	}

	if !ok {
		return nil // No blocks synced yet
	}

	// Check logsyncheight gap (scanner sync height vs blockchain height)
	if currentHeight > scannerSyncHeight && currentHeight-scannerSyncHeight > alertThreshold {
		gap := currentHeight - scannerSyncHeight
		return fmt.Errorf("LogSyncHeight sync gap: %d blocks behind (sync: %d, current: %d)", gap, scannerSyncHeight, currentHeight)
	}

	return nil // No sync gap issues
}
