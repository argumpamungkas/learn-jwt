package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Func ini digunakan untuk meng-hashing password user sebelum disimpan ke dalam database
func HashPass(p string) (result string, err error) {
	salt := 8
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Println("error generate password")
	}

	result = string(hash)
	return
}

// Func ini digunakan untuk mengkomparasi password user yang sudah di hash dengan password user yang diinput ketika login
func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	if err != nil {
		log.Println("error compare hash and password")
		return false
	}

	return true
}
