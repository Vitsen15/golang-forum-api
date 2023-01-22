package handler

import (
	"go_forum/main/entity"
	"go_forum/main/repository/mock_repository"
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

func TestGetThreadById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Get(uint(1)).Return(
		&entity.Thread{
			ID:     1,
			UserID: 1,
			Title:  "Thread title",
			Body:   "Thread body",
		}, nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("GET", "/thread/1", nil)

	router.GET("/thread/:id", threadHandler.GetThreadById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetThreadByIdBadId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("GET", "/thread/bad_id", nil)

	router.GET("/thread/:id", threadHandler.GetThreadById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetThreadByIdNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Get(uint(1)).Return(nil, gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("GET", "/thread/1", nil)

	router.GET("/thread/:id", threadHandler.GetThreadById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetThreadRepliesById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	replies := []*entity.Reply{{ID: 1, ThreadID: 1, Body: "Reply body 1."}, {ID: 1, ThreadID: 1, Body: "Reply body 2."}}
	mock.EXPECT().GetReplies(uint(1)).Return(replies, nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("GET", "/thread/1/replies", nil)

	router.GET("/thread/:id/replies", threadHandler.GetThreadRepliesById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGetThreadRepliesByIdBadId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("GET", "/thread/bad_id/replies", nil)

	router.GET("/thread/:id/replies", threadHandler.GetThreadRepliesById)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestGeAllThreads(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	threads := []*entity.Thread{{ID: 1, UserID: 1, Title: "Thead 1 title"}, {ID: 2, UserID: 2, Title: "Thead 2 title"}}
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().GetAll().Return(threads, nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("GET", "/thread", nil)

	router.GET("/thread", threadHandler.GetAllThreads)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestCreateThread(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Create(&entity.Thread{
		UserID: 1,
		Title:  "Thread title",
		Body:   "Thread body",
	}).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("POST", "/thread", strings.NewReader(`{"UserID": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.POST("/thread", threadHandler.CreateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusCreated, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestCreateThreadCouldntBindJson(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("POST", "/thread", strings.NewReader(`{"WrongField": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.POST("/thread", threadHandler.CreateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateThread(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Update(&entity.Thread{
		ID:     1,
		UserID: 1,
		Title:  "Thread title",
		Body:   "Thread body",
	}).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("PUT", "/thread/1", strings.NewReader(`{"UserID": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.PUT("/thread/:id", threadHandler.UpdateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateThreadBadId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("PUT", "/thread/bad_id", strings.NewReader(`{"UserID": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.PUT("/thread/:id", threadHandler.UpdateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateThreadCouldntbindJSON(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("PUT", "/thread/1", strings.NewReader(`{"WrongField": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.PUT("/thread/:id", threadHandler.UpdateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestUpdateThreadNonFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Update(&entity.Thread{
		ID:     404,
		UserID: 1,
		Title:  "Thread title",
		Body:   "Thread body",
	}).Return(gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("PUT", "/thread/404", strings.NewReader(`{"UserID": "1", "Title": "Thread title", "Body": "Thread body"}`))

	router.PUT("/thread/:id", threadHandler.UpdateThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteThread(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Delete(uint(1)).Return(nil).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("DELETE", "/thread/1", nil)

	router.DELETE("/thread/:id", threadHandler.DeleteThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusOK, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteThreadBadId(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock_repository.NewMockThreadRepository(controller))
	request, _ := http.NewRequest("DELETE", "/thread/bad_id", nil)

	router.DELETE("/thread/:id", threadHandler.DeleteThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusBadRequest, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}

func TestDeleteThreadNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mock := mock_repository.NewMockThreadRepository(controller)
	mock.EXPECT().Delete(uint(1)).Return(gorm.ErrRecordNotFound).Times(1)

	router := gin.Default()
	responseRecorder := httptest.NewRecorder()
	threadHandler := CreateThreadHandler(mock)
	request, _ := http.NewRequest("DELETE", "/thread/1", nil)

	router.DELETE("/thread/:id", threadHandler.DeleteThread)
	router.ServeHTTP(responseRecorder, request)

	if want, got := http.StatusNotFound, responseRecorder.Result().StatusCode; got != want {
		t.Logf("expected a %d, instead got: %d", want, got)
		t.Logf("response: %s", responseRecorder.Body.String())
		t.Fail()
	}
}
