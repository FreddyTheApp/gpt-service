package handlers

import (
	"net/http"

	promptbuilders "github.com/FreddyTheApp/gpt-service/pkg/prompt-builders"
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
	Input        string                    `json:"input"`
	Model        string                    `json:"model"`
	PrevMessages []usecases.ContextMessage `json:"context"`
}

func (h GinHandler) HandleJokeReplyRequest(c *gin.Context) {
	handleReply(c, h.UseCase, promptbuilders.JokeReplyOption)
}

func (h GinHandler) HandleSimpleReplyRequest(c *gin.Context) {
	handleReply(c, h.UseCase, promptbuilders.SimpleReplyOption)
}

func (h GinHandler) HandleTwoSentenceHorrorStoryRURequest(c *gin.Context) {
	handleReply(c, h.UseCase, promptbuilders.TwoSentenceHorrorRU)
}

func handleReply(c *gin.Context, uc usecases.UseCase, replyOption string) {
	var reqBody RequestBody

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reqBody = replaceModelsWithConvertedToOpenAIModels(reqBody)

	responseMessage, err := uc.Execute(reqBody.Input, replyOption, reqBody.Model, reqBody.PrevMessages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": responseMessage})
}
