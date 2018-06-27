package request

type AccountchainBlocklist struct {
	Paging Paging `form:"paging" binding:"required"`

	AccountAddress string `form:"accountAddress"`

	TokenId string `form:"TokenId"`
}

type AccountchainBlock struct {

}