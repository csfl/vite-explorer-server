package request

type AccountChainBlocklist struct {
	Paging `form:"paging" binding:"required"`

	AccountAddress string `form:"accountAddress"`

	TokenId string `form:"TokenId"`
}

type AccountChainBlock struct {

}