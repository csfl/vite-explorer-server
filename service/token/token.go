package token

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/pkg/errors"
)

var tokenAccess = access.GetTokenAccess()
var errorHeader = "service.token"

func GetTokenByTokenId (tokenId *types.TokenTypeId) (*ledger.Token, error) {
	token, err := tokenAccess.GetByTokenId(tokenId)
	if err != nil {
		return nil, errors.Wrap(err, errorHeader + ".GetTokenByTokenId")
	}
	return token, err
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