package entity

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      *User
	Replies   []*Reply
	Title     string `gorm:"size:255"`
	Body      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
