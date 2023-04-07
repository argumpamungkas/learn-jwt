package helpers

import (
	"log"
	"os"

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
