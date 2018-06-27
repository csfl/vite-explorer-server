package account

import (
	"vite-explorer-server/type/response"
	"github.com/vitelabs/go-vite/ledger"
)

func GetAccount (accountAddress []byte) *response.Account{
	account := &ledger.Account{}
	return &response.Account{
		AccountAddress: accountAddress,
		BlockHeight: account.GetBlockHeight(),
	}
}
