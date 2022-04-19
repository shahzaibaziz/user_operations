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

// NewDeleteUserHandler custom handler for delete user
func NewDeleteUserHandler(rt *runtime.Runtime) service.DeleteUserHandler {
	return &deleteUser{rt: rt}
}

type deleteUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *deleteUser) Handle(params service.DeleteUserParams) middleware.Responder {
	ctx := context.Background()

	if err := r.rt.Service().DeleteUser(ctx, params.UserID); errors.Is(err, domainErr.ErrNoContent) {
		log(ctx).Errorf("failed to find user with given id ", err)

		return service.NewDeleteUserNotFound().WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusNotFound)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to delete  user ", err)

		return service.NewDeleteUserDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("user delete %v", params.UserID)
	return service.NewDeleteUserNoContent()
}
