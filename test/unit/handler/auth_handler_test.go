package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go_forum/main/entity"
	authEntity "go_forum/main/entity/auth"
	"go_forum/main/handler"
	authHelper "go_forum/main/helper/auth"
	"go_forum/main/security"
	"go_forum/main/test/mock/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestLogin(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockUserRepository := repository.NewMockUserRepository(mockController)

	// Init user data.
	password := "Temp123$#"
	hash, hashErr := security.HashPassword(password)
	if hashErr != nil {
		t.Logf("hash generation error: %s", hashErr)
		t.Fail()
	}
	user := entity.User{
		ID:        1,
		Email:     "email@domain.com",
		FirstName: "John", LastName: "Doe",
		Hash: hash,
	}
	mockUserRepository.EXPECT().GetByEmail("email@domain.com").Return(&user, nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	authHandler := handler.CreateAuthHandler(mockUserRepository)
	request, _ := http.NewRequest("GET", "/auth/login", nil)
	request.SetBasicAuth(user.Email, password)
	router.GET("/auth/login", authHandler.Login)
	router.ServeHTTP(responseRecorder, request)

	// Generate JWT token.
	credentials := authEntity.Credentials{Email: user.Email, Password: password}
	expirationTime := time.Now().Add(30 * time.Minute)
	tokenString, tokenGenError := authHelper.GenerateJWT(&credentials, expirationTime)
	if tokenGenError != nil {
		t.Logf("JWT token generation error: %s", tokenGenError)
		t.Fail()
	}

	if got := responseRecorder.Body.String(); tokenString != got {
		t.Logf("expected a %s, instead got: %s", tokenString, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}
