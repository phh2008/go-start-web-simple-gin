package dao

import (
	"com.gientech/selection/entity"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var RolePermissionSet = wire.NewSet(wire.Struct(new(RolePermissionDao), "*"))

type RolePermissionDao struct {
	Db *gorm.DB
}

func (a *RolePermissionDao) DeleteByRoleId(roleId int64) error {
	db := a.Db
	db = db.Where("role_id=?", roleId).Delete(&entity.RolePermissionEntity{})
	return db.Error
}

func (a *RolePermissionDao) BatchAdd(list []*entity.RolePermissionEntity) error {
	if len(list) == 0 {
		return nil
	}
	db := a.Db
	db = db.Create(list)
	return db.Error
}
