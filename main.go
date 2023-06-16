package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	db := connectDB()

	boil.SetDB(db)
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:pass@localhost:2345/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
