package accountchain

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/util"
	typeRequest "github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/service/accountchain"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	"encoding/hex"
)

func BlockList (c *gin.Context)  {
	var accountChainBlocklistQuery typeRequest.AccountChainBlocklist

	if err := c.BindJSON(&accountChainBlocklistQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	accountChainBlocklistQuery.PagingSetDefault()

	var blockList []*ledger.AccountBlock
	index, num, count := accountChainBlocklistQuery.Index, accountChainBlocklistQuery.Num, accountChainBlocklistQuery.Count

	if accountChainBlocklistQuery.AccountAddress != "" {
		accountAddress, err:= types.HexToAddress(accountChainBlocklistQuery.AccountAddress)
		if err != nil {
			util.RespondFailed(c, 1, err, "")
			return
		}
		blockList, err = accountchain.GetBlockListByAccountAddress( index, num, count, &accountAddress)
		if err != nil {
			util.RespondFailed(c, 2, err, "")
			return
		}
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
	} else {
		var err error
		blockList, err = accountchain.GetBlockList(index, num, count)
		if err != nil {
			util.RespondFailed(c, 5, err, "")
			return
		}
	}


	confirmInfoList, err := accountchain.GetConfirmInfoList(blockList)

	if err != nil {
		util.RespondFailed(c, 6, err, "")
		return
	}

	util.RespondSuccess(c, response.NewAccountBlockList(blockList, confirmInfoList), "")
	return

}

func Block (c *gin.Context)  {
	var accountChainBlockQuery typeRequest.AccountChainBlock

	if err := c.Bind(&accountChainBlockQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	blockHash, err := hex.DecodeString(accountChainBlockQuery.BlockHash)
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return
	}
	block, err := accountchain.GetBlockByHash(blockHash)
	if err != nil {
		util.RespondFailed(c, 2, err, "")
		return
	}

	confirmInfoList, err := accountchain.GetConfirmInfoList([]*ledger.AccountBlock{block})

	if err != nil {
		util.RespondFailed(c, 3, err, "")
		return
	}

	util.RespondSuccess(c, response.NewAccountBlock(block, confirmInfoList[0]), "")

}