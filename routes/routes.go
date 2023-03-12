package routes

import (
	"go_forum/main/database"
	"go_forum/main/handler"
	"go_forum/main/middleware"
	"go_forum/main/repository/GORM"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() error {
	router := gin.New()
	dbConnection := database.Connection()

	// Init thread handler and repositories.
	userRepository := GORM.CreateUserRepository(dbConnection)
	threadRepository := GORM.CreateThreadRepository(dbConnection)
	replyRepository := GORM.CreateReplyRepository(dbConnection)

	// Init reply handlers.
	userHandler := handler.CreateAuthHandler(userRepository)
	threadHandler := handler.CreateThreadHandler(threadRepository)
	replyHandler := handler.CreateReplyHandler(replyRepository)

	// Default route to show if API is working.
	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "API server is working!"}) })

	// Authorization routes.
	auth := router.Group("/auth")
	{
		auth.GET("/login", userHandler.Login)
	}

	// Thread routes.
	router.GET("/thread", middleware.AuthMiddleware(), threadHandler.GetAllThreads)
	router.GET("/thread/:id", middleware.AuthMiddleware(), threadHandler.GetThreadById)
	router.GET("/thread/:id/replies", middleware.AuthMiddleware(), threadHandler.GetThreadRepliesById)
	router.POST("/thread", middleware.AuthMiddleware(), threadHandler.CreateThread)
	router.PUT("/thread/:id", middleware.AuthMiddleware(), threadHandler.UpdateThread)
	router.DELETE("/thread/:id", middleware.AuthMiddleware(), threadHandler.DeleteThread)

	//Reply routes.
	router.GET("/reply/:id", middleware.AuthMiddleware(), replyHandler.GetReplyById)
	router.POST("/reply", middleware.AuthMiddleware(), replyHandler.CreateReply)
	router.PUT("/reply/:id", middleware.AuthMiddleware(), replyHandler.UpdateReply)
	router.DELETE("/reply/:id", middleware.AuthMiddleware(), replyHandler.DeleteReply)

	return router.Run(":8080")
}
