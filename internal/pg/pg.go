package pg

import (
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	fn := "pg.Connect"

	cnf := config.Get()

	db, err := gorm.Open(postgres.Open(cnf.PgURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	return db, nil
}
