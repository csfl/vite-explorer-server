package snapshotchain

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/util"
)

func BlockList (c *gin.Context)  {
	var snapshotChainBlockListQuery request.SnapshotChainBlocklist

	if err := c.BindJSON(&snapshotChainBlockListQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	snapshotChainBlockListQuery.PagingSetDefault()

}

func Block (c *gin.Context)  {
	var snapshotchainBlockQuery request.SnapshotChainBlock

	if err := c.Bind(&snapshotchainBlockQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}
}
