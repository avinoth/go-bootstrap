package db

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbInstance *gorm.DB

// InitializeDB fetches DB NAME from ENV and connects to the DB
func InitializeDB(goAppEnvironment string, dbName string) {
	dbConnectionString := "dbname=" + dbName + " sslmode=disable"

	log.Info("DB connection string: " + dbConnectionString)

	dbConnect, err := gorm.Open("postgres", dbConnectionString)

	if err != nil {
		log.Fatal("Error Connecting to Database: ", dbConnectionString, ". ", err)
	}

	dbConnect.LogMode(goAppEnvironment == "development")

	if goAppEnvironment == "production" {
		dbConnect.DB().SetMaxIdleConns(4)
		dbConnect.DB().SetMaxOpenConns(20)
	}

	dbInstance = dbConnect
}

// Instance returns the current database Instance
func Instance() *gorm.DB {
	return dbInstance
}
