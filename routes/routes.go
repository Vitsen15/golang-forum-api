package routes

import (
	"go_forum/main/database"
	"go_forum/main/handler"
	"go_forum/main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.New()
	dbConnection := database.Connection()

	// Init thread handler and repository.
	threadRepository := repository.CreateGORMThreadRepository(dbConnection)
	threadHandler := handler.CreateThreadHandler(threadRepository)

	// Init reply handler and repository.
	replyRepository := repository.CreateGORMReplyRepository(dbConnection)
	replyHandler := handler.CreateReplyHandler(replyRepository)

	// Default route to show if API is working.
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "API server is working!"}) })

	// Thread routes.
	router.GET("/thread", threadHandler.GetAllThreads)
	router.GET("/thread/:id", threadHandler.GetThreadById)
	router.GET("/thread/:id/replies", threadHandler.GetThreadRepliesById)
	router.POST("/thread", threadHandler.CreateThread)
	router.PUT("/thread/:id", threadHandler.UpdateThread)
	router.DELETE("/thread/:id", threadHandler.DeleteThread)

	//Reply routes.
	router.GET("/reply/:id", replyHandler.GetReplyById)
	router.POST("/reply", replyHandler.CreateReply)
	router.PUT("/reply/:id", replyHandler.UpdateReply)
	router.DELETE("/reply/:id", replyHandler.DeleteReply)

	return router.Run(":8080")
}
