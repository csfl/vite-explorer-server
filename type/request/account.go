package request

type AccountDetail struct {
	AccountAddress string `form:"accountAddress" binding:"required"`
}

type AccountNewTestToken struct {
	AccountAddress string `form:"accountAddress" binding:"required"`
}