package main

import (
	"database/sql"
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/pg"
	"gorm.io/gorm"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting Segmentation Service")
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Println("Reading config")
	cfg := config.Get()
	_ = cfg

	db, err := InitDB()
	if err != nil {
		return err
	}

	_ = db
	return nil
}

func InitDB() (*gorm.DB, error) {
	log.Println("Opening DB connection")
	db, err := pg.Connect()
	if err != nil {
		return nil, fmt.Errorf("%s, %w", "Error initializing database:", err)
	}
	sqlDb, err := db.DB()
	defer func(sqlDb *sql.DB) {
		log.Println("Closing DB connection")
		err = sqlDb.Close()
		if err != nil {
			err = fmt.Errorf("%s, %w", "Error closing database:", err)
		}
	}(sqlDb)

	return db, err
}
