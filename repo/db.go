package repo

import (
	"DTS/Chapter-3/sesi/sesi2-go-jwt/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	port     = os.Getenv("PGPORT")
	dbname   = os.Getenv("PGDATABASE")
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
	// db.Debug().Migrator().DropTable(models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
