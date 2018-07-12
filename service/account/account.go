package account

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	serviceToken "github.com/vitelabs/vite-explorer-server/service/token"
	serviceAccountChain "github.com/vitelabs/vite-explorer-server/service/accountchain"
	"math/big"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/pkg/errors"
)
var accountAccess = access.GetAccountAccess()
var errorHeader = "service.account"

func GetAccount (accountAddress *types.Address) (*response.Account, error) {
	accountMeta, gmErr:= accountAccess.GetAccountMeta(accountAddress)
	if gmErr != nil{
		//return nil, errors.New("Error getting accountMeta.")
		return nil, errors.Wrap(gmErr, errorHeader + ".GetAccount(GetAccountMeta)")
	}
	accountBLockHeight, ghErr := serviceAccountChain.GetLatestBlockHeightByAccountId(accountMeta.AccountId)
	if ghErr != nil{
		//return nil, errors.New("Error getting account block height (number of trades)")
		return nil, errors.Wrap(ghErr, errorHeader + ".GetAccount(GetLatestBlockHeightByAccountId)")
	}
	accountTokenList, err:= GetAccountTokenList(accountMeta)
	if err != nil {
		return  nil, err
	}
	return response.NewAccount(accountAddress, accountBLockHeight, accountTokenList), nil
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

func GetAccountToken (tokenId *types.TokenTypeId, accountId *big.Int, blockHeight *big.Int) (*response.AccountToken, error) {
	token, gtErr := serviceToken.GetTokenByTokenId(tokenId)
	if gtErr != nil {
		//return nil, errors.New("Error getting token information")
		return nil, errors.Wrap(gtErr, errorHeader + ".GetAccountToken(GetTokenByTokenId)")
	}
	balance, balanceErr := serviceAccountChain.GetAccountBalance(accountId, blockHeight)
	if balanceErr != nil {
		//return nil,errors.New("Error getting current account token balance")
		return nil, errors.Wrap(balanceErr, errorHeader + ".GetAccountToken(GetAccountBalance)")
	}
	return response.NewAccountToken(token, balance), nil
}