package handler

import (
	"errors"
	"go_forum/main/entity"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReplyRepository interface {
	Create(reply *entity.Reply) error
	Update(reply *entity.Reply) error
	Delete(id uint) error
	Get(id uint) (*entity.Reply, error)
}

type ReplyHandler struct {
	repository ReplyRepository
}

func CreateReplyHandler(repository ReplyRepository) *ReplyHandler {
	return &ReplyHandler{repository: repository}
}

func (handler *ReplyHandler) GetReplyById(c *gin.Context) {
	id, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	reply, searchErr := handler.repository.Get(uint(id))

	if searchErr != nil && errors.Is(searchErr, gorm.ErrRecordNotFound) {
		log.Println(searchErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Thread not found."})
		return
	} else if searchErr != nil {
		log.Println(searchErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": searchErr.Error()})
		return
	}

	c.JSON(http.StatusOK, reply)
}

func (handler *ReplyHandler) CreateReply(c *gin.Context) {
	reply := entity.Reply{}

	if err := c.ShouldBindJSON(&reply); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createErr := handler.repository.Create(&reply)

	if createErr != nil {
		log.Println(createErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reply created successfully."})
}

func (handler *ReplyHandler) UpdateReply(c *gin.Context) {
	id, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	reply := entity.Reply{ID: uint(id)}

	if err := c.ShouldBindJSON(&reply); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updateErr := handler.repository.Update(&reply)

	if updateErr != nil && errors.Is(updateErr, gorm.ErrRecordNotFound) {
		log.Println(updateErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Reply not found."})
		return
	} else if updateErr != nil {
		log.Println(updateErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Unable to update reply."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reply updated successfully."})
}

func (handler *ReplyHandler) DeleteReply(c *gin.Context) {
	id, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	deleteErr := handler.repository.Delete(uint(id))

	if deleteErr != nil && errors.Is(gorm.ErrRecordNotFound, deleteErr) {
		log.Println(deleteErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": deleteErr.Error()})
		return
	} else if deleteErr != nil {
		log.Println(deleteErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": deleteErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reply has been deleted successfully."})
}
