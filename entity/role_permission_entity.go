package entity

type RolePermissionEntity struct {
	Id     int64 // 主键id
	RoleId int64
	PermId int64
}

func (RolePermissionEntity) TableName() string {
	return "sys_role_permission"
}
