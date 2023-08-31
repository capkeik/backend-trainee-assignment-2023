package service

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service/interfaces"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
	"log"
)

type User struct {
	User     interfaces.UserService
	Recorder interfaces.UpdateRecorder
}

func NewUserService(user interfaces.UserService, record interfaces.UpdateRecorder) *User {
	return &User{User: user, Recorder: record}
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
	log.Println("UserService: ", "Updating user ", userID, " slugs")
	updates, err := s.User.UpdateUserSegments(ctx, slugsToAdd, slugsToRemove, userID)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}
	log.Println("UserService: ", "Recording changes")
	err = s.Recorder.RecordUpdate(updates.Added, updates.Removed, updates.ID)
	return updates, nil
}
