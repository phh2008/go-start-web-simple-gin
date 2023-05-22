package model

import (
	"time"
)

type RoleModel struct {
	Id       int64     `json:"id"` // 主键id
	RoleCode string    `json:"roleCode" binding:"required"`
	RoleName string    `json:"roleName" binding:"required"`
	CreateAt time.Time `json:"createAt"` // 创建时间
	UpdateAt time.Time `json:"updateAt"` // 更新时间
	CreateBy int64     `json:"createBy"` // 创建人
	UpdateBy int64     `json:"updateBy"` // 更新人
}

type RoleAssignPermModel struct {
	RoleId     int64   `json:"roleId" binding:"required"`     // 角色ID
	PermIdList []int64 `json:"permIdList" binding:"required"` // 权限ID列表
}
