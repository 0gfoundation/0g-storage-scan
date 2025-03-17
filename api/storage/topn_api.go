package storage

import (
	"github.com/0glabs/0g-storage-scan/store"
	"github.com/Conflux-Chain/go-conflux-util/api"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	dataSizeTopn   = "data_size"
	storageFeeTopn = "storage_fee"
	txsTopn        = "txs"
	filesTopn      = "files"

	maxRecords = 100
)

func topnDataSize(c *gin.Context) (interface{}, error) {
	return topnByType(c, dataSizeTopn)
}

func topnStorageFee(c *gin.Context) (interface{}, error) {
	return topnByType(c, storageFeeTopn)
}

func topnTxs(c *gin.Context) (interface{}, error) {
	return topnByType(c, txsTopn)
}

func topnFiles(c *gin.Context) (interface{}, error) {
	return topnByType(c, filesTopn)
}

func topnByType(c *gin.Context, t string) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := store.TopnSpanTypes[topnP.SpanType]
	records := cache.topnAddresses[t][statSpan]

	result := make(map[string]interface{})
	switch t {
	case dataSizeTopn:
		list := make([]DataTopn, 0)
		for rank, r := range records {
			list = append(list, DataTopn{
				Topn:     Topn{Rank: rank + 1, Address: r.Address},
				DataSize: r.DataSize,
			})
		}
		result["list"] = list
	case storageFeeTopn:
		list := make([]FeeTopn, 0)
		for rank, r := range records {
			list = append(list, FeeTopn{
				Topn:       Topn{Rank: rank + 1, Address: r.Address},
				StorageFee: r.StorageFee,
			})
		}
		result["list"] = list
	case txsTopn:
		list := make([]TxsTopn, 0)
		for rank, r := range records {
			list = append(list, TxsTopn{
				Topn: Topn{Rank: rank + 1, Address: r.Address},
				Txs:  r.Txs,
			})
		}
		result["list"] = list
	case filesTopn:
		list := make([]FilesTopn, 0)
		for rank, r := range records {
			list = append(list, FilesTopn{
				Topn:  Topn{Rank: rank + 1, Address: r.Address},
				Files: r.Files,
			})
		}
		result["list"] = list
	default:
		return nil, api.ErrValidation(errors.Errorf("Invalid topn type %v", t))
	}

	return result, nil
}

func topnReward(c *gin.Context) (interface{}, error) {
	var topnP topnParam
	if err := c.ShouldBind(&topnP); err != nil {
		return nil, api.ErrValidation(errors.Errorf("Invalid topn query param"))
	}

	statSpan := store.TopnSpanTypes[topnP.SpanType]
	miners := cache.topnMiners[statSpan]
	if len(miners) == 0 {
		return map[string]interface{}{"list": []RewardTopn{}}, nil
	}

	addrIDs := make([]uint64, 0)
	for _, miner := range miners {
		addrIDs = append(addrIDs, miner.ID)
	}
	minerMap, err := db.BatchGetMiners(addrIDs)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to batch get miners")
	}

	list := make([]RewardTopn, 0)
	for rank, m := range miners {
		reward := RewardTopn{
			Rank:           rank + 1,
			Address:        m.Address,
			Amount:         m.Amount,
			WinCount:       m.WinCount,
			MiningAttempts: minerMap[m.ID].MiningAttempts,
		}
		if reward.WinCount > reward.MiningAttempts {
			reward.MiningAttempts = reward.WinCount
		}
		list = append(list, reward)
	}

	return map[string]interface{}{"list": list}, nil
}
