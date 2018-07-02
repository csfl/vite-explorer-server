package request

type TokenList struct {
	Paging
}

type TokenDetail struct {
	TokenId string `form:"tokenId"`
	TokenName string `form:"tokenName"`
	TokenSymbol string `form:"tokenSymbol"`
}