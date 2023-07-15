package handlers

import (
	"testing"

	openai "github.com/sashabaranov/go-openai"
)

func TestConvertToOpenAIModel(t *testing.T) {
	testCases := []struct {
		name          string
		customModel   string
		expectedModel string
	}{
		{"gpt-4", "gpt-4", openai.GPT4},
		{"gpt-3.5", "gpt-3.5", openai.GPT3Dot5Turbo},
		{"gpt-4-32k", "gpt-4-32k", openai.GPT432K0613},
		{"gpt-3.5-16k", "gpt-3.5-16k", openai.GPT3Dot5Turbo16K0613},
		{"default", "non-existing-model", openai.GPT3Dot5Turbo},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := convertToOpenAIModel(tc.customModel)
			if result != tc.expectedModel {
				t.Errorf("Expected %v, but got %v", tc.expectedModel, result)
			}
		})
	}
}

func TestReplaceModelsWithConvertedToOpenAIModels(t *testing.T) {
	testCases := []struct {
		name          string
		input         RequestBody
		expectedModel string
	}{
		{"gpt-4", RequestBody{Model: "gpt-4"}, openai.GPT4},
		{"gpt-3.5", RequestBody{Model: "gpt-3.5"}, openai.GPT3Dot5Turbo},
		{"gpt-4-32k", RequestBody{Model: "gpt-4-32k"}, openai.GPT432K0613},
		{"gpt-3.5-16k", RequestBody{Model: "gpt-3.5-16k"}, openai.GPT3Dot5Turbo16K0613},
		{"default", RequestBody{Model: "non-existing-model"}, openai.GPT3Dot5Turbo},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceModelsWithConvertedToOpenAIModels(tc.input)
			if result.Model != tc.expectedModel {
				t.Errorf("Expected %v, but got %v", tc.expectedModel, result.Model)
			}
		})
	}
}
