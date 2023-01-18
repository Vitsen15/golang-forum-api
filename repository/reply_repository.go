package repository

import (
	"go_forum/main/entity"
	"gorm.io/gorm"
)

func (repository *Repository) GetReplyById(id uint) (reply entity.Reply, err error) {
	err = repository.Db.First(&reply, id).Error
	return
}

func (repository *Repository) CreateReply(reply entity.Reply) (err error) {
	err = repository.Db.Create(&reply).Error
	return
}

func (repository *Repository) UpdateReply(reply entity.Reply) (err error) {
	result := repository.Db.Updates(&reply)

	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	} else if result.Error != nil {
		err = result.Error
	}

	return
}

func (repository *Repository) DeleteReplyById(id uint) (err error) {
	reply := entity.Reply{ID: id}
	result := repository.Db.Delete(&reply)

	if result.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	} else if result.Error != nil {
		err = result.Error
	}

	return
}
