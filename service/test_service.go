package service

import (
	"com.gientech/selection/dao"
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"github.com/google/wire"
)

var TestSet = wire.NewSet(wire.Struct(new(TestService), "*"))

type TestService struct {
	TestDao *dao.TestDAO
}

func (a *TestService) Query(testQuery model.TestQuery) *result.Result[model.PageData[model.TestResult]] {
	page := a.TestDao.Query(testQuery)
	return result.Success[model.PageData[model.TestResult]]().SetData(page)
}
