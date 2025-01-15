package models

import (
	"testing"
	"time"
)

func TestURLValidation(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected bool
	}{
		{
			name:     "Valid Amazon URL",
			url:      "https://www.amazon.com/product/123",
			expected: true,
		},
		{
			name:     "Valid MercadoLivre URL",
			url:      "https://www.mercadolivre.com/product/123",
			expected: true,
		},
		{
			name:     "Invalid URL format",
			url:      "not-a-url",
			expected: false,
		},
		{
			name:     "Valid URL but not marketplace",
			url:      "https://www.google.com",
			expected: false,
		},
		{
			name:     "Empty URL",
			url:      "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := &URL{
				URL:       tt.url,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			result := url.Validate()
			if result != tt.expected {
				t.Errorf("URL.Validate() = %v, want %v", result, tt.expected)
			}
		})
	}
}
