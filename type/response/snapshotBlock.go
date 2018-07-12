package response

import (
	"math/big"
	"github.com/gin-gonic/gin"
	"encoding/hex"
	"github.com/vitelabs/go-vite/ledger"
)

type SnapshotBlock struct {
	Hash []byte
	PrevHash []byte
	Height *big.Int
	Producer []byte
	Snapshot map[string][]byte
	Signature []byte
	Timestamp uint64
	Amount *big.Int
}

type SnapshotBlockList struct {
	BlockList []*SnapshotBlock
	ChainHeight *big.Int
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

func NewSnapshotBlockList (snapshotBlcok []*ledger.SnapshotBlock, chainHeight *big.Int) *SnapshotBlockList {
	var blockList []*SnapshotBlock
	for _, snapshotBlock := range snapshotBlcok {
		blockList = append(blockList, NewSnapshotBlock(snapshotBlock))
	}
	return &SnapshotBlockList{
		BlockList: blockList,
		ChainHeight: chainHeight,
	}
}

func NewSnapshotChainHeight (chainHeight *big.Int) * SnapshotChainHeight {
	return &SnapshotChainHeight{
		ChainHeight: chainHeight,
		}
}

func (sb *SnapshotBlock) ToResponse () gin.H{
	accountStatusList := make(map[string]string)
	for k, v := range sb.Snapshot {
		accountStatusList[k] = hex.EncodeToString(v)
	}
	return gin.H{
		"hash": hex.EncodeToString(sb.Hash),
		"prevHash": hex.EncodeToString(sb.PrevHash),
		"height": sb.Height.String(),
		"producer": hex.EncodeToString(sb.Producer),
		"snapshot": accountStatusList,
		"signature": hex.EncodeToString(sb.Signature),
		"timestamp": sb.Timestamp,
		"amount": sb.Amount.String(),
	}
}

func (sbList *SnapshotBlockList) ToResponse () gin.H  {
	var hSbList []gin.H
	for _, hSb := range sbList.BlockList {
		hSbList = append(hSbList, hSb.ToResponse())
	}
	return gin.H{
		"blockList": hSbList,
		"chainHeight": sbList.ChainHeight.String(),
	}
}

func (sch *SnapshotChainHeight) ToResponse () gin.H {
	return gin.H{
		"chainHeight": sch.ChainHeight.String(),
	}
}