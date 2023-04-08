package main

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/router"
	"os"
)

func main() {
	repo.StartDB()

	var PORT = os.Getenv("PORT")
	router.StartApp().Run(":" + PORT)
}
