package helpers

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, email string) (res string, err error) {

	// jwt.MapClaims sebuah tipe data map[string]interface{} yang berasal dari package jwt-go
	// data yg tidak boleh disimpan dalam jwt ini adalah data yg sifatnya kredensial seperti password, pin atm
	// maka dari itu sangat disarankan hanya menyimpan id, email, username, waktu login dari user
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	// metode enkripsi yang digunakan adalah HS256
	// method jwt.NewWithClaims digunakan untuk memasukkan data-data user yang kita simpan pada jwt.MapClaims dan sekaligus
	// menentukan metode enkripsinya, dan akan mengembalikan sebuah nilai dengan struct *jwt.token
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// jwt/parse.SignedString adalah method dari struct *jwt.token. method ini digunakan untuk parsing token menjadi sebuah string panjang
	// yang nantinya akan dikirimkan oleh server kepada client. method ini menerima satu param yaitu secret key. Secret key data yg sangat
	// kredensial karena akan digunakan untuk autentikasi token.
	res, err = parseToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		log.Println("error signed token")
		return
	}

	return
}

func VerifyToken(c *gin.Context) (res interface{}, err error) {

	// Custom error
	errResponse := errors.New("Sign in to proceed")

	// mendapatkan nilai dari variable Authorization yang terletak pada headers yang dikirimkan oleh client. dimana setiap client akan
	// mengirimkan request, yang dimana setiap request endpoint memerlukan autentikasi token, maka client harus mengirimkan dan menyimpannya
	// dalam headers variable Authorization
	headerToken := c.Request.Header.Get("Authorization")

	// memeriksa jika token yang dikirimkan memiliki prefix Bearer, token yang dikirim harus berupa Beare Token
	if !strings.HasPrefix(headerToken, "Bearer") || headerToken == "" {
		log.Println("Invalid token")
		err = errors.New(errResponse.Error())
		return
	}

	// pengambilan token tanpa prefix Bearer, jadi Beare Token(kode tokennya)
	stringToken := strings.Split(headerToken, " ")[1]

	// mencoba untuk memparsing tokennya menjadi sebuah struct dari *jwt.Token. Kemudian kita memeriksa apakah metode enkripsi dari
	// tokennya adalah metode HS256 dengan cara mengcasting metodenya menjadi tipe data pointer dari struct jwt.SigningMethodHMAC
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	// memeriksa apakah saat mengcasting claim tokennya menjadi tipe data jwt.MapClaims menghasilkan error atau tidak
	// sekaligus memeriksa apakah tokennya merupakan token yang valid atau tidak
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	// mengembalikan nilai berupa claim dari tokennya yang dimana claim tersebut berisikan data yang kita simpan pada tokennya
	// ketika pertama kali dibuat. isinya email dan id user yang sudah berhasil melakukan login
	res = token.Claims.(jwt.MapClaims)

	return
}
