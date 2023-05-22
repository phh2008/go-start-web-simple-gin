package service

import (
	"com.gientech/selection/dao"
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/logger"
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var PermissionSet = wire.NewSet(wire.Struct(new(PermissionService), "*"))

type PermissionService struct {
	PermissionDao *dao.PermissionDao
}

func (a *PermissionService) Add(ctx context.Context, perm model.PermissionModel) *result.Result[entity.PermissionEntity] {
	var permission entity.PermissionEntity
	copier.Copy(&permission, &perm)
	res, err := a.PermissionDao.Add(ctx, permission)
	if err != nil {
		logger.Errorf("添加权限失败，%s", err.Error())
		return result.Failure[entity.PermissionEntity]("添加权限失败")
	}
	return result.Ok(res)
}
