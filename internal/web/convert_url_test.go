package web

import "testing"

func TestReplaceWhitespacesWithUnderscores(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "string with single space",
			input:    "hello world",
			expected: "hello_world",
		},
		{
			name:     "string with multiple spaces",
			input:    "hello  world  test",
			expected: "hello__world__test",
		},
		{
			name:     "string without spaces",
			input:    "helloworld",
			expected: "helloworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceWhitespacesWithUnderscores(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceWhitespacesWithUnderscores() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReplaceUnderscoresWithWhitespaces(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "string with single underscore",
			input:    "hello_world",
			expected: "hello world",
		},
		{
			name:     "string with multiple underscores",
			input:    "hello__world__test",
			expected: "hello  world  test",
		},
		{
			name:     "string without underscores",
			input:    "helloworld",
			expected: "helloworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceUnderscoresWithWhitespaces(tt.input)
			if result != tt.expected {
				t.Errorf("ReplaceUnderscoresWithWhitespaces() = %v, want %v", result, tt.expected)
			}
		})
	}
}
