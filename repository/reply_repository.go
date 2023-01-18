package repository

import "go_forum/main/entity"

func (repository *Repository) GetReplyById(id uint) (reply entity.Reply, err error) {
	err = repository.Db.First(&reply, id).Error
	return
}

func (repository *Repository) CreateReply(reply entity.Reply) (err error) {
	err = repository.Db.Create(&reply).Error
	return
}
