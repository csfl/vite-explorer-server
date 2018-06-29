package account

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	"github.com/vitelabs/vite-explorer-server/util"
	"log"
	serviceToken "github.com/vitelabs/vite-explorer-server/service/token"
	serviceAccountChain "github.com/vitelabs/vite-explorer-server/service/accountchain"
)
var accountAccess = access.AccountAccess{}.New()


func GetAccount (accountAddress []byte) (*response.Account, error) {
	account := &ledger.Account{}
	accountMeta, err := accountAccess.GetAccountMeta(accountAddress)
	if err != nil{
		log.Fatal(err)
		return nil, err
	}
	accountTokenList, err := GetAccountTokenList(accountMeta)
	if err != nil {
		util.RespondFailed(c, 1, err, "")
		return  nil, err
	}
	return &response.Account{
		AccountAddress: accountAddress,
		BlockHeight: account.GetBlockHeight(),
		TokenList: accountTokenList,
	}, err
}

func GetAccountTokenList (accountMeta *ledger.AccountMeta) ([]*response.AccountToken, error) {
	accountId := accountMeta.AccountId
	var accountTokenList []*response.AccountToken
	for _, accountSimpleToken := range accountMeta.TokenList {
		accountToken, err := GetAccountToken(accountSimpleToken.TokenId, accountId.Bytes(),
			accountSimpleToken.LastAccountBlockHeight.Bytes())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		accountTokenList =  append(accountTokenList, accountToken)
	}
	return accountTokenList, nil
}

func GetAccountToken (tokenId []byte, args ...[]byte) (*response.AccountToken, error) {
	token, err := serviceToken.GetToken(tokenId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	balance, balanceErr := serviceAccountChain.GetAccountBalanceByTokenId(args)
	if balanceErr != nil {
		log.Fatal(err)
		return nil, balanceErr
	}

	return &response.AccountToken{
		Token: serviceToken.NewToken(token),
		Balance: balance.String(),
	},err
}