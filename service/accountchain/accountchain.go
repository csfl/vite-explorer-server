package accountchain

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"math/big"
)

var accountChainAccess = access.AccountChainAccess{}.New()

func GetAccountBalance(keyPartionList ...interface{}) (*big.Int, error){
	balance, err := accountChainAccess.GetAccountBalance(keyPartionList)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
