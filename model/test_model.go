package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type TestQuery struct {
	QueryPage
	Status int
	Name   string
	UserId int
	Id     int
}

type TestResult struct {
	Id       int
	Name     string
	UserId   int
	Amount   decimal.Decimal
	CreateAt time.Time
	Status   int
}
