package token

import (
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/type/request"
	"github.com/vitelabs/vite-explorer-server/util"
	tokenService "github.com/vitelabs/vite-explorer-server/service/token"
	"github.com/vitelabs/vite-explorer-server/type/response"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/common/types"
)

func List (c *gin.Context)  {
	var tokenListQuery request.TokenList

	if err := c.BindJSON(&tokenListQuery); err != nil{
		util.RespondError(c, 400, err)
		return
	}

	tokenListQuery.PagingSetDefault()
	ledgerTokenList, err := tokenService.GetTokenList(tokenListQuery.Index, tokenListQuery.Num, tokenListQuery.Count)

	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return
	}


	util.RespondSuccess(c, response.NewTokenList(ledgerTokenList), "")

}

func Detail (c *gin.Context) {
	var tokenDetailQuery request.TokenDetail

	if err := c.Bind(&tokenDetailQuery); err != nil {
		util.RespondError(c, 400, err)
		return
	}


	var tokenList *response.TokenList
	if tokenDetailQuery.TokenId != "" {
		tokenId, err := types.HexToTokenTypeId(tokenDetailQuery.TokenId)

		if err != nil {
			util.RespondFailed(c, 1, err, "")
			return
		}

		ledgerToken, err := tokenService.GetTokenByTokenId(&tokenId)

		if err != nil {
			util.RespondFailed(c, 2, err, "")
			return
		}

		tokenList = response.NewTokenList([]*ledger.Token{ledgerToken})

	} else if tokenDetailQuery.TokenName != "" {

		ledgerTokenList, err := tokenService.GetTokenListByTokenName(tokenDetailQuery.TokenName)

		if err != nil {
			util.RespondFailed(c, 3, err, "")
			return
		}

		tokenList = response.NewTokenList(ledgerTokenList)
	} else if tokenDetailQuery.TokenSymbol != "" {
		ledgerTokenList, err := tokenService.GetTokenListByTokenSymbol(tokenDetailQuery.TokenSymbol)

		if err != nil {
			util.RespondFailed(c, 4, err, "")
			return
		}

		tokenList = response.NewTokenList(ledgerTokenList)
	}

	util.RespondSuccess(c, tokenList, "")
}