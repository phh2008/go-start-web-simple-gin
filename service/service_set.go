package service

import (
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	TestSet,
	PermissionSet,
	RolePermissionSet,
	RoleSet,
	UserSet,
)
