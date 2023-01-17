package handler

import (
	"go_forum/main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *repository.Repository
}

func CreateHandler(repository *repository.Repository) *Handler {
	return &Handler{Repository: repository}
}

func (handler *Handler) DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API server is working!"})
}
