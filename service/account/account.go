package account

import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
)

func GetAccount (accountAddress []byte) *response.Account{
	account := &ledger.Account{}
	return &response.Account{
		AccountAddress: accountAddress,
		BlockHeight: account.GetBlockHeight(),
	}
}
