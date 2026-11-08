package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBUrl dbConnString is the connection string used to connect to the MySQL database.
	DBUrl = "root:root@/socialnetwork"

	// APIPort is the port on which the API server will listen to.
	APIPort = 8080

	// SecretKey is the secret key used to sign JWT tokens.
	SecretKey = []byte("ORBahUm1ITQAJWrLPEGHzIN8W2kpi+DZbSKdh+Wzelyq88NrZMOi0TumNiK/qsqMo4o7iiMN2eELQSwlJHLqfw==")
)

func Init() {
	var err error

	// Ignore the error loading .env file, as variables might be provided by the environment (e.g., Docker).
	_ = godotenv.Load()

	APIPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		APIPort = 8080
	}

	dbConnString := os.Getenv("DB_CONN_STRING")
	if dbConnString != "" {
		DBUrl = dbConnString
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
