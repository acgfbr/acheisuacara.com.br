package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type URL struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	URL       string    `json:"url" gorm:"type:varchar(2048);not null"`
	ShortCode string    `json:"short_code" gorm:"type:varchar(10);unique;not null"`
	Clicks    uint      `json:"clicks" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *URL) Validate() bool {
	// Check if it's a valid URL
	if !govalidator.IsURL(u.URL) {
		return false
	}

	// Check if it's from a marketplace (you can add more marketplaces as needed)
	marketplaces := []string{
		"amazon.com",
		"mercadolivre.com",
		"americanas.com",
		"magazineluiza.com",
		"shopee.com",
		"aliexpress.com",
	}

	for _, marketplace := range marketplaces {
		if govalidator.Contains(u.URL, marketplace) {
			return true
		}
	}

	return false
}
