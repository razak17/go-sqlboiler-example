package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	dbmodels "github.com/razak17/go-sqlboiler-example/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()
	db := connectDB()

	boil.SetDB(db)

	author := createAuthor(ctx)
	fmt.Fprintf(os.Stderr, "DEBUGPRINT[1]: main.go:19: author=%+v\n", author)
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:2345/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}

func createAuthor(ctx context.Context) dbmodels.Author {
	author := dbmodels.Author{
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	err := author.InsertG(ctx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	return author
}
