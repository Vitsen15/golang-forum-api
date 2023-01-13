package seed

import (
	"go_forum/main/database/entity"
	"gorm.io/gorm"
)

func CreateReply(db *gorm.DB, reply entity.Reply) error {
	err := db.Create(&entity.Reply{
		UserID:   reply.UserID,
		ThreadID: reply.ThreadID,
		Body:     reply.Body,
	}).Error

	return err
}

func CreateReplies(db *gorm.DB, replies []entity.Reply) error {
	var err error

	for _, reply := range replies {
		err = CreateReply(db, reply)
	}

	return err
}
