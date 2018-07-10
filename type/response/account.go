package response

import (
	"math/big"
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/common/types"
)

type AccountToken struct {
	Balance string

	Token *Token
}

type Account struct {
	AccountAddress  *types.Address

	BlockHeight *big.Int

	TokenList []*AccountToken
}

func (account *Account) ToResponse () gin.H {
	var hTokenList []gin.H
	for _, token := range account.TokenList {
		hTokenList = append(hTokenList, token.ToResponse())
	}
	return gin.H{
		"accountAddress": account.AccountAddress.String(),
		"blockHeight": account.BlockHeight.String(),
		"tokenList": hTokenList,
	}
}

func (at *AccountToken) ToResponse () gin.H {
	return gin.H{
		"balance": at.Balance,
		"token":   at.Token.ToResponse(),
	}
}

func NewAccount (accountAddress *types.Address, blockHeight *big.Int, accountTokenList []*AccountToken) *Account {
	return &Account{
		AccountAddress: accountAddress,
		BlockHeight: blockHeight,
		TokenList: accountTokenList,
	}
}

func NewAccountToken (token *ledger.Token, balance *big.Int) *AccountToken {
	return &AccountToken{
		Balance: balance.String(),
		Token: NewToken(token),
	}
}