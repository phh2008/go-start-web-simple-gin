package dao

import (
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/pkg/orm"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var TestSet = wire.NewSet(wire.Struct(new(TestDAO), "*"))

type TestDAO struct {
	Db *gorm.DB
}

func (a *TestDAO) Query(testQuery model.TestQuery) model.PageData[model.TestResult] {
	db := a.Db
	db = db.Model(&entity.Test{})
	if testQuery.Id != 0 {
		db = db.Where("id=?", testQuery.Id)
	}
	if testQuery.Name != "" {
		db = db.Where("name like ?", "%"+testQuery.Name+"%")
	}
	if testQuery.UserId != 0 {
		db = db.Where("user_id=?", testQuery.UserId)
	}
	if testQuery.Status != 0 {
		db = db.Where("status=?", testQuery.Status)
	}
	pageData, db := orm.QueryPageData[model.TestResult](db, testQuery.GetPageNo(), testQuery.GetPageSize())
	return pageData
}
