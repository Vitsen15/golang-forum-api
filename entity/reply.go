package entity

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	ThreadID  uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" faker:"-"`
	Thread    Thread         `json:"Thread" faker:"-"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" faker:"-"`
	User      User           `json:"User" faker:"-"`
	Body      string         `gorm:"type:text" faker:"paragraph"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
