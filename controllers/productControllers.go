package controllers

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateProduct(c *gin.Context) {
	db := repo.GetDB()

	// method mustget digunakan untuk mendapatkan claim dari token yang telah disimpan di auth yg dibuat sebelumnya
	userData := c.MustGet("userData").(jwt.MapClaims) // langsung di convert menjadi data sebelumnya

	contengType := helpers.GetContentType(c)

	var product models.Product

	// mengambil userID berdasarkan id dari user yang telah disimpan dalam jwt.MapClaims
	userID := uint(userData["id"].(float64))

	if contengType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	err := db.Debug().Create(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}
