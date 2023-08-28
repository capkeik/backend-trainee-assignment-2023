package pg

import (
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"time"
)

const Timeout = 5

type DB struct {
	*pg.DB
}

func Connect() (*DB, error) {
	cfg := config.Get()

	fn := "pg.Connect"

	pgOpts, err := pg.ParseURL(cfg.PgURL)
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	pgDB := pg.Connect(pgOpts)

	pgDB.WithTimeout(time.Second * Timeout)

	return &DB{pgDB}, nil
}
