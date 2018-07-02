package snapshotchain

import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/ledger/access"
)


var snapshotChainAccess = access.GetSnapshotChainAccess()

func GetBlock (blockHash []byte) (*ledger.SnapshotBlock, error){
	return snapshotChainAccess.GetBlockByHash(blockHash)
}


func GetBlockList (index int, num int, count int) ([]*ledger.SnapshotBlock, error){
	return snapshotChainAccess.GetBlockList(index, num, count)
}