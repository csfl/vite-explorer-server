package account

import (
	"vite-explorer-server/type/response"
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
)
var accountAccess = access.AccountAccess{}.New()
var accountChainAccess = access.AccountAccess{}.New()

func GetAccount (accountAddress []byte) *response.Account {
	account := &ledger.Account{}
	accountMeta := GetAccountMeta(accountAddress)
	return &response.Account{
		AccountAddress: accountAddress,
		BlockHeight: account.GetBlockHeight(),
		TokenList: GetAccountTokenList(accountMeta),
	}
}

func GetAccountMeta (accountAddress []byte) *ledger.AccountMeta {
	accountMeta := accountAccess.GetAccountMeta(accountAddress)
	return &accountMeta
}

func GetAccountTokenList (accountMeta *ledger.AccountMeta) error {
	accountId := accountMeta.AccountId
	var accountTokenList []*response.AccountToken
	for _, accountSimpleToken := range accountMeta.TokenList {

		// transform the AccountAddress into  AccountAddressHash
		//
		var accountAddressHash byte

		accountToken, err := GetAccountToken(accountId.Bytes(), accountSimpleToken.LastAccountBlockHeight.Bytes(),
			accountAddressHash, accountSimpleToken.TokenId)
		if err != nil {
			fmt.println(err)
			return err
		}
		accountTokenList =  append(accountTokenList, accountToken)
	}
	return nil
}

func GetAccountToken (tokenid []byte, args ...byte) (*response.AccountToken, error) {
	token, err := serviceToken.GetToken(tokenid)
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return nil, err
	}

	// assembly the key to find accountBlock
	//
	accountBlockKey := make([]byte, "c.".Bytes())
	for _, value := range args {
		accountBlockKey = append(accountBlockKey, value)
	}
	balance, balanceErr := accountChainAccess.GetAccountBalance(accountBlockKey)
	if balanceErr != nil {
		util.RespondFailed(c, 1, err, "")
		return nil, balanceErr
	}

	return &response.AccountToken{
		Token: token,
		Balance: balance.String(),
	},err
}