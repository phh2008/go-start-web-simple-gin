package service

import (
	"com.gientech/selection/dao"
	"com.gientech/selection/entity"
	"com.gientech/selection/model"
	"com.gientech/selection/model/result"
	"com.gientech/selection/pkg/exception"
	"com.gientech/selection/pkg/logger"
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var PermissionSet = wire.NewSet(wire.Struct(new(PermissionService), "*"))

type PermissionService struct {
	PermissionDao *dao.PermissionDao
	Enforcer      *casbin.Enforcer
}

func (a *PermissionService) ListPage(ctx context.Context, req model.PermissionListReq) *result.Result[model.PageData[model.PermissionModel]] {
	data := a.PermissionDao.ListPage(ctx, req)
	return result.Ok[model.PageData[model.PermissionModel]](data)
}

func (a *PermissionService) Add(ctx context.Context, perm model.PermissionModel) *result.Result[entity.PermissionEntity] {
	var permission entity.PermissionEntity
	copier.Copy(&permission, &perm)
	res, err := a.PermissionDao.Add(ctx, permission)
	if err != nil {
		logger.Errorf("添加权限资源失败，%s", err.Error())
		return result.Failure[entity.PermissionEntity]("添加权限资源失败")
	}
	return result.Ok(res)
}

func (a *PermissionService) Update(ctx context.Context, perm model.PermissionModel) *result.Result[*entity.PermissionEntity] {
	oldPerm := a.PermissionDao.GetById(ctx, perm.Id)
	if oldPerm.Id == 0 {
		return result.Error[*entity.PermissionEntity](exception.NotFound)
	}
	var permission entity.PermissionEntity
	copier.Copy(&permission, &perm)
	// 更新权限资源表
	res, err := a.PermissionDao.Update(ctx, permission)
	if err != nil {
		logger.Errorf("更新权限资源失败，%s", err.Error())
		return result.Failure[*entity.PermissionEntity]("更新权限资源失败")
	}
	// 获取角色与资源列表,比如：[[systemAdmin /api/v1/user/list get] [guest /api/v1/user/list get]]
	perms := a.Enforcer.GetFilteredPolicy(1, oldPerm.Url, oldPerm.Action)
	// 更新casbin中的数据
	if len(perms) > 0 {
		for i, v := range perms {
			item := v
			item[1] = res.Url
			item[2] = res.Action
			perms[i] = item
		}
		_, err = a.Enforcer.UpdateFilteredPolicies(perms, 1, oldPerm.Url, oldPerm.Action)
		if err != nil {
			logger.Errorf("更新casbin中的权限错误: %s", err)
		}
	}
	return result.Ok(&res)
}
