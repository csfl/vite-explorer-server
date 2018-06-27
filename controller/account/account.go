package account

import (
	"github.com/gin-gonic/gin"
	typeRequest "vite-explorer-server/type/request"
	"vite-explorer-server/util"
	serviceAccount "vite-explorer-server/service/account"
)

func Detail(c *gin.Context)  {
	var accountDetailQuery typeRequest.AccountDetail

	if err := c.Bind(&accountDetailQuery); err != nil {
		utils.RespondError(c, 400, err)
		return
	}

	account := serviceAccount.GetAccount([]byte{1, 2, 3})

	utils.RespondSuccess(c, account, "")
}