package token

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
)

var tokenAccess = access.TokenAccess{}.New()

func GetToken (tokenId []byte) (*ledger.Token, error) {
	token, err := tokenAccess.GetByTokenId(tokenId)
	return token, err
}

