package GORM

import (
	"go_forum/main/entity"

	"gorm.io/gorm"
)

type ThreadRepository struct {
	Db *gorm.DB
}

func CreateThreadRepository(db *gorm.DB) *ThreadRepository {
	return &ThreadRepository{Db: db}
}

func (repository *ThreadRepository) Create(thread *entity.Thread) error {
	return repository.Db.Create(thread).Error
}

func (repository *ThreadRepository) Update(thread *entity.Thread) error {
	result := repository.Db.Updates(&thread)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ThreadRepository) Delete(id uint) error {
	thread := entity.Thread{ID: id}
	result := repository.Db.Delete(&thread)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ThreadRepository) Get(id uint) (*entity.Thread, error) {
	thread := entity.Thread{ID: id}
	result := repository.Db.Find(&thread)

	if result.Error != nil {
		return nil, result.Error
	}

	return &thread, nil
}

func (repository *ThreadRepository) GetAll() (threads []*entity.Thread, err error) {
	err = repository.Db.Find(&threads).Error
	return
}

func (repository *ThreadRepository) GetReplies(id uint) (replies []*entity.Reply, err error) {
	err = repository.Db.Find(&replies).Where("thread_id", id).Error
	return
}
