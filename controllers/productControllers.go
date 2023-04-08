package controllers

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"fmt"
	"net/http"
	"strconv"

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

func UpdateProduct(c *gin.Context) {
	db := repo.GetDB()
	var product models.Product

	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	productID, _ := strconv.Atoi(c.Param("productID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	product.UserID = userID

	product.ID = uint(productID)

	err := db.Model(&product).Where("id = ? ", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := repo.GetDB()
	var product models.Product

	productID, _ := strconv.Atoi(c.Param("productID"))

	product.ID = uint(productID)

	err := db.Where("id = ?", productID).Delete(&product).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("product with id %v success deleted", productID),
	})
}
