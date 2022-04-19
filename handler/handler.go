package handler

import (
	runtime "github.com/shahzaibaziz/user_operations"
	"github.com/shahzaibaziz/user_operations/gen/restapi/operations"
)

// NewCustomHandler mapping custom handler to auto generated handler
func NewCustomHandler(api *operations.UserOperationsAPI, rt *runtime.Runtime) {
	// custom handler
	api.ServiceRegisterUserHandler = NewRegisterUserHandler(rt)
	api.ServiceGetUserHandler = NewGetUserHandler(rt)
	api.ServiceUpdateUserHandler = NewUpdateUserHandler(rt)
	api.ServiceGetAllUsersHandler = NewGetAllUsers(rt)
	api.ServiceDeleteUserHandler = NewDeleteUserHandler(rt)
}
