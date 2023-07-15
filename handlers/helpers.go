package handlers

import (
	"net/http"

	"github.com/FreddyTheApp/gpt-service/usecases"
	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

var modelMapping = map[string]string{
	"gpt-4":       openai.GPT4,
	"gpt-3.5":     openai.GPT3Dot5Turbo,
	"gpt-4-32k":   openai.GPT432K0613,
	"gpt-3.5-16k": openai.GPT3Dot5Turbo16K0613,
}

func convertToOpenAIModel(customModel string) string {
	openAIModel, ok := modelMapping[customModel]
	if !ok {
		return openai.GPT3Dot5Turbo
	}
	return openAIModel
}

func replaceModelsWithConvertedToOpenAIModels(reqBody RequestBody) RequestBody {
	reqBody.Model = convertToOpenAIModel(reqBody.Model)
	return reqBody
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
