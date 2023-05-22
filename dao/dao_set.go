package dao

import (
	"github.com/google/wire"
)

var DaoSet = wire.NewSet(
	TestSet,
	PermissionSet,
	RoleSet,
	RolePermissionSet,
	UserSet,
)
