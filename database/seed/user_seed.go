package seed

import (
	"go_forum/main/database/entity"
	"gorm.io/gorm"
)

// CreateUser is seeder to create user.
func CreateUser(db *gorm.DB, user entity.User) error {
	return db.Create(&entity.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Hash:      user.Hash,
	}).Error
}

func CreateUsers(db *gorm.DB, users []entity.User) error {
	var err error

	for _, user := range users {
		err = CreateUser(db, user)
	}

	return err
}
