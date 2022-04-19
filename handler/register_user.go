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

// NewRegisterUserHandler custom handler of save new User
func NewRegisterUserHandler(rt *runtime.Runtime) service.RegisterUserHandler {
	return &registerUser{rt: rt}
}

type registerUser struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *registerUser) Handle(params service.RegisterUserParams) middleware.Responder {
	ctx := context.Background()

	user, err := r.rt.Service().SaveUser(ctx, toUserModel(params.User))
	if errors.Is(err, domainErr.ErrConflict) {
		log(ctx).Errorf("user with given email is already exist in database", err)

		return service.NewRegisterUserConflict().WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusConflict)),
			Message: swag.String(err.Error()),
		})
	} else if err != nil {
		log(ctx).Errorf("failed to register new user", err)

		return service.NewRegisterUserDefault(http.StatusInternalServerError).WithPayload(&genModel.Error{
			Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
			Message: swag.String(err.Error()),
		})
	}

	log(ctx).Infof("got user %v", user)
	return service.NewRegisterUserCreated().WithPayload(user)
}
