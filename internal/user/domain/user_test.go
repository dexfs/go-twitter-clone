package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Should create a new User instance correct",
			input:    "User 1",
			expected: "User 1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sut := NewUser(test.input)
			if sut.Username != test.expected {
				t.Errorf("got %q want %q", sut.Username, test.expected)
			}

			if sut.ID == "" {
				t.Errorf("got %q want %q", sut.ID, "")
			}

			if sut.CreatedAt == (time.Time{}) {
				t.Errorf("got %q want %q", sut.CreatedAt, time.Time{})
			}
			if sut.UpdatedAt == (time.Time{}) {
				t.Errorf("got %q want %q", sut.UpdatedAt, time.Time{})
			}
		})
	}
}
