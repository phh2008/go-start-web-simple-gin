package dao

import (
	"com.gientech/selection/entity"
	"context"
	"gorm.io/gorm"
)

type RolePermissionDao struct {
	BaseDao[entity.RolePermissionEntity]
}

// NewRolePermissionDAO 创建 dao
func NewRolePermissionDAO(db *gorm.DB) *RolePermissionDao {
	return &RolePermissionDao{
		NewBaseDAO[entity.RolePermissionEntity](db),
	}
}

func (a *RolePermissionDao) DeleteByRoleId(ctx context.Context, roleId int64) error {
	db := a.GetDb(ctx).Where("role_id=?", roleId).Delete(&entity.RolePermissionEntity{})
	return db.Error
}

func (a *RolePermissionDao) BatchAdd(ctx context.Context, list []*entity.RolePermissionEntity) error {
	if len(list) == 0 {
		return nil
	}
	db := a.GetDb(ctx).Create(list)
	return db.Error
}

func (a *RolePermissionDao) ListRoleIdByPermId(ctx context.Context, permId int64) []int64 {
	var roleIds []int64
	a.GetDb(ctx).Model(&entity.RolePermissionEntity{}).
		Where("perm_id=?", permId).
		Pluck("role_id", &roleIds)
	return roleIds
}
