package models

import (
	"errors"
	"time"

	"github.com/mrmissx/stashbin-backend/utils"
	"gorm.io/gorm"
)

type Document struct {
	Id        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Slug      string     `gorm:"type:varchar(10);unique;not null" json:"key"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt *time.Time `gorm:"not null" json:"date"`
}

func (Document) TableName() string {
	return "documents"
}

func insertDocument(document *Document, retry int) error {
	document.Slug = utils.CreateSlug()

	if err := DB.Create(&document).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) && retry <= 3 {
			return insertDocument(document, retry+1)
		}
		return err
	}
	return nil
}

// Save document to database with unique slug (key)
// Retry 3 times if slug is duplicated
func SaveDocument(document *Document) error {
	return insertDocument(document, 0)
}

// Get document by slug (key) that are assigned to it
func GetDocumentBySlug(slug string) (*Document, error) {

	var document Document

	if err := DB.Where("slug = ?", slug).First(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}
