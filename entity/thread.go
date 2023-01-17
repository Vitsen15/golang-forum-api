package entity

import (
	"gorm.io/gorm"
	"time"
)

type Thread struct {
	ID        uint           `gorm:"primarykey" json:"ID" faker:"-"`
	UserID    uint           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"UserID,string" binding:"required" faker:"-"`
	User      User           `json:"-" faker:"-"`
	Replies   []Reply        `json:"-" faker:"-"`
	Title     string         `gorm:"size:255" json:"Title" binding:"required" faker:"sentence"`
	Body      string         `gorm:"type:text" json:"Body" binding:"required" faker:"paragraph"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"-" faker:"-"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"-" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" faker:"-"`
}
