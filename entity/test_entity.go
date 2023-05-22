package entity

import (
	"com.gientech/selection/pkg/localtime"
	"github.com/shopspring/decimal"
)

type Test struct {
	Id       int
	Name     string
	Amount   decimal.Decimal
	CreateAt localtime.LocalTime
	Status   int
	UserId   int
}

func (Test) TableName() string {
	return "test"
}
