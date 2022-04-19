package service

import (
	"time"

	"github.com/google/uuid"

	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	domain "github.com/shahzaibaziz/user_operations/models"
)

func addAudit(user *domain.User) {
	user.CreatedAt = time.Now()
	user.ID = uuid.New().String()
}

func asUserResponse(user *domain.User) *genModel.UserResponse {
	return &genModel.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Address:   user.Address,
		CreatedAt: user.CreatedAt.String(),
	}
}
