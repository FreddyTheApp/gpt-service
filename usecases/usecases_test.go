package usecases

import (
	"testing"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
)

// MockGenerator mocks the Generator interface
type MockGenerator struct {
	mock.Mock
}

func (m *MockGenerator) Generate(prompt string, model string, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	args := m.Called(prompt, model, messages)
	return args.Get(0).(openai.ChatCompletionResponse), args.Error(1)
}

// MockPromptBuilder mocks the PromptBuilder interface
type MockPromptBuilder struct {
	mock.Mock
}

func (m *MockPromptBuilder) Build(input string, replyOption string) string {
	args := m.Called(input, replyOption)
	return args.String(0)
}

func TestExecute(t *testing.T) {
	generator := new(MockGenerator)
	promptBuilder := new(MockPromptBuilder)
	useCase := NewGenerateTextUseCase(generator, promptBuilder)

	// set up expectations
	promptBuilder.On("Build", "Hello", "Reply").Return("prompt")
	generator.On("Generate", "prompt", "model", []openai.ChatCompletionMessage{}).Return(
		openai.ChatCompletionResponse{
			Choices: []openai.ChatCompletionChoice{
				{Message: openai.ChatCompletionMessage{Content: "response"}},
			},
			Model: "model",
			Usage: openai.Usage{TotalTokens: 10, PromptTokens: 5, CompletionTokens: 5},
		}, nil)

	// run the function
	result, err := useCase.Execute("Hello", "Reply", "model", []ContextMessage{})

	// assertions
	if err != nil {
		t.Error("Error should be nil")
	}

	expected := ExecutionResult{
		Content:          "response",
		Model:            "model",
		TokenUsage:       10,
		TokenUsageInput:  5,
		TokenUsageOutput: 5,
	}

	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// check that the expectations were met
	mock.AssertExpectationsForObjects(t, promptBuilder, generator)
}
