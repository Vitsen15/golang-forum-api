package GORM

import (
	"go_forum/main/entity"

	"gorm.io/gorm"
)

type ReplyRepository struct {
	Db *gorm.DB
}

func CreateReplyRepository(db *gorm.DB) *ReplyRepository {
	return &ReplyRepository{Db: db}
}

func (repository *ReplyRepository) Create(reply *entity.Reply) error {
	return repository.Db.Create(&reply).Error
}

func (repository *ReplyRepository) Update(reply *entity.Reply) error {
	result := repository.Db.Updates(&reply)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ReplyRepository) Delete(id uint) error {
	reply := entity.Reply{ID: id}
	result := repository.Db.Delete(&reply)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ReplyRepository) Get(id uint) (reply *entity.Reply, err error) {
	err = repository.Db.First(&reply, id).Error
	return
}
