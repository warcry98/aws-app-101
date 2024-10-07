package initializer

import (
	"backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataDB struct {
	Data map[string]string
}

func LoadEnv() DataDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dataEnv := &DataDB{}

	dataEnv.Data = make(map[string]string)

	dataEnv.Data["PG_URL"] = os.Getenv("PG_URL")
	dataEnv.Data["PG_PORT"] = os.Getenv("PG_PORT")
	dataEnv.Data["PG_USERNAME"] = os.Getenv("PG_USERNAME")
	dataEnv.Data["PG_PASSWORD"] = os.Getenv("PG_PASSWORD")
	dataEnv.Data["PG_DATABASE"] = os.Getenv("PG_DATABASE")

	return *dataEnv
}

func ConnectToDB(dataEnv DataDB) *gorm.DB {
	DSN := "host=" + dataEnv.Data["PG_URL"] + " user=" + dataEnv.Data["PG_USERNAME"] + " password=" + dataEnv.Data["PG_PASSWORD"] + " dbname=" + dataEnv.Data["PG_DATABASE"] + " port=" + dataEnv.Data["PG_PORT"] + " sslmode=disable TimeZone=Asia/Jakarta"
	DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB")
	}

	return DB
}

func SyncDatabase(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
