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
)

// NewGetUserHandler custom handler for get  user
func NewGetUserHandler(rt *runtime.Runtime) service.GetUserHandler {
	return &getUser{rt: rt}
}

type getUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *getUser) Handle(params service.GetUserParams) middleware.Responder {
	ctx := context.Background()

	user, err := r.rt.Service().FindUser(ctx, params.UserID)
	if errors.Is(err, domainErr.ErrNoContent) {
		log(ctx).Errorf("failed to find user with given id", err)

		return service.NewGetUserNotFound().WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusNotFound)),
				Message: swag.String(err.Error()),
			})
	} else if err != nil {
		log(ctx).Errorf("failed to get user", err)

		return service.NewGetUserDefault(http.StatusInternalServerError).WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
				Message: swag.String(err.Error()),
			})
	}

	log(ctx).Infof("user found %v", user)
	return service.NewGetUserOK().WithPayload(user)
}
