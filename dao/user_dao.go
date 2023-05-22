package dao

import (
	"com.gientech/selection/entity"
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserSet = wire.NewSet(wire.Struct(new(UserDao), "*"))

type UserDao struct {
	Db *gorm.DB
}

// GetByEmail 根据 email 查询
func (a *UserDao) GetByEmail(ctx context.Context, email string) entity.UserEntity {
	db := a.Db
	var user entity.UserEntity
	db.Model(entity.UserEntity{}).Where("email=?", email).First(&user)
	return user
}

// Add 添加用户
func (a *UserDao) Add(ctx context.Context, user entity.UserEntity) (entity.UserEntity, error) {
	db := a.Db.WithContext(ctx)
	ret := db.Create(&user)
	return user, ret.Error
}

// SetRole 设置角色
func (a *UserDao) SetRole(ctx context.Context, userId int64, role string) error {
	db := a.Db.WithContext(ctx)
	db = db.Model(entity.UserEntity{}).Where("id=?", userId).Update("role_code", role)
	return db.Error
}

func (a *UserDao) DeleteById(ctx context.Context, id int64) error {
	db := a.Db.WithContext(ctx)
	db = db.Delete(&entity.UserEntity{}, id)
	return db.Error
}
