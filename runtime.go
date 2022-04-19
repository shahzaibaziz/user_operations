package runtime

import (
	wraperrors "github.com/pkg/errors"

	"github.com/shahzaibaziz/user_operations/db"
	"github.com/shahzaibaziz/user_operations/db/mongodb"
	"github.com/shahzaibaziz/user_operations/service"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	dbc     db.DataStore
	service service.Manager
}

// NewRuntime creates a new runtime
func NewRuntime() (*Runtime, error) {
	client, err := mongodb.NewClient(db.Option{})
	if err != nil {
		return nil, wraperrors.Wrap(err, "failed to connect with database")
	}
	return &Runtime{service: service.NewService(&client)}, nil
}

// Service return  service layer object
func (rt *Runtime) Service() service.Manager {
	return rt.service
}
