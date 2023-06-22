package dao

import (
	"github.com/google/wire"
)

// ProviderSet is dao provider set
var ProviderSet = wire.NewSet(
	NewPermissionDAO,
	NewRoleDAO,
	NewRolePermissionDAO,
	NewUserDAO,
)
