package response

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/ledger"
	"math/big"
	"github.com/vitelabs/go-vite/common/types"
	"encoding/hex"
)

type AccountBlockList struct {
	blockList []*AccountBlock
}

func NewAccountBlockList (ledgerBlockList []*ledger.AccountBlock, confirmInfoList []gin.H) *AccountBlockList{
	var blockList []*AccountBlock
	for index, legerBlock := range ledgerBlockList {
		var confirmInfo gin.H
		if confirmInfoList != nil && index < len(confirmInfoList) {
			confirmInfo =  confirmInfoList[index]
		}
		blockList = append(blockList, NewAccountBlock(legerBlock, confirmInfo))
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
	Height *big.Int

	AccountAddress *types.Address

	To *types.Address

	From *types.Address

	Hash []byte

	FromHash []byte

	PrevHash []byte

	Status int

	Timestamp uint64

	Balance *big.Int

	Amount *big.Int

	Data string

	SnapshotTimestamp []byte

	Signature []byte

	Nounce []byte

	Difficulty []byte

	FAmount *big.Int

	ConfirmBlockHash []byte

	ConfirmTimes *big.Int
}

func NewAccountBlock (ledgerBlock *ledger.AccountBlock, confirmInfo gin.H) *AccountBlock{
	accountBlock := &AccountBlock{
		Height: ledgerBlock.Meta.Height,
		AccountAddress: ledgerBlock.AccountAddress,
		To: ledgerBlock.To,
		From: ledgerBlock.From,
		Hash: ledgerBlock.Hash,
		FromHash: ledgerBlock.FromHash,
		PrevHash: ledgerBlock.PrevHash,

		Status: ledgerBlock.Meta.Status,
		Balance: ledgerBlock.Balance,
		Amount: ledgerBlock.Amount,
		Data: ledgerBlock.Data,

		Timestamp: ledgerBlock.Timestamp,
		SnapshotTimestamp: ledgerBlock.SnapshotTimestamp,
		Signature: ledgerBlock.Signature,
		Nounce: ledgerBlock.Nounce,

		Difficulty: ledgerBlock.Difficulty,
		FAmount: ledgerBlock.FAmount,
	}

	if confirmInfo != nil {
		accountBlock.ConfirmBlockHash = confirmInfo["ConfirmBlockHash"].([]byte)
		accountBlock.ConfirmTimes = confirmInfo["ConfirmTimes"].(*big.Int)
	}

	return accountBlock
}

func (ab *AccountBlock) ToResponse () gin.H{
	response := gin.H{
		"height": ab.Height.String(),
		"accountAddress": ab.AccountAddress.String(),
		"hash": hex.EncodeToString(ab.Hash),
		"prevHash": hex.EncodeToString(ab.PrevHash),
		"status": ab.Status,
		"balance": ab.Balance.String(),
		"amount": ab.Amount.String(),
		"data": ab.Data,

		"timestamp": ab.Timestamp,
		"snapshotTimestamp": hex.EncodeToString(ab.SnapshotTimestamp),
		"signature": hex.EncodeToString(ab.Signature),

		"nounce": hex.EncodeToString(ab.Nounce),
		"difficulty": hex.EncodeToString(ab.Difficulty),

		"fAmount": ab.FAmount.String(),
	}

	if ab.To != nil {
		response["to"] = ab.To.String()
	}
	if ab.From != nil {
		response["from"] = ab.From.String()
	}

	if ab.FromHash != nil {
		response["fromHash"] = hex.EncodeToString(ab.FromHash)
	}
	return response
}