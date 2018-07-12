package snapshotchain

import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/ledger/access"
	"math/big"
)


var snapshotChainAccess = access.GetSnapshotChainAccess()

func GetBlock (blockHash []byte) (*ledger.SnapshotBlock, error) {
	return snapshotChainAccess.GetBlockByHash(blockHash)
}

func GetBlockList (index int, num int, count int) ([]*ledger.SnapshotBlock, error){
	return snapshotChainAccess.GetBlockList(index, num, count)
}

func GetSnapshotChainHeight() (* big.Int, error) {
	latestBlock, err := snapshotChainAccess.GetLatestBlock()
	if err != nil {
		return nil, err
	}
	return latestBlock.Height, nil
}