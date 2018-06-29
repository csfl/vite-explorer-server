package response

import (
	"math/big"
	"github.com/gin-gonic/gin"
)

type AccountToken struct {
	Balance string

	Token Token
}

type Account struct {
	AccountAddress []byte

	BlockHeight *big.Int

	TokenList []*AccountToken
}

func (account *Account) ToResponse () gin.H {
	return gin.H{
		"accountAddress": "123",
		"blockHeight": account.BlockHeight.String(),
		"TokenList": account.TokenList.String(),
	}
}

