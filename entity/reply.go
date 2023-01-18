package entity

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        uint           `gorm:"primarykey"`
	ThreadID  uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"ThreadID,string" binding:"required" faker:"-"`
	Thread    Thread         `json:"-" binding:"-" faker:"-"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"UserID,string" binding:"required" faker:"-"`
	User      User           `json:"-" binding:"-" faker:"-"`
	Body      string         `gorm:"type:text" binding:"required" faker:"paragraph"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" faker:"-"`
}
