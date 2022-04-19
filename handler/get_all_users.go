package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/shahzaibaziz/user_operations"
	genModel "github.com/shahzaibaziz/user_operations/gen/models"
	"github.com/shahzaibaziz/user_operations/gen/restapi/operations/service"
)

// NewGetAllUsers custom handler for get all users
func NewGetAllUsers(rt *runtime.Runtime) service.GetAllUsersHandler {
	return &allUsers{rt: rt}
}

type allUsers struct {
	rt *runtime.Runtime
}

// Handle implementation of custom handler
func (r *allUsers) Handle(params service.GetAllUsersParams) middleware.Responder {
	ctx := context.Background()
	filter := make(map[string]interface{})
	if params.Limit != nil {
		filter["limit"] = params.Limit
	}
	if params.Name != nil {
		filter["name"] = params.Name
	}
	users, err := r.rt.Service().GetUsers(ctx, filter)
	if err != nil {
		log(ctx).Errorf("failed to get users", err)

		return service.NewGetAllUsersDefault(http.StatusInternalServerError).WithPayload(
			&genModel.Error{
				Code:    swag.String(fmt.Sprintf("%v", http.StatusInternalServerError)),
				Message: swag.String(err.Error()),
			})
	}

	log(ctx).Infof("found users %v", len(users))
	return service.NewGetAllUsersOK().WithPayload(users)
}
