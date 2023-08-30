package pg

import (
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"gorm.io/gorm"
	"log"
)

type SegmentRepo struct {
	db *gorm.DB
}

func NewSegmentRepo(db *gorm.DB) SegmentRepo {
	return SegmentRepo{db: db}
}

func (r SegmentRepo) Create(slug string) (*model.Segment, error) {
	fn := "repository.pg.SegmentRepo.Create"
	segment := &model.Segment{Slug: slug}
	log.Println("slug:", slug)
	res := r.db.Create(segment)
	err := res.Error

	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return segment, nil
}

func (r SegmentRepo) Delete(slug string) error {
	fn := "repository.pg.SegmentRepo.Delete"
	var seg *model.Segment
	res := r.db.
		Where("slug = ?", slug).Delete(&seg)
	err := res.Error
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

// All TODO Remove before pushing with all usages
func (r SegmentRepo) All() ([]*model.Segment, error) {
	fn := "repository.pg.SegmentRepo.All"
	var segments []*model.Segment

	res := r.db.Find(&segments)
	err := res.Error
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return segments, nil
}
