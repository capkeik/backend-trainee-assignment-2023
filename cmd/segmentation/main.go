package main

import (
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

	log.Println("Opening DB connection")
	db, err := pg.Connect()
	if err != nil {
		return fmt.Errorf("%s, %w", "Error initializing database:", err)
	}
	defer func(db2 *gorm.DB) error {
		sqlDb, err := db.DB()

		log.Println("Closing DB connection")
		err = sqlDb.Close()
		if err != nil {
			return fmt.Errorf("%s, %w", "Error closing DB connection:", err)
		}
		return nil
	}(db)

	return nil
}
