package handler

import (
	"go_forum/main/entity"
	"go_forum/main/handler"
	"go_forum/main/test/mock/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGetReplyById(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	reply := entity.Reply{ID: 1, UserID: 1, Body: "Reply body"}
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Get(uint(1)).Return(&reply, nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("GET", "/reply/1", nil)

	router.GET("/reply/:id", replyHandler.GetReplyById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetReplyByIdWithBadId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(repository.NewMockReplyRepository(mockController))
	request, _ := http.NewRequest("GET", "/reply/bad_id", nil)

	router.GET("/reply/:id", replyHandler.GetReplyById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetReplyByIdNotFound(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Get(uint(228)).Return(nil, gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("GET", "/reply/228", nil)

	router.GET("/reply/:id", replyHandler.GetReplyById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestCreateReply(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Create(&entity.Reply{ThreadID: 1, UserID: 1, Body: "New reply body"}).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("POST", "/reply", strings.NewReader(`{"ThreadID" : "1", "UserID" : "1", "Body" : "New reply body"}`))

	router.POST("/reply", replyHandler.CreateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusCreated, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestCreateReplyCouldntBindJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(repository.NewMockReplyRepository(mockController))
	request, _ := http.NewRequest("POST", "/reply", strings.NewReader(`{"WrongField" : "1", "UserID" : "1", "Body" : "New reply body"}`))

	router.POST("/reply", replyHandler.CreateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateReply(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Update(&entity.Reply{ID: 1, ThreadID: 1, UserID: 1, Body: "Updated reply body"}).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("PUT", "/reply/1", strings.NewReader(`{"ThreadID" : "1", "UserID" : "1", "Body" : "Updated reply body"}`))

	router.PUT("/reply/:id", replyHandler.UpdateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateReplyBadId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(repository.NewMockReplyRepository(mockController))
	request, _ := http.NewRequest("PUT", "/reply/bad_id", strings.NewReader(`{"ThreadID" : "1", "UserID" : "1", "Body" : "Updated reply body"}`))

	router.PUT("/reply/:id", replyHandler.UpdateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateReplyNotFound(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Update(&entity.Reply{ID: 228, ThreadID: 1, UserID: 1, Body: "Updated reply body"}).Return(gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("PUT", "/reply/228", strings.NewReader(`{"ThreadID" : "1", "UserID" : "1", "Body" : "Updated reply body"}`))

	router.PUT("/reply/:id", replyHandler.UpdateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateReplyCouldntBindJSON(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(repository.NewMockReplyRepository(mockController))
	request, _ := http.NewRequest("PUT", "/reply/228", strings.NewReader(`{"WrongField" : "1", "UserID" : "1", "Body" : "Updated reply body"}`))

	router.PUT("/reply/:id", replyHandler.UpdateReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteReply(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Delete(uint(1)).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("DELETE", "/reply/1", nil)

	router.DELETE("/reply/:id", replyHandler.DeleteReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteReplyBadId(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(repository.NewMockReplyRepository(mockController))
	request, _ := http.NewRequest("DELETE", "/reply/bad_id", nil)

	router.DELETE("/reply/:id", replyHandler.DeleteReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteReplyNotFound(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mock := repository.NewMockReplyRepository(mockController)
	mock.EXPECT().Delete(uint(404)).Return(gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	replyHandler := handler.CreateReplyHandler(mock)
	request, _ := http.NewRequest("DELETE", "/reply/404", nil)

	router.DELETE("/reply/:id", replyHandler.DeleteReply)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; want != got {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}
