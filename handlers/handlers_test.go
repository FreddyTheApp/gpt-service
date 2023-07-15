package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FreddyTheApp/gpt-service/usecases"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUseCase struct {
	mock.Mock
}

func (m *MockUseCase) Execute(input string, replyOption string, model string, context []usecases.ContextMessage) (usecases.ExecutionResult, error) {
	args := m.Called(input, replyOption, model, context)
	return args.Get(0).(usecases.ExecutionResult), args.Error(1)
}

func TestHandleSimpleReplyRequest(t *testing.T) {
	// Arrange
	mockUseCase := new(MockUseCase)
	h := NewGinHandler(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/simple-reply", h.HandleSimpleReplyRequest)

	json := `{"input": "test input", "model": "gpt-4", "context": [{"IsFromUser": true, "Message": "Hello, world!"}]}`
	req, _ := http.NewRequest("POST", "/simple-reply", bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")

	// Expect that the use case's Execute method will be called with the input from the request
	mockUseCase.On("Execute", "test input", "simple", "gpt-4", mock.Anything).Return(usecases.ExecutionResult{}, nil)

	// Act
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 status code but got %d", resp.Code)
	}

	// Check that the expected methods were called
	mockUseCase.AssertExpectations(t)
}

func TestHandleJokeReplyRequest(t *testing.T) {
	// Arrange
	mockUseCase := new(MockUseCase)
	h := NewGinHandler(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/joke-reply", h.HandleJokeReplyRequest)

	json := `{"input": "test input", "model": "gpt-4", "context": [{"IsFromUser": true, "Message": "Hello, world!"}]}`
	req, _ := http.NewRequest("POST", "/joke-reply", bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")

	// Expect that the use case's Execute method will be called with the input from the request
	mockUseCase.On("Execute", "test input", "joke", "gpt-4", mock.Anything).Return(usecases.ExecutionResult{}, nil)

	// Act
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 status code but got %d", resp.Code)
	}

	// Check that the expected methods were called
	mockUseCase.AssertExpectations(t)
}

func TestHandleTwoSentenceHorrorStoryRURequest(t *testing.T) {
	// Arrange
	mockUseCase := new(MockUseCase)
	h := NewGinHandler(mockUseCase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/horror-story-ru", h.HandleTwoSentenceHorrorStoryRURequest)

	json := `{"input": "test input", "model": "gpt-4", "context": [{"IsFromUser": true, "Message": "Hello, world!"}]}`
	req, _ := http.NewRequest("POST", "/horror-story-ru", bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")

	// Expect that the use case's Execute method will be called with the input from the request
	mockUseCase.On("Execute", "test input", "horror", "gpt-4", mock.Anything).Return(usecases.ExecutionResult{}, nil)

	// Act
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert
	if resp.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 status code but got %d", resp.Code)
	}

	// Check that the expected methods were called
	mockUseCase.AssertExpectations(t)
}
