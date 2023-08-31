package interfaces

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/model"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
)

type UserService interface {
	GetUserSegments(id int32) (*[]string, error)
	CreateUser(ctx context.Context, id int32) (*model.User, error)
	UpdateUserSegments(ctx context.Context, slugsToAdd, slugsToRemove *[]string, userID int32) (*response.UserChanges, error)
}

type SegmentService interface {
	Create(slug string) (*model.Segment, error)
	Delete(slug string) error
}

type UpdateRecorder interface {
	RecordUpdate(added, removed *[]string, id int32) error
}
