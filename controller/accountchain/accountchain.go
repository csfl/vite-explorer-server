package accountchain

import (
	"github.com/gin-gonic/gin"
	"vite-explorer-server/utils"
	typeRequest "vite-explorer-server/types/request"
)

func BlockList (c *gin.Context)  {
	var accountchainBlocklistQuery typeRequest.AccountchainBlocklist

	if err := c.BindJSON(&accountchainBlocklistQuery); err != nil {
		utils.RespondError(c, 400, err)
		return
	}
}

func Block (c *gin.Context)  {
}