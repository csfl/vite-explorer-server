package accountchain

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"math/big"
)

var accountChainAccess = access.AccountAccess{}.New()

func GetAccountBalanceByTokenId(args ...[]byte) (*big.Int, error){
	return nil, nil
}
