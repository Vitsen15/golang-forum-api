package repository

import (
	"go_forum/main/entity"
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
	err = repository.Db.Delete(&thread).Error
	return
}
