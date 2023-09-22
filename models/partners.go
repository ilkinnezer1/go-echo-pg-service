package models

import (
	"gorm.io/gorm"
	"time"
)

type Partners struct {
	ID        int    `json:"id" gorm:"primary_key"`
	ImagePath string `json:"imagePath"`
	Title     string `json:"title"`
	AltText   string `json:"altText"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GenerateNewId update new partner id while being created
func (p *Partners) GenerateNewId(db *gorm.DB) (err error) {
	lastPartner := &Partners{}

	if err := db.Order("created_at DESC").First(&lastPartner).Error; err != nil {
		// If there is not any last partner, ID will be 1 as a start point
		p.ID = 1
	} else {
		// Update new id based on the last one
		p.ID = lastPartner.ID + 1
	}

	return nil
}
