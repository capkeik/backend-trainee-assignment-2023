package pg

import (
	"context"
	"fmt"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return UserRepo{db}
}

func (r *UserRepo) GetUserSegments(id int32) (*[]string, error) {
	fn := "repository.pg.CreateUser"
	var user *model.User
	res := r.db.Preload("Segments").First(&user, id)
	err := res.Error

	if err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	var slugs = []string{}
	for _, s := range user.Segments {
		slugs = append(slugs, s.Slug)
	}
	return &slugs, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, id int32) (*model.User, error) {
	fn := "repository.pg.CreateUser"
	user := model.User{ID: id}
	res := r.db.Create(&user)
	err := res.Error

	if err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	if res.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}

func (r *UserRepo) UpdateUserSegments(
	ctx context.Context,
	slugsToAdd, slugsToRemove *[]string,
	userID int32,
) (*response.UserChanges, error) {
	fn := "repository.pg.UpdateUserSegments"

	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	var segmentsToRemove []model.Segment
	if err := r.db.Where("slug IN (?)", *slugsToRemove).Find(&segmentsToRemove).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	var idsToRemove = []int32{}

	for _, seg := range segmentsToRemove {
		idsToRemove = append(idsToRemove, seg.ID)
	}
	err := r.db.Exec("DELETE FROM user_segments WHERE user_id = ? AND segment_id IN ?", user.ID, idsToRemove).Error
	if err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	var segmentsToAdd []model.Segment
	if err := r.db.Where("slug IN (?)", *slugsToAdd).Find(&segmentsToAdd).Error; err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}

	if err = r.db.Model(&user).Association("Segments").Append(&segmentsToAdd); err != nil {
		return nil, fmt.Errorf("%s, %w", fn, err)
	}
	res := response.UserChanges{
		ID:      user.ID,
		Removed: slugsToRemove,
		Added:   slugsToAdd,
	}
	return &res, nil
}
