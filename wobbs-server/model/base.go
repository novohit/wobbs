package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID         int32          `gorm:"primaryKey"`
	CreateTime time.Time      `gorm:"column:create_time"`
	UpdateTime time.Time      `gorm:"column:update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreateTime = time.Now()
	b.UpdateTime = b.CreateTime
	return
}
