package main

import (
	configuration "github.com/Satsawat-Rodketkul/domain-user-api/config"
	"github.com/Satsawat-Rodketkul/domain-user-api/internal/database"
)

func main() {
	configuration.LoadEnv()
	database.DBconnection()
}
