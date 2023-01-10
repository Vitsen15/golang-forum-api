package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primarykey"`
	Email     string `gorm:"size:255"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Hash      string `gorm:"size:255"`
	Threads   []*Thread
	Replies   []*Reply
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
