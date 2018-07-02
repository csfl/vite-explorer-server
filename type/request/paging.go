package request

type Paging struct {
	Index int `form:"index"`
	Num int `form:"num"`
	Count int `form:"count"`
}

func (paging *Paging) PagingSetDefault () {
	if paging.Num == 0 {
		paging.Num = 1 // The default is one page
	}

	if paging.Count == 0 {
		paging.Count = 10 // The default is ten per page
	}
}