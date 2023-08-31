package service

import (
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"gorm.io/gorm"
	"time"
)

type RecorderPG struct {
	db *gorm.DB
}

func NewRecordService(db *gorm.DB) RecorderPG {
	return RecorderPG{db: db}
}

func (r RecorderPG) RecordUpdate(added, removed *[]string, id int32) error {
	records := []*model.Record{}

	for _, s := range *removed {
		records = append(
			records,
			&model.Record{
				UserID:          id,
				Slug:            s,
				Action:          model.ActionRemove,
				ActionTimestamp: time.Now(),
			},
		)
	}

	for _, s := range *added {
		records = append(
			records,
			&model.Record{
				UserID:          id,
				Slug:            s,
				Action:          model.ActionAdd,
				ActionTimestamp: time.Now(),
			},
		)
	}
	if err := r.db.Create(records).Error; err != nil {
		return fmt.Errorf("%s: %w", "service.record.RecorderPG.RecordUpdate", err)
	}

	return nil
}
