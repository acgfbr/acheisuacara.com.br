package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"acheisuacara.com.br/pkg/handlers"
	"acheisuacara.com.br/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockURLService struct {
	mock.Mock
}

func (m *MockURLService) CreateShortURL(url string) (*models.URL, error) {
	args := m.Called(url)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

func (m *MockURLService) GetLongURL(shortCode string) (*models.URL, error) {
	args := m.Called(shortCode)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.URL), args.Error(1)
}

func TestCreateShortURL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    map[string]interface{}
		setupMock      func(*MockURLService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid URL",
			requestBody: map[string]interface{}{
				"url": "https://www.amazon.com/product/123",
			},
			setupMock: func(m *MockURLService) {
				m.On("CreateShortURL", "https://www.amazon.com/product/123").Return(
					&models.URL{
						URL:       "https://www.amazon.com/product/123",
						ShortCode: "abc123",
					}, nil)
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
			setupMock: func(m *MockURLService) {
				m.On("CreateShortURL", "not-a-url").Return(nil,
					errors.New("invalid URL or not from a supported marketplace"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "invalid URL or not from a supported marketplace",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockURLService)
			tt.setupMock(mockService)
			handler := handlers.NewURLHandler(mockService)

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
			assert.Equal(t, tt.expectedBody, responseBody)
		})
	}
}

func TestRedirectToLongURL(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		shortCode      string
		setupMock      func(*MockURLService)
		expectedStatus int
		expectedURL    string
	}{
		{
			name:      "Valid Short Code",
			shortCode: "abc123",
			setupMock: func(m *MockURLService) {
				m.On("GetLongURL", "abc123").Return(
					&models.URL{
						URL:       "https://www.amazon.com/product/123",
						ShortCode: "abc123",
					}, nil)
			},
			expectedStatus: http.StatusMovedPermanently,
			expectedURL:    "https://www.amazon.com/product/123",
		},
		{
			name:      "Invalid Short Code",
			shortCode: "nonexist",
			setupMock: func(m *MockURLService) {
				m.On("GetLongURL", "nonexist").Return(nil,
					errors.New("URL not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedURL:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockURLService)
			tt.setupMock(mockService)
			handler := handlers.NewURLHandler(mockService)

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
