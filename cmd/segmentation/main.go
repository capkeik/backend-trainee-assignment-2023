package main

import (
	"errors"
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/pg"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting Segmentation Service")
	cfg := config.Get()
	_ = cfg

	pgDB, err := pg.Connect()

	if err != nil {
		log.Fatal(err)
	}

	if pgDB != nil {
		log.Println("Running PostgreSQL migrations")
		if err = runMigrations(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("db connection no")
	}
}

func runMigrations() error {
	fn := "main.runMigrations"
	// can load config because it's implemented to load once
	cfg := config.Get()

	if cfg.PgMigrationsPath == "" {
		return nil
	}

	if cfg.PgURL == "" {
		return errors.New("no PgURL provided")
	}

	m, err := migrate.New(cfg.PgMigrationsPath, cfg.PgURL)
	if err != nil {
		return fmt.Errorf("%s, %w", fn, err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		fmt.Printf("%T: %v", err, err)
		return fmt.Errorf("%s, %w", fn, err)
	}

	return nil
}
