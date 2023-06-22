package service

import (
	"com.gientech/selection/dao"
)

// RolePermissionService 角色权限关系
type RolePermissionService struct {
	rolePermissionDao *dao.RolePermissionDao
}

// NewRolePermissionService 创建服务
func NewRolePermissionService(rolePermissionDao *dao.RolePermissionDao) *RolePermissionService {
	return &RolePermissionService{
		rolePermissionDao: rolePermissionDao,
	}
}
