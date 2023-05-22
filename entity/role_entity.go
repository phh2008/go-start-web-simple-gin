package entity

type RoleEntity struct {
	BaseEntity
	RoleCode string
	RoleName string
}

func (RoleEntity) TableName() string {
	return "sys_role"
}
