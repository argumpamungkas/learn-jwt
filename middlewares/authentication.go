package middlewares

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		// method set untuk menyimpan claim tokennya kedalam data request agar dapat diambil pada endpoint berikutnya
		c.Set("userData", verifyToken)

		// next digunakan agar dapat melajutkan proses kepada endpoint berikutnya
		c.Next()
	}
}
