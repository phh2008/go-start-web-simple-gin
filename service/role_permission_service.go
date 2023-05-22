package service

import (
	"com.gientech/selection/dao"
	"github.com/google/wire"
)

var RolePermissionSet = wire.NewSet(wire.Struct(new(RolePermissionService), "*"))

type RolePermissionService struct {
	RolePermissionDao *dao.RolePermissionDao
}
