package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBUrl dbConnString is the connection string used to connect to the MySQL database.
	DBUrl = "root:xablau@tcp(localhost:3306)/social_network"

	// APIPort is the port on which the API server will listen to.
	APIPort = 5000

	// SecretKey is the secret key used to sign JWT tokens.
	SecretKey []byte
)

func Init() {
	var err error

	// Ignore the error loading .env file, as variables might be provided by the environment (e.g., Docker).
	_ = godotenv.Load()

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 5000
	}

	DBUrl = os.Getenv("DB_CONN_STRING")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
