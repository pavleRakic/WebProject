package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	datab "github.com/pavleRakic/testGoApi/db"
)

func main() {
	db, err := datab.NewMSSQLStorage()

	if err != nil {
		log.Fatal(err)
	}

	driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mssqldb",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	/*if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migration complete")*/

}
