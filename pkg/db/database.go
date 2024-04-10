// db package
package db

import (
	"database/sql"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
)

var db *goqu.Database

func init() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Check if the connection to the database is successful
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	// Create a goqu.Database instance using the SQL database
	db = goqu.New("postgres", sqlDB)
}

func GetDB() *goqu.Database {
	return db
}
