package shared

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type PostgresManager struct {
	Db *gorm.DB
}

var Manager *PostgresManager

func init() {
	log.Println("PostgresManager init")

	if os.Getenv("TESTING") == "true" {
		return
	}

	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connect Error !")
	}

	Manager = &PostgresManager{Db: db}
}

// GetManager : Get PostgresManager
func GetManager() *PostgresManager {
	return Manager
}
