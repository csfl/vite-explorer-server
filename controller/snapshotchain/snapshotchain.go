package snapshotchain

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/util"
	"github.com/vitelabs/go-vite/ledger"
	serviceSnapshotChain "github.com/vitelabs/vite-explorer-server/service/snapshotchain"
	"github.com/vitelabs/vite-explorer-server/type/response"
	"encoding/hex"
	"github.com/vitelabs/go-vite/log"
)

func BlockList (c *gin.Context)  {
	var snapshotChainBlockListQuery request.SnapshotChainBlocklist
	if err := c.BindJSON(&snapshotChainBlockListQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}
	snapshotChainBlockListQuery.Paging.PagingSetDefault()

	var blockList []*ledger.SnapshotBlock
	index, num, count := snapshotChainBlockListQuery.Paging.Index, snapshotChainBlockListQuery.Paging.Num, snapshotChainBlockListQuery.Paging.Count
	log.Info("query for SnapshotChainBlockList's [index,num,count]=",index,num,count)
	blockList, err := serviceSnapshotChain.GetBlockList(index, num, count)
	if err != nil {
		util.RespondFailed(c, 1, err,"")
		return
	}
	totalNumber, err := serviceSnapshotChain.GetSnapshotChainHeight()
	if err != nil {
		util.RespondFailed(c, 2, err, "")
		return
	}
	util.RespondSuccess(c, response.NewSnapshotBlockList(blockList, totalNumber),"")

}

func Block (c *gin.Context)  {
	var snapshotChainBlockQuery request.SnapshotChainBlock

	if err := c.Bind(&snapshotChainBlockQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}
	blockHash, err := hex.DecodeString(snapshotChainBlockQuery.BlockHash)
	if err != nil {
		util.RespondFailed(c, 3, err, "")
		return
	}
	snapshotBlock, err := serviceSnapshotChain.GetBlock(blockHash)
	if err != nil {
		util.RespondFailed(c, 4, err, "")
		return
	}

	util.RespondSuccess(c, response.NewSnapshotBlock(snapshotBlock), "")
}