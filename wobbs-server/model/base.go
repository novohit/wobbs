package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID         int32          `gorm:"primaryKey" json:"id"`
	CreateTime time.Time      `gorm:"column:create_time" json:"-"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"-"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"-"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreateTime = time.Now()
	b.UpdateTime = b.CreateTime
	return
}
