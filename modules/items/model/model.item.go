package model

import (
	"test-pari/constant"
	"test-pari/schemas"
	"time"

	"gorm.io/gorm"
)

type Items struct {
	schemas.FullAudit

	Name        *string `json:"name,omitempty"  gorm:"column:Name"`
	Description *string `json:"description,omitempty"  gorm:"column:Description"`
	Price       *int64  `json:"price,omitempty"  gorm:"column:Price"`
	Qty         *int64  `json:"qty,omitempty"  gorm:"column:Qty"`
	IsActive    *bool   `json:"is_active,omitempty"  gorm:"column:IsActive;default:true"`
}

// ? this is just gorm way of custom table name
func (t *Items) TableName() string {
	return constant.TABLE_ITEMS_NAME
}

func (Items) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Items{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Items) InitAudit(operation string /*, user string, user_id int64*/) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		// t.CreatedByUserName = user
		t.CreatedTime = timeNow
		// t.ModifiedByUserName = user
		// t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_UPDATE:
		// t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_DELETE:
		// t.DeletedByUserId = &user_id
		t.DeletedTime = gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}
