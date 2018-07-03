package accountchain


import (
	"github.com/vitelabs/go-vite/ledger"
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/common/types"
)


var accountChainAccess = access.GeAccountChainAccess()

func GetBlockByHash (blockHash []byte) (*ledger.AccountBlock, error){
	return accountChainAccess.GetBlockByHash(blockHash)
}

func GetBlockListByAccountAddress (index int, num int, count int, accountAddress *types.Address) ([]*ledger.AccountBlock, error){
	return accountChainAccess.GetBlockListByAccountAddress(index, num, count, accountAddress)
}

func GetBlockListByTokenId (index int, num int, count int, tokenId *types.TokenTypeId) ([]*ledger.AccountBlock, error){
	return accountChainAccess.GetBlockListByTokenId(index, num, count, tokenId)
}

func GetBlockList (index int, num int, count int) ([]*ledger.AccountBlock, error){
	return accountChainAccess.GetBlockList(index, num, count)
}