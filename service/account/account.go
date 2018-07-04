package account

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	serviceToken "github.com/vitelabs/vite-explorer-server/service/token"
	serviceAccountChain "github.com/vitelabs/vite-explorer-server/service/accountchain"
	"math/big"
)
var accountAccess = access.GetAccountAccess()


func GetAccount (accountAddress string) (*response.Account, error) {
	account := &ledger.Account{}
	accountMeta, err := accountAccess.GetAccountMeta(accountAddress)
	if err != nil{
		return nil, err
	}
	accountTokenList, err := GetAccountTokenList(accountMeta)
	if err != nil {
		return  nil, err
	}
	return &response.Account{
		AccountAddress: []byte(accountAddress),
		BlockHeight: account.GetBlockHeight(),
		TokenList: accountTokenList,
	}, nil
}

func GetAccountTokenList (accountMeta *ledger.AccountMeta) ([]*response.AccountToken, error) {
	accountId := accountMeta.AccountId
	var accountTokenList []*response.AccountToken
	for _, accountSimpleToken := range accountMeta.TokenList {
		accountToken, err := GetAccountToken(accountSimpleToken.TokenId, accountId,
			accountSimpleToken.LastAccountBlockHeight)
		if err != nil {
			return nil, err
		}
		accountTokenList =  append(accountTokenList, accountToken)
	}
	return accountTokenList, nil
}

func GetAccountToken (tokenId []byte, accountId *big.Int, blockHeight *big.Int) (*response.AccountToken, error) {
	token, err := serviceToken.GetTokenByTokenId(tokenId)
	if err != nil {
		return nil, err
	}
	balance, balanceErr := serviceAccountChain.GetAccountBalance(accountId, blockHeight)
	if balanceErr != nil {
		return nil, balanceErr
	}
	return &response.AccountToken{
		Token: *response.NewToken(token),
		Balance: balance.String(),
	},nil
}