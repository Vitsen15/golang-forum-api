package seed

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

func CreateThread(db *gorm.DB, thread entity.Thread) error {
	return db.Create(&entity.Thread{
		UserID: thread.UserID,
		Title:  thread.Title,
		Body:   thread.Body,
	}).Error
}
