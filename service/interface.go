package service

import (
	"context"

	"github.com/shahzaibaziz/user_operations/db"
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	domain "github.com/shahzaibaziz/user_operations/models"
)

// Manager defines the available functions for the given service implementation.
type Manager interface {
	DeleteUser(ctx context.Context, id string) error
	FindUser(ctx context.Context, id string) (*genModel.UserResponse, error)
	GetUsers(ctx context.Context, filter map[string]interface{}) ([]*genModel.UserResponse, error)
	SaveUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error)
	UpdateUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error)
}

type service struct {
	db db.DataStore
}

// NewService return new service object
func NewService(store *db.DataStore) Manager {
	return &service{db: *store}
}
