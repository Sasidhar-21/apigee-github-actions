package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDatabase(dbName string, db *sql.DB) error {
	var err error
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///migrations"),
		dbName,
		driver)
	if err != nil {
		log.Printf("Failed creating DB migration %v", err)
		return err
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No DB migration was necessary")
			return nil
		} else {
			log.Printf("Failed migrated DB %v", err)
			return err

		}
	}
	return nil
}
