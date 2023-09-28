package controllers

import (
	"net/http"

	"github.com/mrmissx/stashbin-backend/models"
	"github.com/mrmissx/stashbin-backend/response"

	"github.com/gin-gonic/gin"
)

func GetDocumentBySlug(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrNoKeyQuery)
		return
	}

	document, err := models.GetDocumentBySlug(key)
	if err != nil {
		c.AbortWithStatusJSON(response.GormErrorToResponse(err))
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewResult(
			"successfully retrieved document",
			gin.H{"content": document.Content},
		),
	)
}

func CreateDocument(c *gin.Context) {
	var document *models.Document

	if err := c.ShouldBindJSON(&document); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.ErrContentEmpty)
		return
	}

	err := models.SaveDocument(document)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrInternalServerError)
		return
	}

	// return redirect header for HTMX request
	if c.GetHeader("HX-Request") == "true" {
		c.Header("HX-Redirect", "/"+document.Slug)
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
