package dao

import (
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/pkg/orm"
	"context"
	"gorm.io/gorm"
)

type UserDao struct {
	BaseDao[entity.UserEntity]
}

// NewUserDAO 创建 userDAO
func NewUserDAO(db *gorm.DB) *UserDao {
	return &UserDao{
		NewBaseDAO[entity.UserEntity](db),
	}
}

func (a *UserDao) ListPage(ctx context.Context, req model.UserListReq) model.PageData[model.UserModel] {
	db := a.GetDb(ctx)
	db = db.Model(&entity.UserEntity{})
	if req.RealName != "" {
		db = db.Where("real_name like ?", "%"+req.RealName+"%")
	}
	if req.Email != "" {
		db = db.Where("email like ?", "%"+req.Email+"%")
	}
	if req.Status != 0 {
		db = db.Where("status=?", req.Status)
	}
	var pageData model.PageData[model.UserModel]
	pageData.PageNo = req.GetPageNo()
	pageData.PageSize = req.GetPageSize()
	db.Count(&pageData.Count).
		Scopes(orm.OrderScope(req.Sort, req.Direction), orm.PageScope(req.GetPageNo(), req.GetPageSize())).
		Find(&pageData.Data)
	return pageData
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
	db := a.GetDb(ctx).Model(&entity.UserEntity{}).Where("id=?", userId).Update("role_code", role)
	return db.Error
}

// DeleteById 删除用户
func (a *UserDao) DeleteById(ctx context.Context, id int64) error {
	db := a.GetDb(ctx).Delete(&entity.UserEntity{}, id)
	return db.Error
}

// CancelRole 撤销用户角色
func (a *UserDao) CancelRole(ctx context.Context, roleCode string) error {
	ret := a.GetDb(ctx).Model(&entity.UserEntity{}).Where("role_code=?", roleCode).Update("role_code", "")
	return ret.Error
}
