// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSceneCase = "scene_case"

// SceneCase mapped from table <scene_case>
type SceneCase struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:Case ID" json:"id"` // Case ID
	ParentID   int32     `gorm:"column:parent_id;not null;comment:父节点id" json:"parent_id"`          // 父节点id
	Type       int32     `gorm:"column:type;not null;comment:case类型" json:"type"`                   // case类型
	SceneID    int32     `gorm:"column:scene_id;not null;comment:所属场景id" json:"scene_id"`           // 所属场景id
	Sort       int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                       // 排序
	Disabled   bool      `gorm:"column:disabled;not null;comment:是否禁用" json:"disabled"`             // 是否禁用
	Extend     string    `gorm:"column:extend;not null" json:"extend"`
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`   // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:最后修改时间" json:"update_time"` // 最后修改时间
	IsDelete   bool      `gorm:"column:is_delete;not null;comment:是否删除" json:"is_delete"`                                 // 是否删除
}

// TableName SceneCase's table name
func (*SceneCase) TableName() string {
	return TableNameSceneCase
}
