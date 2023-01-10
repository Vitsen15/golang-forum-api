package entity

import (
	"gorm.io/gorm"
	"time"
)

type Reply struct {
	ID        uint `gorm:"primarykey"`
	ThreadID  uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Thread    *Thread
	UserID    uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User
	Body      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
