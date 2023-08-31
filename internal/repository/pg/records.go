package pg

import (
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"gorm.io/gorm"
	"time"
)

type RecordsRepo struct {
	db *gorm.DB
}

func NewRecordRepo(db *gorm.DB) RecordsRepo {
	return RecordsRepo{db: db}
}

func (r *RecordsRepo) GetRecords(userID int32, from, to time.Time) (*[]*model.Record, error) {
	if from.IsZero() {
		from = time.Unix(0, 0).UTC()
	}

	if to.IsZero() {
		to = time.Now().UTC()
	}

	records := []*model.Record{}

	if err := r.db.Where("user_id = ? AND action_timestamp BETWEEN ? AND ?", userID, from, to).
		Find(&records).Error; err != nil {
		return nil, err
	}

	return &records, nil
}
