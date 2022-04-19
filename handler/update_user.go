package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/shahzaibaziz/user_operations"
	domainErr "github.com/shahzaibaziz/user_operations/errors"
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	"github.com/shahzaibaziz/user_operations/gen/restapi/operations/service"
	domain "github.com/shahzaibaziz/user_operations/models"
)

// NewUpdateUserHandler custom handler for update user
func NewUpdateUserHandler(rt *runtime.Runtime) service.UpdateUserHandler {
	return &updateUser{rt: rt}
}

type updateUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *updateUser) Handle(params service.UpdateUserParams) middleware.Responder {
	ctx := context.Background()
	user := &domain.User{ID: params.UserID}
	if params.User.Name == "" && params.User.Address == "" {
		log(ctx).Errorf("body is empty")

		return service.NewUpdateUserDefault(http.StatusBadRequest).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusBadRequest)),
			Message: swag.String("name or address is required in body"),
		})
	}

	if params.User.Name != "" {
		user.Name = params.User.Name
	}
	if params.User.Address != "" {
		user.Name = params.User.Address
	}

	userRes, err := r.rt.Service().UpdateUser(ctx, user)
	if errors.Is(err, domainErr.ErrNoContent) {
		log(ctx).Errorf("failed to find user ", err)

		return service.NewUpdateUserNotFound().WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusNotFound)),
				Message: swag.String(err.Error()),
			})
	} else if err != nil {
		log(ctx).Errorf("failed to update user ", err)
		return service.NewUpdateUserDefault(http.StatusInternalServerError).WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
				Message: swag.String(err.Error()),
			})
	}

	log(ctx).Infof("update user %v", params.UserID)
	return service.NewUpdateUserOK().WithPayload(userRes)
}
