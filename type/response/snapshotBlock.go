package response

import (
	"math/big"
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/common/types"
)

type SnapshotBlock struct {
	Hash *types.Hash
	PrevHash *types.Hash
	Height *big.Int
	Producer *types.Address

	Snapshot map[string]*ledger.SnapshotItem
	Signature []byte
	Timestamp uint64
	Amount *big.Int
}

type SnapshotBlockList struct {
	BlockList []*SnapshotBlock
	TotalNumber *big.Int
}

type SnapshotChainHeight struct {
	ChainHeight *big.Int
}

func NewSnapshotBlock (snapshotBlcok *ledger.SnapshotBlock) *SnapshotBlock {
	return &SnapshotBlock{
		Hash: snapshotBlcok.Hash,
		PrevHash: snapshotBlcok.PrevHash,
		Height: snapshotBlcok.Height,
		Producer: snapshotBlcok.Producer,
		Snapshot: snapshotBlcok.Snapshot,
		Signature: snapshotBlcok.Signature,
		Timestamp: snapshotBlcok.Timestamp,
		Amount: snapshotBlcok.Amount,
	}
}

func NewSnapshotBlockList (snapshotBlcok []*ledger.SnapshotBlock, totalNumber *big.Int) *SnapshotBlockList {
	var blockList []*SnapshotBlock
	for _, snapshotBlock := range snapshotBlcok {
		blockList = append(blockList, NewSnapshotBlock(snapshotBlock))
	}
	return &SnapshotBlockList{
		BlockList: blockList,
		TotalNumber: totalNumber,
	}
}

func (sb *SnapshotBlock) ToResponse () gin.H{
	accountStatusList := make(map[string]gin.H)
	for k, v := range sb.Snapshot {
		accountStatusList[k] = gin.H{
			"accountBlockHash": v.AccountBlockHash,
			"accountBlockHeight": v.AccountBlockHeight,
		}
	}

	response := gin.H{
		"hash": sb.Hash.String(),
		"height": sb.Height.String(),
		"producer": sb.Producer.String(),
		"snapshot": accountStatusList,
		"signature": hex.EncodeToString(sb.Signature),
		"timestamp": sb.Timestamp,
		"amount": sb.Amount.String(),
	}

	if sb.PrevHash != nil {
		response["prevHash"] = sb.PrevHash.String()
	}
	return response
}

func (sbList *SnapshotBlockList) ToResponse () gin.H  {
	var hSbList []gin.H
	for _, hSb := range sbList.BlockList {
		hSbList = append(hSbList, hSb.ToResponse())
	}
	return gin.H{
		"blockList": hSbList,
		"totalNumber": sbList.TotalNumber.String(),
	}
}