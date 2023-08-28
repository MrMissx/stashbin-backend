package models

import "time"

type Document struct {
	Id        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Slug      string     `gorm:"type:varchar(10);unique;not null" json:"key"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt *time.Time `gorm:"not null" json:"date"`
}

func (Document) TableName() string {
	return "documents"
}
