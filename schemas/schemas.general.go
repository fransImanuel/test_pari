package schemas

import (
	"time"

	"gorm.io/gorm"
)

type FullAudit struct {
	ID           int64          `gorm:"column:id;primaryKey" json:"id"`
	DeletedTime  gorm.DeletedAt `json:"deleted_time,omitempty" gorm:"column:DeletedTime"`
	CreatedTime  time.Time      `json:"created_time,omitempty" gorm:"column:CreatedTime"`
	ModifiedTime time.Time      `json:"modified_time,omitempty" gorm:"column:ModifiedTime"`
	// DeletedByUserId    *int64         `json:"deleted_by_user_id,omitempty" gorm:"column:DeletedByUserId"`
	// CreatedByUserName  string         `json:"created_by_user_name,omitempty" gorm:"column:CreatedByUserName"`
	// ModifiedByUserName string         `json:"modified_by_user_name,omitempty" gorm:"column:ModifiedByUserName"`
}

type Response struct {
	// Code    int64       `json:"code,omitempty"  `
	Status  string      `json:"status,omitempty"  `
	Message string      `json:"message,omitempty"  `
	Data    interface{} `json:"data,omitempty"  `
}
