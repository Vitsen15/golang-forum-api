package entity

import (
	"go_forum/main/security"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" faker:"-"`
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

func (user *User) BeforeCreate(*gorm.DB) (err error) {
	hash, err := security.HashPassword(user.Hash)
	if err != nil {
		log.Println("Couldn't generate password hash for user", err)
		os.Exit(1)
	}
	user.Hash = hash

	return
}
