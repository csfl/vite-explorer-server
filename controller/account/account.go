package account

import (
	"github.com/gin-gonic/gin"
	typeRequest "github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/util"
	serviceAccount "github.com/vitelabs/vite-explorer-server/service/account"
)

func Detail(c *gin.Context)  {
	var accountDetailQuery typeRequest.AccountDetail

	if err := c.Bind(&accountDetailQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}
	//
	//account := serviceAccount.GetAccount([]byte{1, 2, 3})
	//token, err := serviceToken.GetToken([]byte{4, 5, 6})
	//if err != nil {
	//	util.RespondFailed(c, 1, err, "")
	//	return
	//}
	//
	//fmt.Println(token)
	//util.RespondSuccess(c, account, "")

	account, err := serviceAccount.GetAccount(accountDetailQuery.AccountAddress)
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return
	}
	util.RespondSuccess(c, account, "")

}