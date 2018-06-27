package request

type Paging struct {
	Index int `form:"index" binding:"required"`
	Num int `form:"num"`
	Count int `form:"count"`
}
