package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/0glabs/0g-storage-scan/metrics"
	"github.com/Conflux-Chain/go-conflux-util/alert"
	"github.com/Conflux-Chain/go-conflux-util/health"
	set "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/go-rpc-provider"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/types"
	"github.com/pkg/errors"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrChainReorged = errors.New("chain re-orged")
)

type EthData struct {
	Number   uint64
	Block    *types.Block
	Receipts map[common.Hash]*types.Receipt
	Logs     []types.Log
}

func IsTxExecutedInBlock(tx types.TransactionDetail, receipt types.Receipt) bool {
	return tx.BlockHash != nil && receipt.Status != nil && *receipt.Status < 2
}

func GetEthDataByReceipts(w3c *web3go.Client, blockNumber uint64) (*EthData, error) {
	start := time.Now()
	data, err := getEthDataByReceipts(w3c, blockNumber)

	metrics.Registry.Sync.QueryEthData("getBlockReceipts").UpdateSince(start)
	metrics.Registry.Sync.QueryEthDataAvailability("getBlockReceipts").
		Mark(err == nil || errors.Is(err, ErrChainReorged))

	return data, err
}

func getEthDataByReceipts(w3c *web3go.Client, blockNumber uint64) (*EthData, error) {
	// Use batch function for single block
	batchData, err := getEthDataBatchByReceipts(w3c, blockNumber, blockNumber)
	if err != nil {
		return nil, err
	}

	// Return the single block data or nil if no data
	if len(batchData) == 0 {
		return nil, nil
	}

	return batchData[0], nil
}

// GetEthDataBatchByReceipts gets data for multiple blocks using batch receipt retrieval
func GetEthDataBatchByReceipts(w3c *web3go.Client, blockFrom, blockTo uint64) ([]*EthData, error) {
	start := time.Now()
	data, err := getEthDataBatchByReceipts(w3c, blockFrom, blockTo)

	metrics.Registry.Sync.QueryEthData("getBatchReceipts").UpdateSince(start)
	metrics.Registry.Sync.QueryEthDataAvailability("getBatchReceipts").Mark(err == nil || errors.Is(err, ErrChainReorged))

	return data, err
}

func getEthDataBatchByReceipts(w3c *web3go.Client, blockFrom, blockTo uint64) ([]*EthData, error) {
	// get blocks for the range
	var blockNumbers []types.BlockNumber
	for blockNum := blockFrom; blockNum <= blockTo; blockNum++ {
		blockNumbers = append(blockNumbers, types.BlockNumber(blockNum))
	}

	// batch get blocks
	ctx := context.Background()
	blockMap, err := BatchGetBlocks(ctx, w3c, blockNumbers, true, 16)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to batch get blocks %v-%v", blockFrom, blockTo)
	}

	// create EthData for each block by getting receipts individually
	var result []*EthData
	for blockNum := blockFrom; blockNum <= blockTo; blockNum++ {
		block := blockMap[blockNum]

		// batch get receipts for this block
		blockNumOrHash := types.BlockNumberOrHashWithNumber(types.BlockNumber(blockNum))
		blockReceipts, err := w3c.Parity.BlockReceipts(&blockNumOrHash)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to get block receipts for block %v", blockNum)
		}
		if blockReceipts == nil {
			return nil, errors.WithMessagef(ErrChainReorged, "batch retrieved block receipts nil for block %v", blockNum)
		}

		// process receipts with reorg checking
		txReceipts := map[common.Hash]*types.Receipt{}
		blockTxs := block.Transactions.Transactions()
		if len(blockTxs) != len(blockReceipts) {
			return nil, errors.Errorf("block receipts number mismatch at block %v, rcpts %v, txs %v", blockNum, len(blockReceipts), len(blockTxs))
		}

		for i := 0; i < len(blockTxs); i++ {
			tx := blockTxs[i]
			receipt := &blockReceipts[i]

			// check re-org
			switch {
			case receipt.BlockHash != block.Hash:
				return nil, errors.WithMessagef(ErrChainReorged, "receipt block hash mismatch at block %v, rcptBlkHash %v, blkHash %v",
					blockNum, receipt.BlockHash, block.Hash)
			case receipt.BlockNumber != blockNum:
				return nil, errors.WithMessagef(ErrChainReorged, "receipt block num mismatch at block %v, rcptBlkNum %v, blkNum %v",
					blockNum, receipt.BlockNumber, blockNum)
			case receipt.TransactionHash != tx.Hash:
				return nil, errors.WithMessagef(ErrChainReorged, "receipt tx hash mismatch at block %v, rcptTxHash %v, TxHash %v",
					blockNum, receipt.TransactionHash, tx.Hash)
			}
			txReceipts[tx.Hash] = receipt
		}

		ethData := &EthData{
			Number:   blockNum,
			Block:    &block,
			Receipts: txReceipts,
		}
		result = append(result, ethData)
	}

	return result, nil
}

func GetEthDataByLogs(w3c *web3go.Client, blockNumber uint64, addresses []common.Address, topics [][]common.Hash) (*EthData, error) {
	start := time.Now()
	data, err := getEthDataByLogs(w3c, blockNumber, addresses, topics)

	metrics.Registry.Sync.QueryEthData("getLogs").UpdateSince(start)
	metrics.Registry.Sync.QueryEthDataAvailability("getLogs").Mark(err == nil || errors.Is(err, ErrChainReorged))

	return data, err
}

func getEthDataByLogs(w3c *web3go.Client, blockNumber uint64, addresses []common.Address, topics [][]common.Hash) (*EthData, error) {
	// Use batch function for single block
	batchData, err := getEthDataBatchByLogs(w3c, blockNumber, blockNumber, addresses, topics)
	if err != nil {
		return nil, err
	}

	// Return the single block data or nil if no data
	if len(batchData) == 0 {
		return nil, nil
	}

	return batchData[0], nil
}

// GetEthDataBatchByLogs gets data for multiple blocks using batch log retrieval
func GetEthDataBatchByLogs(w3c *web3go.Client, blockFrom, blockTo uint64, addresses []common.Address, topics [][]common.Hash) ([]*EthData, error) {
	start := time.Now()
	data, err := getEthDataBatchByLogs(w3c, blockFrom, blockTo, addresses, topics)

	metrics.Registry.Sync.QueryEthData("getBatchLogs").UpdateSince(start)
	metrics.Registry.Sync.QueryEthDataAvailability("getBatchLogs").Mark(err == nil || errors.Is(err, ErrChainReorged))

	return data, err
}

func getEthDataBatchByLogs(w3c *web3go.Client, blockFrom, blockTo uint64, addresses []common.Address, topics [][]common.Hash) ([]*EthData, error) {
	// batch get logs for the range
	logArray, err := BatchGetLogs(w3c, blockFrom, blockTo, addresses, topics)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get batch logs for blocks %v-%v", blockFrom, blockTo)
	}

	// group logs by block number
	blockLogs := make(map[uint64][]types.Log)
	for _, log := range logArray {
		blockNum := log.BlockNumber
		blockLogs[blockNum] = append(blockLogs[blockNum], log)
	}

	// get blocks for the range
	var blockNumbers []types.BlockNumber
	for blockNum := blockFrom; blockNum <= blockTo; blockNum++ {
		blockNumbers = append(blockNumbers, types.BlockNumber(blockNum))
	}

	// batch get blocks
	ctx := context.Background()
	blockMap, err := BatchGetBlocks(ctx, w3c, blockNumbers, true, 16)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to batch get blocks %v-%v", blockFrom, blockTo)
	}

	// create EthData for each block
	var result []*EthData
	for blockNum := blockFrom; blockNum <= blockTo; blockNum++ {
		block := blockMap[blockNum]

		logs := blockLogs[blockNum]
		if logs == nil {
			logs = []types.Log{} // empty logs for blocks without relevant logs
		}

		// validate logs belong to the correct block and check for reorg
		validLogs := make([]types.Log, 0, len(logs))
		txIndex2Tx := make(map[uint64]types.TransactionDetail)

		// build transaction index map for reorg checking
		txs := block.Transactions.Transactions()
		for _, tx := range txs {
			txIndex2Tx[*tx.TransactionIndex] = tx
		}

		for _, log := range logs {
			// check re-org
			if tx, exists := txIndex2Tx[uint64(log.TxIndex)]; exists {
				switch {
				case log.BlockHash != block.Hash:
					return nil, errors.WithMessagef(ErrChainReorged, "log block hash mismatch at block %v, logBlkHash %v, blkHash %v",
						blockNum, log.BlockHash, block.Hash)
				case log.BlockNumber != blockNum:
					return nil, errors.WithMessagef(ErrChainReorged, "log block num mismatch at block %v, logBlkNum %v, blkNum %v",
						blockNum, log.BlockNumber, blockNum)
				case log.TxHash != tx.Hash:
					return nil, errors.WithMessagef(ErrChainReorged, "log tx hash mismatch at block %v, logTxHash %v, txHash %v",
						blockNum, log.TxHash, tx.Hash)
				case uint64(log.TxIndex) != *tx.TransactionIndex:
					return nil, errors.WithMessagef(ErrChainReorged, "log tx index mismatch at block %v, logTxIndex %v, txIndex %v",
						blockNum, log.TxIndex, tx.TransactionIndex)
				}
				validLogs = append(validLogs, log)
			}
		}

		ethData := &EthData{
			Number: blockNum,
			Block:  &block,
			Logs:   validLogs,
		}
		result = append(result, ethData)
	}

	return result, nil
}

func BatchGetLogs(w3c *web3go.Client, blockFrom, blockTo uint64, addresses []common.Address,
	topics [][]common.Hash) ([]types.Log, error) {
	bnFrom := types.NewBlockNumber(int64(blockFrom))
	bnTo := types.NewBlockNumber(int64(blockTo))
	logFilter := types.FilterQuery{
		FromBlock: &bnFrom,
		ToBlock:   &bnTo,
		Addresses: addresses,
		Topics:    topics,
	}
	return w3c.Eth.Logs(logFilter)
}

func BatchGetBlockTimes(ctx context.Context, w3c *web3go.Client, blkNums []types.BlockNumber,
	batchSize uint64) (map[uint64]uint64, error) {
	blockNum2Block, err := BatchGetBlocks(ctx, w3c, blkNums, false, batchSize)
	if err != nil {
		return nil, err
	}

	blockNum2Time := make(map[uint64]uint64)
	for _, block := range blockNum2Block {
		blockNum2Time[block.Number.Uint64()] = block.Timestamp
	}

	return blockNum2Time, nil
}

func BatchGetBlocks(ctx context.Context, w3c *web3go.Client, blkNums []types.BlockNumber, withDetail bool,
	batchSize uint64) (map[uint64]types.Block, error) {
	if len(blkNums) == 0 {
		return nil, errors.New("no block numbers")
	}

	blkNumSet := set.NewSet()
	for _, num := range blkNums {
		blkNumSet.Add(num)
	}

	blockNum2Block := make(map[uint64]types.Block)
	blkNumSlice := blkNumSet.ToSlice()
	blkNumSize := len(blkNumSlice)
	for i := 0; i < blkNumSize; i += int(batchSize) {
		end := i + int(batchSize)
		if end > blkNumSize {
			end = blkNumSize
		}
		blockNums := blkNumSlice[i:end]

		batch := make([]rpc.BatchElem, 0)
		for _, blkNum := range blockNums {
			elem := rpc.BatchElem{
				Method: "eth_getBlockByNumber",
				Args:   []interface{}{blkNum, withDetail},
				Result: new(types.Block),
			}
			batch = append(batch, elem)
		}

		err := w3c.Eth.BatchCallContext(ctx, batch)
		if err != nil {
			return nil, err
		}

		for _, elem := range batch {
			err := elem.Error
			if err != nil {
				return nil, err
			}
			block := elem.Result.(*types.Block)
			blockNum2Block[block.Number.Uint64()] = *block
		}
	}

	for _, num := range blkNums {
		if _, exists := blockNum2Block[uint64(num)]; !exists {
			return nil, errors.Errorf("block not found at number %v", num)
		}
	}


	return blockNum2Block, nil
}

func GetBlockByNumber(w3c *web3go.Client, blockNumber types.BlockNumber, isFull bool) (val *types.Block, err error) {
	block, err := w3c.Eth.BlockByNumber(blockNumber, isFull)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get block by number %v", blockNumber)
	}

	if block == nil {
		return nil, errors.Errorf("invalid block data at %v (must not be nil)", blockNumber)
	}

	if block.Number == nil {
		return nil, errors.Errorf("invalid block data at %v (block number must not be nil)", blockNumber)
	}

	return block, nil
}

type AlertContent struct {
	URL     string
	Reason  string
	Elapsed time.Duration
}

func (c AlertContent) String() string {
	if c.URL == "" {
		return fmt.Sprintf("---\nReason: %s\nElapsed: %v", c.Reason, c.Elapsed)
	}
	return fmt.Sprintf("---\nReason: %s\nURL: %s\nElapsed: %v", c.Reason, c.URL, c.Elapsed)
}

func AlertErr(ctx context.Context, title, channel string, err error,
	report health.TimedCounterConfig, health *health.TimedCounter, urls ...string) error {

	ch, ok := alert.DefaultManager().Channel(channel)
	if !ok {
		return errors.Errorf("Alert channel %s not found", channel)
	}

	var url string
	if len(urls) > 0 {
		url = urls[0]
	}

	if err == nil {
		if recovered, elapsed := health.OnSuccess(report); recovered {
			return ch.Send(ctx, &alert.Notification{
				Title: title, Content: AlertContent{url, "Recovered", elapsed}.String(),
				Severity: alert.SeverityLow,
			})
		}
		return nil
	}

	if unhealthy, unrecovered, elapsed := health.OnFailure(report); unhealthy || unrecovered {
		return ch.Send(ctx, &alert.Notification{
			Title: title, Content: AlertContent{url, err.Error(), elapsed}.String(),
			Severity: alert.SeverityHigh,
		})
	}

	return nil
}
