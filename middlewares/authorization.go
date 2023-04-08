package middlewares

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := repo.GetDB()

		// mencoba mengambil route parameter berupa productID
		productID, err := strconv.Atoi(c.Param("productID"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}

		var product models.Product

		// mencoba claim token yang telah disimpan dalam autentikasi.
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		// mencoba mendapatkan data berdasarkan id product yang didapatkan dari route productID
		err = db.Select("user_id").First(&product, uint(productID)).Error

		// jika productnya tidak ada maka akan memasuki status not found
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data not found",
				"message": "data doesn't exist",
			})
			return
		}

		// pengecekan jika pada product.userID tidak sama dengan userID yang disimpan dalam jwt.MapClaims maka akan masuk error
		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		// jika sama maka akan masuk ke proses berikutnya, atau endpoint selanjutnya
		c.Next()
	}
}
