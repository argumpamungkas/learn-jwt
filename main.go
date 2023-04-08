package main

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/repo"
	"DTS/Chapter-3/sesi/sesi2-go-jwt/router"
)

func main() {
	repo.StartDB()
	
	router.StartApp().Run(":8080")
}
