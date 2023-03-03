package repository

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*entity.User, error)
}

type GORMUserRepository struct {
	Db *gorm.DB
}

func CreateGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{Db: db}
}

func (repository *GORMUserRepository) GetByEmail(email string) (user *entity.User, err error) {
	return user, repository.Db.Where("email = ?", email).First(&user).Error
}
