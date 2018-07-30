package accountchain

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/util"
	typeRequest "github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/service/accountchain"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	"github.com/vitelabs/vite-explorer-server/service/token"
	"github.com/pkg/errors"
	"fmt"
	"github.com/vitelabs/go-vite/log"
	"math/big"
)

func BlockList (c *gin.Context)  {
	var accountChainBlocklistQuery typeRequest.AccountChainBlocklist

	if err := c.BindJSON(&accountChainBlocklistQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	accountChainBlocklistQuery.Paging.PagingSetDefault()

	var blockList []*ledger.AccountBlock
	index, num, count := accountChainBlocklistQuery.Paging.Index, accountChainBlocklistQuery.Paging.Num, accountChainBlocklistQuery.Paging.Count
	log.Info("query for AccountBlockList's [index,num,count]=",index,num,count)
	var tokenList []*ledger.Token
	var totalNum *big.Int

	// Because there are only one token now, so query transaction list by tokenId and account is not implemented
	if accountChainBlocklistQuery.AccountAddress != "" {
		accountAddress, err:= types.HexToAddress(accountChainBlocklistQuery.AccountAddress)
		if err != nil {
			stackErr := errors.WithStack(errors.New(err.Error()))
			fmt.Println(stackErr)
			util.RespondFailed(c, 1, err, "")
			return
		}
		blockList, err = accountchain.GetBlockListByAccountAddress( index, num, count, &accountAddress)

		if err != nil {
			util.RespondFailed(c, 2, err, "")
			return
		}

		blockHeight, err := accountchain.GetLatestBlockHeightByAccountAddr(&accountAddress)

		if err != nil {
			util.RespondFailed(c, 9, err, "")
			return
		}

		totalNum = blockHeight

	} else if accountChainBlocklistQuery.TokenId != "" {
		tokenId, err := types.HexToTokenTypeId(accountChainBlocklistQuery.TokenId)

		if err != nil {
			util.RespondFailed(c, 3, err, "")
			return
		}
		blockList, err= accountchain.GetBlockListByTokenId( index, num, count, &tokenId)
		if err != nil {
			util.RespondFailed(c, 4, err, "")
			return
		}

		tokenInfo, err := token.GetTokenByTokenId(&tokenId)
		if err != nil {
			util.RespondFailed(c, 7, err, "")
		}

		tokenList = make([]*ledger.Token, len(blockList))
		for index := range tokenList {
			tokenList[index] = tokenInfo
		}

		totalNum, err = accountchain.GetTotalNumber()
		if err != nil {
			util.RespondFailed(c, 10, err, "")
		}
	} else {
		var err error
		blockList, err = accountchain.GetBlockList(index, num, count)
		if err != nil {
			util.RespondFailed(c, 5, err, "")
			return
		}

		totalNum, err = accountchain.GetTotalNumber()
		if err != nil {
			util.RespondFailed(c, 11, err, "")
		}


	}

	if tokenList == nil {
		for _, block := range blockList {
			var tokenInfo *ledger.Token
			if block.TokenId != nil {
				var err error
				tokenInfo, err = token.GetTokenByTokenId(block.TokenId)

				if err != nil {
					util.RespondFailed(c, 8, err, "")
					return
				}
			}

			tokenList = append(tokenList, tokenInfo)
		}
	}


	confirmInfoList, err := accountchain.GetConfirmInfoList(blockList)
	if err != nil {
		util.RespondFailed(c, 6, err, "")
		return
	}

	util.RespondSuccess(c, response.NewAccountBlockList(blockList, totalNum,confirmInfoList, tokenList), "")
	return

}

func Block (c *gin.Context)  {
	var accountChainBlockQuery typeRequest.AccountChainBlock

	if err := c.Bind(&accountChainBlockQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	blockHash, err := types.HexToHash(accountChainBlockQuery.BlockHash)
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return
	}
	block, err := accountchain.GetBlockByHash(&blockHash)
	if err != nil {
		util.RespondFailed(c, 2, err, "")
		return
	}

	confirmInfoList, err := accountchain.GetConfirmInfoList([]*ledger.AccountBlock{block})

	if err != nil {
		util.RespondFailed(c, 3, err, "")
		return
	}

	var confirmInfo gin.H
	if confirmInfoList != nil && len(confirmInfoList) > 0{
		confirmInfo = confirmInfoList[0]
	}

	tokenInfo, err := token.GetTokenByTokenId(block.TokenId)
	if err != nil {
		util.RespondFailed(c, 4, err, "")
	}

	util.RespondSuccess(c, response.NewAccountBlock(block, confirmInfo, tokenInfo), "")

}