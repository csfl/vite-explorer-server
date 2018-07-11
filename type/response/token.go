package response

import (
	"math/big"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/common/types"
)

type TokenList struct {
	tokenList []*Token
	totalNumber *big.Int
}

func NewTokenList (tokens []*ledger.Token) *TokenList{
	var tokenList []*Token
	for _, token := range tokens {
		tokenList = append(tokenList, NewToken(token))
	}
	return &TokenList{
		tokenList: tokenList,
		totalNumber: big.NewInt(1),
	}
}


func (t *TokenList) ToResponse () gin.H{
	var tokenList []gin.H
	for _, token := range t.tokenList {
		tokenList = append(tokenList, token.ToResponse())
	}
	return gin.H{
		"tokenList": tokenList,
		"totalNumber": t.totalNumber.String(),
	}
}

type Token struct {
	Name string
	Id *types.TokenTypeId

	Symbol string

	Owner *types.Address

	Decimals int

	TotalSupply *big.Int
}

func NewToken (ledgerToken *ledger.Token) *Token {
	return &Token{
		Name: ledgerToken.Mintage.Name,
		Id: ledgerToken.Mintage.Id,
		Symbol: ledgerToken.Mintage.Symbol,

		Owner: ledgerToken.Mintage.Owner,
		Decimals: ledgerToken.Mintage.Decimals,
		TotalSupply: ledgerToken.Mintage.TotalSupply,
	}
}

func (t*Token) ToResponse () gin.H {
	return gin.H{
		"name": t.Name,
		"id": t.Id.String(),

		"symbol": t.Symbol,
		"owner": t.Owner.String(),

		"decimals": t.Decimals,

		"totalSupply": t.TotalSupply.String(),
	}
}