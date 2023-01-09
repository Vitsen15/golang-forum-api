package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"size:255"`
	FirstName string `gorm:"size:255"`
	LastName  string `gorm:"size:255"`
	Hash      string `gorm:"size:255"`
	Threads   []*Thread
	Replies   []*Reply
}

type Thread struct {
	gorm.Model
	UserID  uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    *User
	Replies []*Reply
	Title   string `gorm:"size:255"`
	Body    string `gorm:"type:text"`
}

type Reply struct {
	gorm.Model
	ThreadID uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Thread   *Thread
	UserID   uint `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User     *User
	Body     string `gorm:"type:text"`
}
