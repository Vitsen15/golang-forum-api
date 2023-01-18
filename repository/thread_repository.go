package repository

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

func (repository *Repository) GetThreadById(id uint) (thread entity.Thread, err error) {
	err = repository.Db.First(&thread, id).Error
	return
}

func (repository *Repository) GetThreadRepliesById(id uint) (replies []entity.Reply, err error) {
	err = repository.Db.Find(&replies).Where("thread_id", id).Error
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
	result := repository.Db.Updates(&thread)

	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	} else if result.Error != nil {
		err = result.Error
	}

	return
}

func (repository *Repository) DeleteThreadById(id uint) (err error) {
	thread := entity.Thread{ID: id}
	result := repository.Db.Delete(&thread)

	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	} else if result.Error != nil {
		err = result.Error
	}

	return
}
