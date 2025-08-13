package db

import (
	"database/sql"
	"log"
)

func NewMSSQLStorage() (*sql.DB, error) {
	//connString := "sqlserver://apiuser:NewStrongPassword123!@localhost:1433?database=WebProjekatTest"
	connString := "server=localhost,1433;database=webProjekatTest;trusted_connection=yes"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	return db, nil
}
