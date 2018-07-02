package response

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/ledger"
)

type AccountBlockList struct {
	blockList []*AccountBlock
}

func NewAccountBlockList (ledgerBlockList []*ledger.AccountBlock) *AccountBlockList{
	var blockList []*AccountBlock
	for _, legerBlock := range ledgerBlockList {
		blockList = append(blockList, NewAccountBlock(legerBlock))
	}
	return &AccountBlockList{
		blockList: blockList,
	}
}

func (abl *AccountBlockList) ToResponse () gin.H{
	var responseBlockList []gin.H
	for _, block := range abl.blockList {
		responseBlockList = append(responseBlockList, block.ToResponse())
	}
	return gin.H{
		"blockList": responseBlockList,
	}
}

type AccountBlock struct {
	Height string

	AccountAddress string

	To []byte

	FromHash []byte

	PrevHash []byte

	Status int

	Balance string

	Amount string

	Data string

	SnapshotTimestamp []byte

	Signature []byte

	Nounce []byte

	Difficulty []byte

	FAmount []byte
}

func NewAccountBlock (ledgerBlock *ledger.AccountBlock) *AccountBlock{
	return &AccountBlock{}
}

func (ab *AccountBlock) ToResponse () gin.H{
	return gin.H{
		"height": ab.Height,
	}
}