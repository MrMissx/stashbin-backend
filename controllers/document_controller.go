package controllers

import (
	"errors"
	"net/http"

	"github.com/mrmissx/stashbin-backend/models"
	"github.com/mrmissx/stashbin-backend/response"
	"github.com/mrmissx/stashbin-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDocumentBySlug(c *gin.Context) {
	var document models.Document
	key := c.Query("key")
	if key == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrNoKeyQuery)
		return
	}

	if err := models.DB.Where("slug = ?", key).First(&document).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, response.ErrDocumentNotFound)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrInternalServerError)
			return
		}
	}

	c.JSON(
		http.StatusCreated,
		response.NewResult(
			"successfully retrieved document",
			gin.H{"content": document.Content},
		),
	)
}

func insertDocument(document *models.Document, retry int) (*models.Document, error) {
	document.Slug = utils.CreateSlug()

	if err := models.DB.Create(&document).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) && retry <= 3 {
			return insertDocument(document, retry+1)
		}
		return nil, err
	}
	return document, nil
}

func CreateDocument(c *gin.Context) {
	var document *models.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrContentEmpty)
		return
	}

	document, err := insertDocument(document, 0)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrInternalServerError)
		return
	}

	c.JSON(
		http.StatusCreated,
		response.NewResult(
			"successfully created document",
			gin.H{
				"key":    document.Slug,
				"length": len(document.Content),
				"date":   document.CreatedAt,
			},
		),
	)
}
