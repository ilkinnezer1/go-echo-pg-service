package models

import (
	"gorm.io/gorm"
	"time"
)

type Projects struct {
	ID         int    `json:"ID" gorm:"primary_key"`
	Name       string `json:"name"`
	Slogan     string `json:"slogan"`
	ShortIntro string `json:"shortIntro"`
	ImagePath  string `json:"imagePath"`
	ImgAltText string `json:"imgAltText"`
	Hyperlink  string `json:"hyperlink"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (p *Projects) ListItemById(db *gorm.DB) (err error) {
	lastSlider := &Partners{}

	if err := db.Order("created_at DESC").First(&lastSlider).Error; err != nil {
		// If there is not any slider, ID will be 1 as a start point
		p.ID = 1
	} else {
		p.ID = lastSlider.ID + 1
	}
	return nil
}
