package usecases

import (
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestConvertMessagesToChatCompletionMessages(t *testing.T) {
	messages := []ContextMessage{
		{IsFromUser: true, Message: "Hello!"},
		{IsFromUser: false, Message: "Hi, how can I help you?"},
	}

	expected := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "Hello!"},
		{Role: openai.ChatMessageRoleAssistant, Content: "Hi, how can I help you?"},
	}

	result := convertMessagesToChatCompletionMessages(messages)

	if len(result) != len(expected) {
		t.Errorf("Expected result and expected length do not match")
	}

	for i := range result {
		if result[i].Role != expected[i].Role || result[i].Content != expected[i].Content {
			t.Errorf("Error in index %d, got: %v, want: %v.", i, result[i], expected[i])
		}
	}
}

func TestGetOpenAIRole(t *testing.T) {
	if getOpenAIRole(true) != openai.ChatMessageRoleUser {
		t.Error("Expected User role, got", getOpenAIRole(true))
	}
	if getOpenAIRole(false) != openai.ChatMessageRoleAssistant {
		t.Error("Expected Assistant role, got", getOpenAIRole(false))
	}
}
