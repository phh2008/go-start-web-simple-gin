package orm

import (
	"com.gientech/selection/model"
	"com.gientech/selection/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(config *config.Config) *gorm.DB {
	var dsn = config.Viper.GetString("db.url")
	var gdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return gdb
}

func PageScope(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo <= 0 {
			pageNo = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func QueryPageData[T any](db *gorm.DB, pageNo, pageSize int) (model.PageData[T], *gorm.DB) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	var pageData model.PageData[T]
	pageData.PageNo = pageNo
	pageData.PageSize = pageSize
	offset := (pageNo - 1) * pageSize
	newDb := db.Count(&pageData.Count).Offset(offset).Limit(pageSize).Find(&pageData.Data)
	return pageData, newDb
}

func QueryPageData2[T any](db *gorm.DB, page model.QueryPage) (model.PageData[T], *gorm.DB) {
	return QueryPageData[T](db, page.PageNo, page.PageSize)
}
