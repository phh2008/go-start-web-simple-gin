package dao

import (
	"com.gientech/selection/entity"
	"com.gientech/selection/pkg/exception"
	"context"
	"github.com/google/wire"
)

var RoleSet = wire.NewSet(wire.Struct(new(RoleDao), "*"))

type RoleDao struct {
	BaseDao
}

func (a *RoleDao) GetById(id int64) entity.RoleEntity {
	db := a.Db
	var role entity.RoleEntity
	db.First(&role, id)
	return role
}

// Add 添加角色
func (a *RoleDao) Add(ctx context.Context, entity entity.RoleEntity) (entity.RoleEntity, error) {
	// 检查角色是否存在
	role := a.GetByCode(ctx, entity.RoleCode)
	if role.Id > 0 {
		return entity, exception.NewBizError("500", "角色已存在")
	}
	db := a.GetDb(ctx).Create(&entity)
	return entity, db.Error
}

// GetByCode 根据角色编号获取角色
func (a *RoleDao) GetByCode(ctx context.Context, code string) entity.RoleEntity {
	var role entity.RoleEntity
	a.GetDb(ctx).Where("role_code=?", code).First(&role)
	return role
}
