package repository

import (
	"go_forum/main/entity"

	"gorm.io/gorm"
)

type ReplyRepository interface {
	Create(reply *entity.Reply) error
	Update(reply *entity.Reply) error
	Delete(id uint) error
	Get(id uint) (*entity.Reply, error)
}

type GORMReplyRepository struct {
	Db *gorm.DB
}

func CreateGORMReplyRepository(db *gorm.DB) *GORMReplyRepository {
	return &GORMReplyRepository{Db: db}
}

func (repository *GORMReplyRepository) Create(reply *entity.Reply) error {
	return repository.Db.Create(&reply).Error
}

func (repository *GORMReplyRepository) Update(reply *entity.Reply) error {
	result := repository.Db.Updates(&reply)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *GORMReplyRepository) Delete(id uint) error {
	reply := entity.Reply{ID: id}
	result := repository.Db.Delete(&reply)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *GORMReplyRepository) Get(id uint) (reply *entity.Reply, err error) {
	err = repository.Db.First(&reply, id).Error
	return
}
