package dao

import (
	"com.gientech/selection/entity"
	"context"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

type UserDao struct {
	BaseDao
}

// GetByEmail 根据 email 查询
func (a *UserDao) GetByEmail(ctx context.Context, email string) entity.UserEntity {
	var user entity.UserEntity
	a.GetDb(ctx).Model(entity.UserEntity{}).Where("email=?", email).First(&user)
	return user
}

// Add 添加用户
func (a *UserDao) Add(ctx context.Context, user entity.UserEntity) (entity.UserEntity, error) {
	ret := a.GetDb(ctx).Create(&user)
	return user, ret.Error
}

// SetRole 设置角色
func (a *UserDao) SetRole(ctx context.Context, userId int64, role string) error {
	db := a.GetDb(ctx).Model(entity.UserEntity{}).Where("id=?", userId).Update("role_code", role)
	return db.Error
}

func (a *UserDao) DeleteById(ctx context.Context, id int64) error {
	db := a.GetDb(ctx).Delete(&entity.UserEntity{}, id)
	return db.Error
}
