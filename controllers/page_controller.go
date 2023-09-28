package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrmissx/stashbin-backend/models"
)

func Index(c *gin.Context) {
	fmt.Println("Index")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func ContentPage(c *gin.Context) {
	fmt.Println("/slug")
	key := c.Param("slug")
	if key == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	document, err := models.GetDocumentBySlug(key)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	fmt.Println(document.Content)

	c.HTML(http.StatusOK, "content.html", gin.H{})
}

func RawPage(c *gin.Context) {
	key := c.Param("slug")
	if key == "" {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	document, err := models.GetDocumentBySlug(key)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.String(http.StatusOK, document.Content)
}

func AboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}
