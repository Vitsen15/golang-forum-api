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

func (handler *Handler) GetReplyById(c *gin.Context) {
	id, castErr := strconv.ParseUint(c.Param("id"), 10, 32)

	if castErr != nil {
		log.Println(castErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": castErr.Error()})
		return
	}

	reply, searchErr := handler.Repository.GetReplyById(uint(id))

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

func (handler *Handler) CreateReply(c *gin.Context) {
	reply := entity.Reply{}

	if err := c.ShouldBindJSON(&reply); err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	createErr := handler.Repository.CreateReply(reply)

	if createErr != nil {
		log.Println(createErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": createErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reply created successfully."})
}

func (handler *Handler) UpdateReply(c *gin.Context) {
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

	updateErr := handler.Repository.UpdateReply(reply)

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
