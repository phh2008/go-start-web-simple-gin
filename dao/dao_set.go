package dao

import (
	"github.com/google/wire"
)

var DaoSet = wire.NewSet(
	BaseDaoSet,
	TestSet,
	PermissionSet,
	RoleSet,
	RolePermissionSet,
	UserSet,
)
