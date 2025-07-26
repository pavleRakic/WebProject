package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pavleRakic/testGoApi/cmd/api"
	"github.com/pavleRakic/testGoApi/db"
)

func main() {
	db, err := db.NewMSSQLStorage()

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer("0.0.0.0:8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}
