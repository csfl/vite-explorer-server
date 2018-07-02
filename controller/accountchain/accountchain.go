package accountchain

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/util"
	typeRequest "github.com/vitelabs/vite-explorer-server/type/request"
)

func BlockList (c *gin.Context)  {
	var accountchainBlocklistQuery typeRequest.AccountChainBlocklist

	if err := c.BindJSON(&accountchainBlocklistQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}

	accountchainBlocklistQuery.PagingSetDefault()
}

func Block (c *gin.Context)  {
}