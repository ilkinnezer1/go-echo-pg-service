package models

import (
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	Description string `json:"description"`
	Hyperlink   string `json:"hyperlink"`
	Pinned      bool   `json:"pinned"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GenerateNewBlogId generates a new ID for the blog before inserting it into the database.
func (b *Blog) GenerateNewBlogId(db *gorm.DB) (err error) {
	lastBlog := &Blog{}

	if err := db.Order("created_at DESC").First(&lastBlog).Error; err != nil {
		// Starting point at 1 if there is not any error
		b.ID = 1
	} else {
		// Update new blog id based on the prev one
		b.ID = lastBlog.ID + 1
	}

	return nil
}
