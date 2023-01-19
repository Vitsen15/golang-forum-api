package repository

import (
	"go_forum/main/entity"

	"gorm.io/gorm"
)

type ThreadRepository interface {
	Create(*entity.Thread) error
	Update(thread *entity.Thread) error
	Delete(id uint) error
	Get(id uint) (*entity.Thread, error)
	GetAll() ([]*entity.Thread, error)
	GetReplies(id uint) ([]*entity.Reply, error)
}

type GORMThreadRepository struct {
	Db *gorm.DB
}

func CreateGORMThreadRepository(db *gorm.DB) *GORMThreadRepository {
	return &GORMThreadRepository{Db: db}
}

func (repository *GORMThreadRepository) Create(thread *entity.Thread) error {
	return repository.Db.Create(thread).Error
}

func (repository *GORMThreadRepository) Update(thread *entity.Thread) error {
	result := repository.Db.Updates(&thread)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *GORMThreadRepository) Delete(id uint) error {
	thread := entity.Thread{ID: id}
	result := repository.Db.Delete(&thread)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *GORMThreadRepository) Get(id uint) (*entity.Thread, error) {
	thread := entity.Thread{ID: id}
	result := repository.Db.Find(&thread)

	if result.Error != nil {
		return nil, result.Error
	}

	return &thread, nil
}

func (repository *GORMThreadRepository) GetAll() (threads []*entity.Thread, err error) {
	err = repository.Db.Find(&threads).Error
	return
}

func (repository *GORMThreadRepository) GetReplies(id uint) (replies []*entity.Reply, err error) {
	err = repository.Db.Find(&replies).Where("thread_id", id).Error
	return
}
