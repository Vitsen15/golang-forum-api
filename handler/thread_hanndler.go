package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go_forum/main/entity"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (handler *Handler) GetThreadById(c *gin.Context) {
	ID, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	thread, getErr := handler.Repository.GetThreadById(uint(ID))

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

func (handler *Handler) GetThreadRepliesById(c *gin.Context) {
	ID, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	replies, searchErr := handler.Repository.GetThreadRepliesById(uint(ID))

	if searchErr != nil {
		log.Println(searchErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": searchErr.Error()})
	}

	c.JSON(http.StatusOK, replies)
}

func (handler *Handler) GetAllThreads(c *gin.Context) {
	threads, getErr := handler.Repository.GetAllThreads()

	if getErr != nil {
		log.Println(getErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Couldn't retrieve thread."})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (handler *Handler) CreateThread(c *gin.Context) {
	thread := entity.Thread{}

	if err := c.ShouldBindJSON(&thread); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createErr := handler.Repository.CreateThread(thread)

	if createErr != nil {
		log.Println(createErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thread created."})
}

func (handler *Handler) UpdateThread(c *gin.Context) {
	ID, castIDErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castIDErr != nil {
		log.Println(castIDErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castIDErr.Error()})
		return
	}

	thread := entity.Thread{}

	if err := c.ShouldBindJSON(&thread); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	thread.ID = uint(ID)

	updateErr := handler.Repository.UpdateThread(thread)

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

func (handler *Handler) DeleteThread(c *gin.Context) {
	ID, castIDErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castIDErr != nil {
		log.Println(castIDErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castIDErr.Error()})
		return
	}

	deleteErr := handler.Repository.DeleteThreadById(uint(ID))

	if deleteErr != nil && errors.Is(gorm.ErrRecordNotFound, deleteErr) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": deleteErr.Error()})
		return
	} else if deleteErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": deleteErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thread has been deleted successfully."})
}
