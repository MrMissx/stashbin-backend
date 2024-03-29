package models

import (
	"errors"
	"time"

	"github.com/mrmissx/stashbin-backend/utils"
	"gorm.io/gorm"
)

type Document struct {
	Id        uint       `json:"id" gorm:"primary_key;auto_increment"`
	Slug      string     `json:"key" gorm:"type:varchar(10);unique;not null;unqiueIndex"`
	Content   string     `json:"content" gorm:"type:text;not null"`
	CreatedAt *time.Time `json:"date" gorm:"not null"`
	Views     uint       `json:"views" gorm:"type:int;not null;default:0"`
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

func incrementViews(document *Document) {
	document.Views += 1
	DB.Model(&document).Update("views", document.Views)
}

// Get document by slug (key) that are assigned to it
func GetDocumentBySlug(slug string) (*Document, error) {

	var document Document

	if err := DB.Where("slug = ?", slug).First(&document).Error; err != nil {
		return nil, err
	}

	go incrementViews(&document)

	return &document, nil
}
