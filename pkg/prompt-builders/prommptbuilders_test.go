package promptbuilders

import (
	"strings"
	"testing"
)

func TestBuild(t *testing.T) {
	builder := NewPromptBuilder()

	testCases := []struct {
		name        string
		input       string
		replyOption string
		expected    string
	}{
		{
			name:        "Test JokeReplyOption",
			input:       "Dogs",
			replyOption: JokeReplyOption,
			expected:    "Generate funny joke about topic: Dogs",
		},
		{
			name:        "Test SimpleReplyOption",
			input:       "Hello",
			replyOption: SimpleReplyOption,
			expected:    "Hello",
		},
		{
			name:        "Test TwoSentenceHorrorRU",
			input:       "Noche",
			replyOption: TwoSentenceHorrorRU,
			expected:    "Noche",
		},
		{
			name:        "Test Default",
			input:       "Hi",
			replyOption: "non_existent_option",
			expected:    "Hi",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := builder.Build(tc.input, tc.replyOption)
			if tc.replyOption == TwoSentenceHorrorRU {
				// Check if the theme input is present in the result
				if !strings.Contains(result, tc.input) {
					t.Errorf("Expected %v to be present in the result", tc.input)
				}
			} else if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
