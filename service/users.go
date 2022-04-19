package service

import (
	"context"

	wraperrors "github.com/pkg/errors"

	domainErr "github.com/shahzaibaziz/user_operations/errors"
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	domain "github.com/shahzaibaziz/user_operations/models"
)

// SaveUser save user object
func (s *service) SaveUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error) {
	filter := map[string]interface{}{"email": user.Email}

	usr, err := s.db.FindUserByFilter(ctx, filter)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user having given email")
	}

	addAudit(user)
	if usr == nil {
		if _, err = s.db.SaveUser(ctx, user); err != nil {
			log(ctx).Errorf("failed to save user in database ", err)

			return nil, wraperrors.Wrap(err, "failed to save user")
		}
	} else {
		log(ctx).Error("user already exist")

		return nil, wraperrors.Wrap(domainErr.ErrConflict, "failed to save user")
	}

	return asUserResponse(user), nil
}

// FindUser find user object using id
func (s *service) FindUser(ctx context.Context, id string) (*genModel.UserResponse, error) {
	usr, err := s.db.FindUser(ctx, id)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user having given email")
	}
	if usr == nil {
		log(ctx).Error("no user found")

		return nil, wraperrors.Wrap(domainErr.ErrNoContent, "failed to find user")
	} else if err != nil {
		return nil, wraperrors.Wrap(err, "failed to make query to database")
	}

	return asUserResponse(usr), nil
}

// GetUsers fetch all users using filter
func (s *service) GetUsers(ctx context.Context, filter map[string]interface{}) ([]*genModel.UserResponse, error) {
	users, err := s.db.GetUsers(ctx, filter)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user having given email")
	}

	usersResponse := make([]*genModel.UserResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, asUserResponse(user))
	}

	return usersResponse, nil
}

// DeleteUser delete user
func (s *service) DeleteUser(ctx context.Context, id string) error {
	user, err := s.db.FindUser(ctx, id)
	if err != nil {
		return wraperrors.Wrap(err, "failed to find user with given id")
	}

	if user == nil {
		return wraperrors.Wrap(domainErr.ErrNoContent, "no user found")
	}

	if err := s.db.DeleteUser(ctx, id); err != nil {
		return wraperrors.Wrap(err, "failed to delete user")
	}

	return nil
}

// UpdateUser update user's name and address
func (s *service) UpdateUser(ctx context.Context, user *domain.User) (*genModel.UserResponse, error) {
	usr, err := s.db.FindUser(ctx, user.ID)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to find user with given id")
	}

	if usr == nil {
		return nil, wraperrors.Wrap(domainErr.ErrNoContent, "no user found")
	}

	if user.Name != "" {
		usr.Name = user.Name
	}
	if user.Address != "" {
		usr.Address = user.Address
	}

	usr, err = s.db.SaveUser(ctx, usr)
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to delete user")
	}

	tmp := asUserResponse(usr)
	return tmp, nil
}
