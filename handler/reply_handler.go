package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
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
