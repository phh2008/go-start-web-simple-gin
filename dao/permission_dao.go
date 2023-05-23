package dao

import (
	"com.gientech/selection/entity"
	"context"
	"github.com/google/wire"
)

var PermissionSet = wire.NewSet(wire.Struct(new(PermissionDao), "*"))

type PermissionDao struct {
	BaseDao
}

func (a *PermissionDao) Add(ctx context.Context, permission entity.PermissionEntity) (entity.PermissionEntity, error) {
	db := a.GetDb(ctx).Create(&permission)
	return permission, db.Error
}

func (a *PermissionDao) FindByIdList(idList []int64) []entity.PermissionEntity {
	var list []entity.PermissionEntity
	if len(idList) == 0 {
		return list
	}
	db := a.Db
	db.Find(&list, idList)
	return list
}
