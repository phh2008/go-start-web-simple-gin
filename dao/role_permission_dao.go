package dao

import (
	"com.gientech/selection/entity"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RolePermissionSet = wire.NewSet(wire.Struct(new(RolePermissionDao), "*"))

type RolePermissionDao struct {
	Db *gorm.DB
}

func (a *RolePermissionDao) DeleteByRoleId(ctx context.Context, roleId int64) error {
	db := a.Db.WithContext(ctx)
	db = db.Where("role_id=?", roleId).Delete(&entity.RolePermissionEntity{})
	return db.Error
}

func (a *RolePermissionDao) BatchAdd(ctx context.Context, list []*entity.RolePermissionEntity) error {
	if len(list) == 0 {
		return nil
	}
	db := a.Db.WithContext(ctx)
	db = db.Create(list)
	return db.Error
}
