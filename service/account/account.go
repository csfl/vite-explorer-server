package account

import (
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/vite-explorer-server/type/response"
	serviceToken "github.com/vitelabs/vite-explorer-server/service/token"
	serviceAccountChain "github.com/vitelabs/vite-explorer-server/service/accountchain"
	"math/big"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/gin-gonic/gin"
	"github.com/vitelabs/vite-explorer-server/util"
)
var accountAccess = access.GetAccountAccess()


func GetAccount (c *gin.Context, accountAddress *types.Address) (*response.Account, error) {
	account := &ledger.Account{}
	accountMeta, err:= accountAccess.GetAccountMeta(accountAddress)
	if err != nil{
		util.RespondFailed(c, 2, err, "")
		return nil, err
	}
	accountTokenList, err:= GetAccountTokenList(c, accountMeta)
	if err != nil {
		return  nil, err
	}
	return response.NewAccount(accountAddress, account.BlockHeight, accountTokenList), nil
}

func GetAccountTokenList (c *gin.Context, accountMeta *ledger.AccountMeta) ([]*response.AccountToken, error) {
	accountId := accountMeta.AccountId
	var accountTokenList []*response.AccountToken
	for _, accountSimpleToken := range accountMeta.TokenList {
		accountToken, err := GetAccountToken(c, accountSimpleToken.TokenId, accountId,
			accountSimpleToken.LastAccountBlockHeight)
		if err != nil {
			return nil, err
		}
		accountTokenList =  append(accountTokenList, accountToken)
	}
	return accountTokenList, nil
}

func GetAccountToken (c *gin.Context, tokenId *types.TokenTypeId, accountId *big.Int, blockHeight *big.Int) (*response.AccountToken, error) {
	token, gtErr := serviceToken.GetTokenByTokenId(tokenId)
	if gtErr != nil {
		util.RespondFailed(c, 3, gtErr, "")
		return nil, gtErr
	}
	balance, balanceErr := serviceAccountChain.GetAccountBalance(accountId, blockHeight)
	if balanceErr != nil {
		util.RespondFailed(c, 4, balanceErr, "")
		return nil, balanceErr
	}
	return response.NewAccountToken(token, balance), nil
}