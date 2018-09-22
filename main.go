package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"

	"github.com/avinoth/go-bootstrap/src/api"
	"github.com/avinoth/go-bootstrap/src/config/db"
	"github.com/avinoth/go-bootstrap/src/model"
)

var appEnvironment string

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func initDB() {
	dbName := os.Getenv("DATABASE_NAME")

	if len(dbName) == 0 {
		log.Fatal("Database details not specified. Set it using DATABASE_NAME")
	}

	db.InitializeDB(appEnvironment, dbName)
}

func initServer() {
	port := getenv("PORT", "12345")

	api.RunServer(port)
}

func main() {
	appEnvironment = getenv("APP_ENV", "development")

	log.Info("Environment: " + appEnvironment)

	initDB()
	model.AutoMigrate()
	initServer()

	defer db.Instance().Close()
}
