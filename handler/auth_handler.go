package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_forum/main/entity"
	authEntity "go_forum/main/entity/auth"
	authHelper "go_forum/main/helper/auth"
	"go_forum/main/security"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type UserRepository interface {
	GetByEmail(email string) (*entity.User, error)
}

type AuthHandler struct {
	userRepository UserRepository
}

func CreateAuthHandler(userRepository UserRepository) *AuthHandler {
	return &AuthHandler{userRepository: userRepository}
}

func (handler *AuthHandler) Login(c *gin.Context) {
	//Retrieve credentials from user input.
	email, password, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Please provide a username and password"})
		return
	}

	credentials := authEntity.Credentials{Email: email, Password: password}

	// Identify user.
	user, userGetErr := handler.userRepository.GetByEmail(credentials.Email)
	if userGetErr != nil && errors.Is(userGetErr, gorm.ErrRecordNotFound) {
		log.Println(userGetErr)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User doesn't exist."})
		return
	} else if userGetErr != nil {
		log.Println(userGetErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Authentication failed."})
		return
	}

	// Authenticate user.
	if !security.DoPasswordsMatch(user.Hash, password) {
		message := fmt.Sprintf("Authentication failed for user: %s.", user.Email)
		log.Println(message)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": message})
	}

	// Set JWT token and cookie expiration time to 30 minutes.
	expirationTime := time.Now().Add(30 * time.Minute)
	tokenString, tokenGenError := authHelper.GenerateJWT(&credentials, expirationTime)
	if tokenGenError != nil {
		log.Println(tokenGenError)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate JWT."})
		return
	}

	_, _ = c.Writer.Write([]byte(tokenString))
}
