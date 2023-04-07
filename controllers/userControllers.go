package controllers

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJson = "application/json"
)

func UserRegister(c *gin.Context) {
	var user models.User

	db := repo.GetDB()

	contentType := helpers.GetContentType(c)

	// _, _ := db, contentType

	if contentType == appJson {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
}
