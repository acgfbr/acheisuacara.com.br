package services

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"

	"acheisuacara.com.br/pkg/models"
	"gorm.io/gorm"
)

type URLService struct {
	db *gorm.DB
}

func NewURLService(db *gorm.DB) *URLService {
	return &URLService{db: db}
}

func (s *URLService) CreateShortURL(longURL string) (*models.URL, error) {
	url := &models.URL{
		URL: longURL,
	}

	if !url.Validate() {
		return nil, errors.New("invalid URL or not from a supported marketplace")
	}

	// Generate short code
	hash := sha256.Sum256([]byte(longURL))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	shortCode := strings.ToLower(encoded[:8])

	url.ShortCode = shortCode

	// Check if URL already exists
	var existingURL models.URL
	if err := s.db.Where("url = ?", longURL).First(&existingURL).Error; err == nil {
		return &existingURL, nil
	}

	if err := s.db.Create(url).Error; err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetLongURL(shortCode string) (*models.URL, error) {
	var url models.URL
	if err := s.db.Where("short_code = ?", shortCode).First(&url).Error; err != nil {
		return nil, err
	}

	// Increment click count
	s.db.Model(&url).Update("clicks", url.Clicks+1)

	return &url, nil
}
