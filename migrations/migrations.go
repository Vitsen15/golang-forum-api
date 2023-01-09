package migrations

import (
	"go_forum/main/models"
	"gorm.io/gorm"
	"log"
	"os"
)

func AutoMigrate(DB *gorm.DB) {
	err := DB.AutoMigrate(models.User{}, models.Thread{}, models.Reply{})

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
