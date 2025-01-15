package services

import (
	"strings"
	"testing"

	"acheisuacara.com.br/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.URL{})
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func TestCreateShortURL(t *testing.T) {
	db := setupTestDB(t)
	service := NewURLService(db)

	tests := []struct {
		name        string
		url         string
		expectError bool
	}{
		{
			name:        "Valid URL - New Entry",
			url:         "https://www.amazon.com/product/123",
			expectError: false,
		},
		{
			name:        "Valid URL - Existing Entry",
			url:         "https://www.amazon.com/product/123",
			expectError: false,
		},
		{
			name:        "Invalid URL",
			url:         "not-a-url",
			expectError: true,
		},
		{
			name:        "Empty URL",
			url:         "",
			expectError: true,
		},
		{
			name:        "Very Long URL",
			url:         "https://www.amazon.com/product/" + strings.Repeat("a", 2048),
			expectError: false,
		},
		{
			name:        "URL with Special Characters",
			url:         "https://www.amazon.com/product/123?q=test&category=electronics#section",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.CreateShortURL(tt.url)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.NotEmpty(t, result.ShortCode)
			}
		})
	}
}

func TestGetLongURL(t *testing.T) {
	db := setupTestDB(t)
	service := NewURLService(db)

	// Create a test URL
	testURL := &models.URL{
		URL:       "https://www.amazon.com/product/123",
		ShortCode: "abc123",
		Clicks:    0,
	}
	db.Create(testURL)

	tests := []struct {
		name        string
		shortCode   string
		expectError bool
		expectURL   string
	}{
		{
			name:        "Existing Short Code",
			shortCode:   "abc123",
			expectError: false,
			expectURL:   "https://www.amazon.com/product/123",
		},
		{
			name:        "Non-existing Short Code",
			shortCode:   "nonexist",
			expectError: true,
			expectURL:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GetLongURL(tt.shortCode)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectURL, result.URL)
				assert.Greater(t, result.Clicks, uint(0))
			}
		})
	}
}
