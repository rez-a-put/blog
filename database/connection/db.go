package connection

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() *sql.DB {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbConnection := os.Getenv("DB_CONN")

	DB, err = sql.Open(dbDriver, dbConnection)
	if err != nil {
		panic(err.Error())
	}

	return DB
}
