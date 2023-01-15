package migrations

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
	"log"
	"os"
)

func AutoMigrate(DB *gorm.DB) {
	err := DB.AutoMigrate(entity.User{}, entity.Thread{}, entity.Reply{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
