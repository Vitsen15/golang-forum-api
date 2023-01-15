package entity

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	ID        uint           `gorm:"primarykey"  json:"id" faker:"-"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" faker:"-"`
	User      User           `json:"User" faker:"-"`
	Replies   []Reply        `json:"Replies" faker:"-"`
	Title     string         `gorm:"size:255" json:"Title" faker:"sentence"`
	Body      string         `gorm:"type:text" json:"Body" faker:"paragraph"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}
