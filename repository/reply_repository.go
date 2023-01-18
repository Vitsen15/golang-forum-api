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
	if repository.Db.Updates(&reply).RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}
