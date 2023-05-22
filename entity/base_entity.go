package entity

import (
	"time"
)

type BaseEntity struct {
	Id       int64     `json:"id"`                             // 主键id
	CreateAt time.Time `gorm:"autoCreateTime" json:"createAt"` // 创建时间
	UpdateAt time.Time `gorm:"autoUpdateTime" json:"updateAt"` // 更新时间
	CreateBy int64     `json:"createBy"`                       // 创建人
	UpdateBy int64     `json:"updateBy"`                       // 更新人
	Deleted  uint8     `gorm:"default:1" json:"deleted"`       // 是否删除 1-否，2-是
}
