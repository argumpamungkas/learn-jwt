package repo

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "gume98"
	port     = "5432"
	dbname   = "jwt-api"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	fmt.Println("sukses koneksi ke database")

	db.Debug().AutoMigrate(models.User{}, models.Product{})
	// db.Debug().Migrator().DropTable(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
