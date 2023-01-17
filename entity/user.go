package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"ID"`
	Email     string         `gorm:"type:varchar(100);unique_index" json:"Email" faker:"email"`
	FirstName string         `gorm:"size:255" json:"FirstName" faker:"first_name"`
	LastName  string         `gorm:"size:255" json:"LastName" faker:"last_name"`
	Hash      string         `gorm:"size:255" json:"-" faker:"password"`
	Threads   []Thread       `json:"Threads" faker:"-"`
	Replies   []Reply        `json:"Replies" faker:"-"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"-" faker:"-"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"-" faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" faker:"-"`
}
