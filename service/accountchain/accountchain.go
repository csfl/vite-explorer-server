package accountchain


import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/common/types"
	"math/big"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)


var accountChainAccess = access.GetAccountChainAccess()
var errorHeader = "service.accountChain"

func GetConfirmInfoList (blockList []*ledger.AccountBlock) ([]gin.H, error) {
	var confirmInfoList []gin.H
	for _, block := range blockList {

		confirmSnapshotBlock, err:= accountChainAccess.GetConfirmBlock(block)

		if err != nil {
			return nil, errors.Wrap(err, errorHeader + ".GetConfirmInfoList(GetConfirmBlock)")
		}

		if confirmSnapshotBlock == nil {
			return nil, nil
		}

		confirmTimes, err := accountChainAccess.GetConfirmTimes(confirmSnapshotBlock)
		if err != nil {
			return nil, errors.Wrap(err, errorHeader + ".GetConfirmInfoList(GetConfirmTimes)")
		}
		confirmInfoList = append(confirmInfoList, gin.H{
			"confirmBlockHash": confirmSnapshotBlock.Hash,
			"confirmTimes": confirmTimes,
		})
	}

	return confirmInfoList, nil
}

func GetBlockByHash (blockHash []byte) (*ledger.AccountBlock, error){
	blocks, err := accountChainAccess.GetBlockByHash(blockHash)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlockByHash")
	}

	return blocks, nil
}

func GetBlockListByAccountAddress (index int, num int, count int, accountAddress *types.Address) ([]*ledger.AccountBlock, error){
	blocks, err := accountChainAccess.GetBlockListByAccountAddress(index, num, count, accountAddress)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlockListByAccountAddress")
	}
	return blocks, nil
}

func GetBlockListByTokenId (index int, num int, count int, tokenId *types.TokenTypeId) ([]*ledger.AccountBlock, error){
	blocks, err := accountChainAccess.GetBlockListByTokenId(index, num, count, tokenId)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlockListByTokenId")
	}
	return blocks, nil
}

func GetBlockList (index int, num int, count int) ([]*ledger.AccountBlock, error){
	blocks, err := accountChainAccess.GetBlockList(index, num, count)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetBlockList")
	}
	return blocks, nil
}

func GetAccountBalance(accountId *big.Int, blockHeight *big.Int) (*big.Int, error){
	balance, err := accountChainAccess.GetAccountBalance(accountId,blockHeight)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetAccountBalance")
	}
	return balance, nil
}

func GetLatestBlockHeightByAccountId (accountId *big.Int) (* big.Int, error){
	blockHeight, err := accountChainAccess.GetLatestBlockHeightByAccountId(accountId)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetLatestBlockHeightByAccountId")
	}
	return blockHeight, nil
}