package dao

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var BaseDaoSet = wire.NewSet(wire.Struct(new(BaseDao), "*"))

var dbKey = "txDb"

type BaseDao struct {
	Db *gorm.DB
}

// Transaction 开启事务
func (a *BaseDao) Transaction(c context.Context, handler func(tx context.Context) error) error {
	db := a.Db
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return handler(context.WithValue(c, dbKey, tx))
	})
}

// GetDb 获取事务的db连接
func (a *BaseDao) GetDb(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		db = a.Db
		return db.WithContext(ctx)
	}
	return db
}
