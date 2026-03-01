package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBUrl dbConnString is the connection string used to connect to the MySQL database.
	DBUrl = "root:xablau@tcp(localhost:3306)/social_network"

	// APIPort is the port on which the API server will listen to.
	APIPort = 5000
)

func Init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 5000
	}

	DBUrl = os.Getenv("DB_CONN_STRING")
}
