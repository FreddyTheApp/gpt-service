package handlers

import (
	"net/http"
	"os"

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

type RequestBody struct {
	Message      string             `json:"message"`
	PrevMessages []usecases.Message `json:"prev-messages"`
}

func (h GinHandler) HandleJokeReplyRequest(c *gin.Context) {
	handleReplyRequest(c, h.UseCase, usecases.JokeReplyOption)
}

func (h GinHandler) HandleSimpleReplyRequest(c *gin.Context) {
	handleReplyRequest(c, h.UseCase, usecases.SimpleReplyOption)
}

func (h GinHandler) HandleTwoSentenceHorrorStoryRURequest(c *gin.Context) {
	handleReplyRequest(c, h.UseCase, usecases.TwoSentenceHorrorRU)
}

func (h GinHandler) SecretMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestSecret := c.GetHeader("X-Secret")

		if requestSecret == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No secret in the header"})
			c.Abort()
			return
		}

		envSecret := os.Getenv("API_SECRET")

		if requestSecret != envSecret {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid secret"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func handleReplyRequest(c *gin.Context, uc usecases.UseCase, replyOption usecases.ReplyOption) {
	var reqBody RequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := uc.Execute(reqBody.Message, replyOption, reqBody.PrevMessages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": responseMessage})
}
