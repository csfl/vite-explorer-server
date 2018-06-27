package account

import (
	"github.com/gin-gonic/gin"
	typeResponse "vite-explorer-server/type/response"
	typeRequest "vite-explorer-server/type/request"
	"vite-explorer-server/util"
)

func Detail(c *gin.Context)  {
	var requestAccountDetail typeRequest.AccountDetail

	if err := c.Bind(&requestAccountDetail); err != nil {
		utils.RespondError(c, 400, err)
		return
	}



	utils.RespondSuccess(c, &typeResponse.Account{
		AccountAddress: "0x123456",

		BlockHeight: "123123123",

		TokenList: []*typeResponse.AccountToken{{
			Balance: "123",
		}},
	}, "")
}