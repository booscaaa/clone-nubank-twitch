package factory

import (
	"database/sql"
	"fmt"
	"os"
)

func GetConnection() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	return db
}
