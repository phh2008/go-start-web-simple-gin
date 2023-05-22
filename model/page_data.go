package model

type PageData[T any] struct {
	Count    int64 `json:"count"`
	PageNo   int   `json:"pageNo"`
	PageSize int   `json:"pageSize"`
	Data     []T   `json:"data"`
}

func NewPageData[T any](pageNo int, pageSize int) *PageData[T] {
	return &PageData[T]{PageNo: pageNo, PageSize: pageSize}
}

func (a *PageData[T]) SetData(data []T) *PageData[T] {
	a.Data = data
	return a
}
