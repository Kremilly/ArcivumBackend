package configs

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadConfig() {
	godotenv.Load()
}

func InitDB() {
	LoadConfig()

	dsn := os.Getenv("DATABASE_PUBLIC_URL")
	if os.Getenv("ENVIRONMENT") == "development" {
		dsn = os.Getenv("LOCAL_DATABASE_URL")
	}

	if dsn == "" {
		log.Fatalf("DATABASE_URL or LOCAL_DATABASE_URL environment variable not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}

func GetRootURL(c fiber.Ctx) string {
	scheme := "http"
	if c.Secure() {
		scheme = "https"
	}

	host := c.Hostname()
	return scheme + "://" + host
}

func DetectEnv(c fiber.Ctx) string {
	host := c.Hostname()

	if host == "localhost" || host == "127.0.0.1" {
		return "development"
	}
	return "production"
}
