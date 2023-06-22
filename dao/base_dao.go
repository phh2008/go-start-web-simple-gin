package dao

import (
	"context"
	"gorm.io/gorm"
)

var dbKey = "txDb"

type BaseDao[T any] struct {
	db *gorm.DB
}

// NewBaseDAO 创建 baseDAO
func NewBaseDAO[T any](db *gorm.DB) BaseDao[T] {
	return BaseDao[T]{db: db}
}

// Transaction 开启事务
func (a *BaseDao[T]) Transaction(c context.Context, handler func(tx context.Context) error) error {
	db := a.db
	return db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return handler(context.WithValue(c, dbKey, tx))
	})
}

// GetDb 获取事务的db连接
func (a *BaseDao[T]) GetDb(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		db = a.db
		return db.WithContext(ctx)
	}
	return db
}

// GetById 根据ID查询
func (a *BaseDao[T]) GetById(ctx context.Context, id int64) (T, error) {
	var domain T
	err := a.GetDb(ctx).Limit(1).Find(&domain, id).Error
	return domain, err
}
