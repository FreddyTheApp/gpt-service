package handlers

import (
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
