package main

import (
	"github.com/Satsawat-Rodketkul/domain-user-api/config"
	"github.com/Satsawat-Rodketkul/domain-user-api/internal/database"
)

func main() {
	config.LoadEnv()
	database.DBconnection()
}
