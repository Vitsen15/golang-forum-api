package repository

import "gorm.io/gorm"

type Repository struct {
	Db *gorm.DB
}

func CreateRepository(db *gorm.DB) *Repository {
	return &Repository{Db: db}
}
