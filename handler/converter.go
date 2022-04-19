package handler

import (
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	domain "github.com/shahzaibaziz/user_operations/models"
)

func toUserModel(user *genModel.User) *domain.User {
	return &domain.User{
		Name:    *user.Name,
		Address: *user.Address,
		Email:   user.Email.String(),
	}
}
