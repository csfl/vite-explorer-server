package snapshotchain

import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/ledger/access"
	"math/big"
	"github.com/pkg/errors"
)

var errorHeader = "service.snapshotChain"
var snapshotChainAccess = access.GetSnapshotChainAccess()

func GetBlock (blockHash []byte) (*ledger.SnapshotBlock, error) {
	snapshotBlock, err := snapshotChainAccess.GetBlockByHash(blockHash)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlock(GetBlockByHash)")
	}
	return snapshotBlock, nil
}

func GetBlockList (index int, num int, count int) ([]*ledger.SnapshotBlock, error){
	snapshotBlockList, err := snapshotChainAccess.GetBlockList(index, num, count)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlockList(GetBlockList)")
	}
	return snapshotBlockList, nil
}

func GetSnapshotChainHeight() (* big.Int, error) {
	latestBlock, err := snapshotChainAccess.GetLatestBlock()
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetSnapshotChainHeight(GetLatestBlock)")
	}
	return latestBlock.Height, nil
}