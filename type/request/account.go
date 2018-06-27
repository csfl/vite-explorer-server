package request

type AccountDetail struct {
	AccountAddress string `form:"accountAddress" binding:"required"`
}