package repository

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

func (repository *Repository) GetThreadById(id uint) (thread entity.Thread, err error) {
	err = repository.Db.First(&thread, id).Error
	return
}

func (repository *Repository) GetAllThreads() (thread []entity.Thread, err error) {
	err = repository.Db.Find(&thread).Error
	return
}

func (repository *Repository) CreateThread(thread entity.Thread) (err error) {
	err = repository.Db.Create(&thread).Error
	return
}

func (repository *Repository) UpdateThread(thread entity.Thread) (err error) {
	err = repository.Db.Updates(&thread).Error
	return
}

func (repository *Repository) DeleteThreadById(id uint) (err error) {
	thread := entity.Thread{ID: id}

	if repository.Db.Delete(&thread).RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}
