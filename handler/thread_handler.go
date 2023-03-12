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

type ThreadRepository interface {
	Create(*entity.Thread) error
	Update(thread *entity.Thread) error
	Delete(id uint) error
	Get(id uint) (*entity.Thread, error)
	GetAll() ([]*entity.Thread, error)
	GetReplies(id uint) ([]*entity.Reply, error)
}

type ThreadHandler struct {
	repository ThreadRepository
}

func CreateThreadHandler(threadRepository ThreadRepository) *ThreadHandler {
	return &ThreadHandler{repository: threadRepository}
}

func (handler *ThreadHandler) GetThreadById(c *gin.Context) {
	ID, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	thread, getErr := handler.repository.Get(uint(ID))

	if getErr != nil && errors.Is(getErr, gorm.ErrRecordNotFound) {
		log.Println(getErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Thread not found."})
		return
	} else if getErr != nil {
		log.Println(getErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": getErr.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

func (handler *ThreadHandler) GetThreadRepliesById(c *gin.Context) {
	ID, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	replies, searchErr := handler.repository.GetReplies(uint(ID))

	if searchErr != nil {
		log.Println(searchErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": searchErr.Error()})
	}

	c.JSON(http.StatusOK, replies)
}

func (handler *ThreadHandler) GetAllThreads(c *gin.Context) {
	threads, getErr := handler.repository.GetAll()

	if getErr != nil {
		log.Println(getErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't retrieve thread."})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (handler *ThreadHandler) CreateThread(c *gin.Context) {
	thread := entity.Thread{}

	if err := c.ShouldBindJSON(&thread); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createErr := handler.repository.Create(&thread)

	if createErr != nil {
		log.Println(createErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thread created."})
}

func (handler *ThreadHandler) UpdateThread(c *gin.Context) {
	ID, castIDErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castIDErr != nil {
		log.Println(castIDErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castIDErr.Error()})
		return
	}

	thread := entity.Thread{}

	if err := c.ShouldBindJSON(&thread); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	thread.ID = uint(ID)

	updateErr := handler.repository.Update(&thread)

	if updateErr != nil && errors.Is(updateErr, gorm.ErrRecordNotFound) {
		log.Println(updateErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Thread not found."})
		return
	} else if updateErr != nil {
		log.Println(updateErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Unable to update thread."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thread has been updated successfully."})
}

func (handler *ThreadHandler) DeleteThread(c *gin.Context) {
	ID, castIDErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castIDErr != nil {
		log.Println(castIDErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castIDErr.Error()})
		return
	}

	deleteErr := handler.repository.Delete(uint(ID))

	if deleteErr != nil && errors.Is(gorm.ErrRecordNotFound, deleteErr) {
		log.Println(deleteErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": deleteErr.Error()})
		return
	} else if deleteErr != nil {
		log.Println(deleteErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": deleteErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thread has been deleted successfully."})
}
