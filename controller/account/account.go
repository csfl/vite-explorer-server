package account

import (
	"github.com/gin-gonic/gin"
	typeRequest "github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/util"
	serviceAccount "github.com/vitelabs/vite-explorer-server/service/account"
	"github.com/vitelabs/go-vite/common/types"
	"errors"
)

func Detail(c *gin.Context)  {
	var accountDetailQuery typeRequest.AccountDetail

	if err := c.Bind(&accountDetailQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}
	if !types.IsValidHexAddress(accountDetailQuery.AccountAddress) {
		util.RespondFailed(c, 1, errors.New("AccountAddress is invalid"), "")
		return
	}
	accountAddress, err:= types.HexToAddress(accountDetailQuery.AccountAddress)
	if err != nil {
		util.RespondFailed(c, 6, err, "")
		return
	}
	account, err := serviceAccount.GetAccount(c, &accountAddress)
	if err != nil {
		return
	}
	util.RespondSuccess(c, account,"")
}