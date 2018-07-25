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
	totalNumber *big.Int
}

func NewAccountBlockList (ledgerBlockList []*ledger.AccountBlock, totalNumber *big.Int, confirmInfoList []gin.H, tokenList []*ledger.Token) *AccountBlockList{
	var blockList []*AccountBlock
	for index, legerBlock := range ledgerBlockList {
		var confirmInfo gin.H
		if confirmInfoList != nil && index < len(confirmInfoList) {
			confirmInfo =  confirmInfoList[index]
		}
		blockList = append(blockList, NewAccountBlock(legerBlock, confirmInfo, tokenList[index]))
	}
	return &AccountBlockList{
		blockList: blockList,
		totalNumber: totalNumber,
	}
}

func (abl *AccountBlockList) ToResponse () gin.H{
	var responseBlockList []gin.H
	for _, block := range abl.blockList {
		responseBlockList = append(responseBlockList, block.ToResponse())
	}
	return gin.H{
		"blockList": responseBlockList,
		"totalNumber": abl.totalNumber.String(),
	}
}

type AccountBlock struct {
	Height *big.Int

	AccountAddress *types.Address

	To *types.Address

	From *types.Address

	Hash *types.Hash

	FromHash *types.Hash

	PrevHash *types.Hash

	Status int

	Timestamp uint64

	Token *Token

	Balance *big.Int

	Amount *big.Int

	Data string

	SnapshotTimestamp *types.Hash

	Signature []byte

	Nounce []byte

	Difficulty []byte

	FAmount *big.Int

	ConfirmBlockHash []byte

	ConfirmTimes *big.Int
}

func NewAccountBlock (ledgerBlock *ledger.AccountBlock, confirmInfo gin.H, token *ledger.Token) *AccountBlock{
	var responseToken *Token
	if token != nil {
		responseToken = NewToken(token)
	}
	accountBlock := &AccountBlock{
		Height: ledgerBlock.Meta.Height,
		AccountAddress: ledgerBlock.AccountAddress,
		To: ledgerBlock.To,
		From: ledgerBlock.From,
		Hash: ledgerBlock.Hash,
		FromHash: ledgerBlock.FromHash,
		PrevHash: ledgerBlock.PrevHash,
		Token: responseToken,

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
		accountBlock.ConfirmBlockHash = confirmInfo["confirmBlockHash"].([]byte)
		accountBlock.ConfirmTimes = confirmInfo["confirmTimes"].(*big.Int)
	}


	return accountBlock
}

func (ab *AccountBlock) ToResponse () gin.H{
	response := gin.H{
		"height": ab.Height.String(),
		"accountAddress": ab.AccountAddress.String(),
		"status": ab.Status,
		"balance": ab.Balance.String(),
		"amount": ab.Amount.String(),
		"data": ab.Data,

		"timestamp": ab.Timestamp,
		"signature": hex.EncodeToString(ab.Signature),

		"nounce": hex.EncodeToString(ab.Nounce),
		"difficulty": hex.EncodeToString(ab.Difficulty),

		"fAmount": ab.FAmount.String(),
		"confirmBlockHash": ab.ConfirmBlockHash,
		"confirmTimes": ab.ConfirmTimes,
	}

	if ab.Hash != nil {
		response["hash"] = ab.Hash.String()
	}

	if ab.PrevHash != nil {
		response["prevHash"] = ab.PrevHash.String()
	}

	if ab.Token != nil {
		response["token"] = ab.Token.ToResponse()
	}

	if ab.To != nil {
		response["to"] = ab.To.String()
	}
	if ab.From != nil {
		response["from"] = ab.From.String()
	}

	if ab.FromHash != nil {
		response["fromHash"] = ab.FromHash.String()
	}

	if ab.SnapshotTimestamp != nil {
		response["snapshotTimestamp"] = ab.SnapshotTimestamp.String()
	}
	return response
}