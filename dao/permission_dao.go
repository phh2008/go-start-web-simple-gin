package dao

import (
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/pkg/orm"
	"context"
	"gorm.io/gorm"
)

type PermissionDao struct {
	BaseDao[entity.PermissionEntity]
}

// NewPermissionDAO 创建dao
func NewPermissionDAO(db *gorm.DB) *PermissionDao {
	return &PermissionDao{
		NewBaseDAO[entity.PermissionEntity](db),
	}
}

func (a *PermissionDao) ListPage(ctx context.Context, req model.PermissionListReq) model.PageData[model.PermissionModel] {
	db := a.GetDb(ctx)
	db = db.Model(&entity.PermissionEntity{})
	if req.PermName != "" {
		db = db.Where("perm_name like ?", "%"+req.PermName+"%")
	}
	if req.Url != "" {
		db = db.Where("url=?", req.Url)
	}
	if req.Action != "" {
		db = db.Where("action=?", req.Action)
	}
	if req.PermType != 0 {
		db = db.Where("perm_type=?", req.PermType)
	}
	pageData, db := orm.QueryPageData[model.PermissionModel](db, req.GetPageNo(), req.GetPageSize())
	return pageData
}

func (a *PermissionDao) Add(ctx context.Context, permission entity.PermissionEntity) (entity.PermissionEntity, error) {
	db := a.GetDb(ctx).Create(&permission)
	return permission, db.Error
}

func (a *PermissionDao) Update(ctx context.Context, permission entity.PermissionEntity) (entity.PermissionEntity, error) {
	db := a.GetDb(ctx).Model(&permission).Updates(permission)
	return permission, db.Error
}

func (a *PermissionDao) FindByIdList(idList []int64) []entity.PermissionEntity {
	var list []entity.PermissionEntity
	if len(idList) == 0 {
		return list
	}
	db := a.db
	db.Find(&list, idList)
	return list
}
