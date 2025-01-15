package tests

import (
	"testing"

	"acheisuacara.com.br/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockDB is a mock implementation of the database
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *MockDB {
	m.Called(query, args)
	return m
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds)
	return args.Error(0)
}

func (m *MockDB) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

func (m *MockDB) Model(value interface{}) *MockDB {
	m.Called(value)
	return m
}

func (m *MockDB) Update(column string, value interface{}) error {
	args := m.Called(column, value)
	return args.Error(0)
}

func TestCreateShortURL(t *testing.T) {
	mockDB := new(MockDB)
	service := services.NewURLService(mockDB)

	tests := []struct {
		name        string
		url         string
		setupMock   func()
		expectError bool
	}{
		{
			name: "Valid URL - New Entry",
			url:  "https://www.amazon.com/product/123",
			setupMock: func() {
				mockDB.On("Where", "url = ?", mock.Anything).Return(mockDB)
				mockDB.On("First", mock.Anything, mock.Anything).Return(gorm.ErrRecordNotFound)
				mockDB.On("Create", mock.Anything).Return(nil)
			},
			expectError: false,
		},
		{
			name: "Valid URL - Existing Entry",
			url:  "https://www.amazon.com/product/123",
			setupMock: func() {
				mockDB.On("Where", "url = ?", mock.Anything).Return(mockDB)
				mockDB.On("First", mock.Anything, mock.Anything).Return(nil)
			},
			expectError: false,
		},
		{
			name:        "Invalid URL",
			url:         "not-a-url",
			setupMock:   func() {},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB = new(MockDB)
			service = services.NewURLService(mockDB)
			tt.setupMock()

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
	mockDB := new(MockDB)
	service := services.NewURLService(mockDB)

	tests := []struct {
		name        string
		shortCode   string
		setupMock   func()
		expectError bool
	}{
		{
			name:      "Existing Short Code",
			shortCode: "abc123",
			setupMock: func() {
				mockDB.On("Where", "short_code = ?", mock.Anything).Return(mockDB)
				mockDB.On("First", mock.Anything, mock.Anything).Return(nil)
				mockDB.On("Model", mock.Anything).Return(mockDB)
				mockDB.On("Update", "clicks", mock.Anything).Return(nil)
			},
			expectError: false,
		},
		{
			name:      "Non-existing Short Code",
			shortCode: "nonexist",
			setupMock: func() {
				mockDB.On("Where", "short_code = ?", mock.Anything).Return(mockDB)
				mockDB.On("First", mock.Anything, mock.Anything).Return(gorm.ErrRecordNotFound)
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB = new(MockDB)
			service = services.NewURLService(mockDB)
			tt.setupMock()

			result, err := service.GetLongURL(tt.shortCode)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}
		})
	}
}
