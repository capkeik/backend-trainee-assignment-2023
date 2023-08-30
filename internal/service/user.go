package service

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/repository/pg"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
)

type User struct {
	User *pg.UserRepo
}

func NewUserService(user *pg.UserRepo) *User {
	return &User{User: user}
}

func (s *User) GetUserSegments(id int32) (*[]string, error) {
	return s.User.GetUserSegments(id)
}

func (s *User) CreateUser(ctx context.Context, id int32) (*model.User, error) {
	return s.User.CreateUser(ctx, id)
}

func (s *User) UpdateUserSegments(
	ctx context.Context,
	slugsToAdd, slugsToRemove *[]string,
	userID int32,
) (*response.UserChanges, error) {
	return s.User.UpdateUserSegments(ctx, slugsToAdd, slugsToRemove, userID)
}
