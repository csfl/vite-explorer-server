package token

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
)

var tokenAccess = access.TokenAccess{}.GetInstance()

func GetTokenByTokenId (tokenId []byte) (*ledger.Token, error) {
	return tokenAccess.GetByTokenId(tokenId)
}

func GetTokenListByTokenName (tokenName string) ([]*ledger.Token, error) {
	return tokenAccess.GetListByTokenName(tokenName)
}

func GetTokenListByTokenSymbol (tokenSymbol string) ([]*ledger.Token, error) {
	return tokenAccess.GetListByTokenSymbol(tokenSymbol)
}

func GetTokenList (index int, num int, count int) ([]*ledger.Token, error) {
	return tokenAccess.GetList(index, num, count)
}