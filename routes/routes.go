package routes

import (
	"go_forum/main/database"
	"go_forum/main/handler"
	"go_forum/main/repository"

	"github.com/gin-gonic/gin"
)

func Start() error {
	var router = gin.New()

	Repository := repository.CreateRepository(database.Connection())
	Handler := handler.CreateHandler(Repository)

	// Default route to show if API is working.
	router.GET("/", Handler.DefaultHandler)

	// Thread routes.
	router.GET("/thread", Handler.GetAllThreads)
	router.GET("/thread/:id", Handler.GetThreadById)
	router.GET("/thread/:id/replies", Handler.GetThreadRepliesById)
	router.POST("/thread", Handler.CreateThread)
	router.PUT("/thread/:id", Handler.UpdateThread)
	router.DELETE("/thread/:id", Handler.DeleteThread)

	return router.Run(":8080")
}
