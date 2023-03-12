package GORM

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repository *UserRepository) GetByEmail(email string) (user *entity.User, err error) {
	return user, repository.Db.Where("email = ?", email).First(&user).Error
}
