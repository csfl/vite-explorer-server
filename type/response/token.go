package response

import (
	"math/big"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/common/types"
)

type TokenList struct {
	tokenList []*Token
}

func NewTokenList (tokens []*ledger.Token) *TokenList{
	var tokenList []*Token
	for _, token := range tokens {
		tokenList = append(tokenList, NewToken(token))
	}
	return &TokenList{
		tokenList: tokenList,
	}
}


func (t *TokenList) ToResponse () gin.H{
	var tokenList []gin.H
	for _, token := range t.tokenList {
		tokenList = append(tokenList, token.ToResponse())
	}
	return gin.H{
		"tokenList": tokenList,
	}
}

type Token struct {
	Name string
	Id *types.TokenTypeId
	Introduction string

	Symbol string

	Owner []byte

	Decimals int

	TotalSupply *big.Int

	TransactionNumber *big.Int
}

func NewToken (ledgerToken *ledger.Token) *Token {
	return &Token{}
}

func (t*Token) ToResponse () gin.H {
	return gin.H{
		"name": t.Name,
		"id": t.Id.String(),
		"introduction": t.Introduction,

		"symbol": t.Symbol,

		"decimals": t.Decimals,

		"totalSupply": t.TotalSupply.String(),

		"transactionNumber": t.TransactionNumber.String(),
	}
}