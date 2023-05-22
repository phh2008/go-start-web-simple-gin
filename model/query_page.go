package model

type QueryPage struct {
	PageNo   int `json:"pageNo" form:"pageNo"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

func NewQueryPage(pageNo int, pageSize int) *QueryPage {
	return &QueryPage{PageNo: pageNo, PageSize: pageSize}
}

func (a QueryPage) GetPageNo() int {
	if a.PageNo <= 0 {
		return 1
	}
	return a.PageNo
}

func (a QueryPage) GetPageSize() int {
	if a.PageSize <= 0 {
		return 10
	}
	return a.PageSize
}
