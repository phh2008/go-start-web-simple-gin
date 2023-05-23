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

var RoleSet = wire.NewSet(wire.Struct(new(RoleService), "*"))

type RoleService struct {
	RoleDao           *dao.RoleDao
	RolePermissionDao *dao.RolePermissionDao
	PermissionDao     *dao.PermissionDao
	Enforcer          *casbin.Enforcer
	UserDao           *dao.UserDao
}

func (a *RoleService) ListPage(ctx context.Context, req model.RoleListReq) *result.Result[model.PageData[model.RoleModel]] {
	data := a.RoleDao.ListPage(ctx, req)
	return result.Ok[model.PageData[model.RoleModel]](data)
}

// Add 添加角色
func (a *RoleService) Add(ctx context.Context, role model.RoleModel) *result.Result[entity.RoleEntity] {
	var roleEntity entity.RoleEntity
	copier.Copy(&roleEntity, &role)
	roleEntity, err := a.RoleDao.Add(ctx, roleEntity)
	if err != nil {
		return result.Error[entity.RoleEntity](err)
	}
	return result.Ok[entity.RoleEntity](roleEntity)
}

// GetByCode 根据角色编号获取角色
func (a *RoleService) GetByCode(ctx context.Context, roleCode string) *result.Result[entity.RoleEntity] {
	role := a.RoleDao.GetByCode(ctx, roleCode)
	return result.Ok[entity.RoleEntity](role)
}

func (a *RoleService) AssignPermission(ctx context.Context, assign model.RoleAssignPermModel) *result.Result[any] {
	// 获取角色
	role := a.RoleDao.GetById(assign.RoleId)
	if role.Id == 0 {
		return result.Failure[any]("角色不存在")
	}
	// 获取权限列表(权限列表为空表示要清空权限)
	permList := a.PermissionDao.FindByIdList(assign.PermIdList)
	// 构建角色与权限关系对象
	var rolePermList []*entity.RolePermissionEntity // 角色权限关系表数据
	var permissions [][]string                      // casbin 中的角色与权限数据
	for _, v := range permList {
		rolePermList = append(rolePermList, &entity.RolePermissionEntity{PermId: v.Id, RoleId: assign.RoleId})
		permissions = append(permissions, []string{role.RoleCode, v.Url, v.Action})
	}
	err := a.RolePermissionDao.Transaction(ctx, func(tx context.Context) error {
		// 删除原来的角色与权限关系
		err := a.RolePermissionDao.DeleteByRoleId(tx, assign.RoleId)
		if err != nil {
			logger.Errorf("删除操作失败，%s", err.Error())
			return exception.SysError
		}
		// 新增角色与权限关系
		err = a.RolePermissionDao.BatchAdd(tx, rolePermList)
		if err != nil {
			logger.Errorf("删除操作失败，%s", err.Error())
			return exception.SysError
		}
		return err
	})
	if err != nil {
		return result.Error[any](err)
	}
	// 更新casbin中的角色与资源关系
	a.Enforcer.DeletePermissionsForUser(role.RoleCode)
	if len(permissions) > 0 {
		a.Enforcer.AddPolicies(permissions)
	}
	return result.Success[any]()
}

// DeleteById 删除角色
func (a *RoleService) DeleteById(ctx context.Context, id int64) *result.Result[any] {
	role := a.RoleDao.GetById(id)
	if role.Id == 0 {
		return result.Success[any]()
	}
	err := a.RoleDao.Transaction(ctx, func(tx context.Context) error {
		// 清空用户表上的角色字段
		err := a.UserDao.CancelRole(ctx, role.RoleCode)
		if err != nil {
			return err
		}
		// 删除角色
		err = a.RoleDao.DeleteById(ctx, id)
		if err != nil {
			return err
		}
		// 删除角色与资源关系
		err = a.RolePermissionDao.DeleteByRoleId(ctx, id)
		if err != nil {
			return err
		}
		// 删除 casbin 中的角色与资源关系、角色与用户关系
		_, err = a.Enforcer.DeleteRole(role.RoleCode)
		return err
	})
	if err != nil {
		logger.Errorf("删除角色操作失败，%s", err)
		return result.Error[any](exception.SysError)
	}
	return result.Success[any]()
}
