package database

import (
	"fmt"
	"log"

	"database/sql"

	configuration "github.com/Satsawat-Rodketkul/domain-user-api/config"
	_ "github.com/lib/pq"
)

func DBconnection() {
	dbHost := configuration.GetConfigValue("DATABASE_HOST")
	dbPort := configuration.GetConfigValue("DATABASE_PORT")
	dbUser := configuration.GetConfigValue("DATABASE_USER")
	dbPassword := configuration.GetConfigValue("DATABASE_PASSWORD")
	dbName := configuration.GetConfigValue("DATABASE_NAME")
	dbMode := configuration.GetConfigValue("DATABASE_SSLMODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbMode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	log.Print("Connect to database success:", db)
}
