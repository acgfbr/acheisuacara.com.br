package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"acheisuacara.com.br/pkg/models"
	"acheisuacara.com.br/pkg/services"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.TestMode)
	db := setupTestDB(t)
	service := services.NewURLService(db)
	handler := NewURLHandler(service)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid URL",
			requestBody: map[string]interface{}{
				"url": "https://www.amazon.com/product/123",
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"url":        "https://www.amazon.com/product/123",
				"short_code": "abc123",
			},
		},
		{
			name: "Invalid URL",
			requestBody: map[string]interface{}{
				"url": "not-a-url",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "URL inválida ou não suportada",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.POST("/api/shorten", handler.CreateShortURL)

			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/shorten", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)

			var responseBody map[string]interface{}
			json.Unmarshal(resp.Body.Bytes(), &responseBody)

			if tt.expectedStatus == http.StatusOK {
				assert.NotEmpty(t, responseBody["short_code"])
				assert.Equal(t, tt.requestBody["url"], responseBody["url"])
			} else {
				assert.Equal(t, tt.expectedBody, responseBody)
			}
		})
	}
}

func TestRedirectToLongURL(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB(t)
	service := services.NewURLService(db)
	handler := NewURLHandler(service)

	// Create a test URL
	testURL := &models.URL{
		URL:       "https://www.amazon.com/product/123",
		ShortCode: "abc123",
		Clicks:    0,
	}
	db.Create(testURL)

	tests := []struct {
		name           string
		shortCode      string
		expectedStatus int
		expectedURL    string
	}{
		{
			name:           "Existing Short Code",
			shortCode:      "abc123",
			expectedStatus: http.StatusMovedPermanently,
			expectedURL:    "https://www.amazon.com/product/123",
		},
		{
			name:           "Non-existing Short Code",
			shortCode:      "nonexist",
			expectedStatus: http.StatusNotFound,
			expectedURL:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/:shortCode", handler.RedirectToLongURL)

			req := httptest.NewRequest("GET", "/"+tt.shortCode, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			if tt.expectedURL != "" {
				assert.Equal(t, tt.expectedURL, resp.Header().Get("Location"))
			}
		})
	}
}
