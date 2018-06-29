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
	//var tokenList
	//for _, accountToken := range account.TokenList{
	//	tokenList = append(tokenList, [accountToken, accountToken.Token.ToResponse()])
	//}
	//return gin.H{
	//	"accountAddress": account.AccountAddress,
	//	"blockHeight": account.BlockHeight.String(),
	//	"TokenList": tokenList,
	//}
	return gin.H{}
}

func NewAccount () *Account{
	return &Account{

	}
}

