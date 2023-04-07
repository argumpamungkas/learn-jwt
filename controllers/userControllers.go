package controllers

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/helpers"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"log"
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

func UserLogin(c *gin.Context) {
	var user models.User

	db := repo.GetDB()
	contentType := helpers.GetContentType(c)

	// cek headers
	if contentType == appJson {
		c.ShouldBindJSON(&user) // input email, password
	} else {
		c.ShouldBind(&user) // input email, password
	}

	// variable password = user.password
	password := user.Password

	// mengambil data berdasarkan email yang dimasukkan, jika ada maka akan masuk ke step selanjutnya, jika tidak maka akan masuk pada err
	err := db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password, Tidak ditemukan berdasarkan email",
		})
		return
	}

	// compare antara password yang sudah dihashing/diambil dari database, dengan password yang baru diinputkan
	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password, Compare",
		})
		return
	}

	// jika compare berhasil, maka server akan membuat token/generate token dan dikirim pada client
	token, err := helpers.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Println("Invalid generate token")
		return
	}

	// disini
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
