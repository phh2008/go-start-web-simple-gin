package dao

import (
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/pkg/orm"
	"context"
	"gorm.io/gorm"
)

type RoleDao struct {
	BaseDao[entity.RoleEntity]
}

// NewRoleDAO 创建 dao
func NewRoleDAO(db *gorm.DB) *RoleDao {
	return &RoleDao{
		NewBaseDAO[entity.RoleEntity](db),
	}
}

func (a *RoleDao) ListPage(ctx context.Context, req model.RoleListReq) model.PageData[model.RoleModel] {
	db := a.GetDb(ctx)
	db = db.Model(&entity.UserEntity{})
	if req.RoleCode != "" {
		db = db.Where("role_code like ?", "%"+req.RoleCode+"%")
	}
	if req.RoleName != "" {
		db = db.Where("role_name like ?", "%"+req.RoleName+"%")
	}
	pageData, db := orm.QueryPageData[model.RoleModel](db, req.GetPageNo(), req.GetPageSize())
	return pageData
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

// DeleteById 删除角色
func (a *RoleDao) DeleteById(ctx context.Context, id int64) error {
	ret := a.GetDb(ctx).Delete(&entity.RoleEntity{}, id)
	return ret.Error
}

// ListByIds 根据角色ID集合查询角色列表
func (a *RoleDao) ListByIds(ctx context.Context, ids []int64) []entity.RoleEntity {
	var list []entity.RoleEntity
	a.GetDb(ctx).Find(&list, ids)
	return list
}
