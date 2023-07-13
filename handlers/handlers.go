package handlers

import (
	"net/http"

	"github.com/FreddyTheApp/gpt-service/usecases"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UseCase usecases.UseCase
}

func NewGinHandler(useCase usecases.UseCase) *GinHandler {
	return &GinHandler{
		UseCase: useCase,
	}
}

type SingleMessageRequestBody struct {
	Message string `json:"message"`
}

func (h GinHandler) HandleJokeReplyRequest(c *gin.Context) {
	handleReplyRequest(c, h.UseCase, usecases.JokeReplyOption)
}

func (h GinHandler) HandleSimpleReplyRequest(c *gin.Context) {
	handleReplyRequest(c, h.UseCase, usecases.SimpleReplyOption)
}

func handleReplyRequest(c *gin.Context, uc usecases.UseCase, replyOption usecases.ReplyOption) {
	var requestMessage SingleMessageRequestBody

	if err := c.ShouldBindJSON(&requestMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := uc.Execute(requestMessage.Message, replyOption)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": responseMessage})
}
