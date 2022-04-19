package db

import (
	"context"
	"log"

	"github.com/shahzaibaziz/user_operations/models"
)

// DataStore is an interface for query ops
type DataStore interface {
	DeleteUser(ctx context.Context, id string) error
	FindUserByFilter(ctx context.Context, filter map[string]interface{}) (*models.User, error)
	FindUser(ctx context.Context, id string) (*models.User, error)
	GetUsers(ctx context.Context, filter map[string]interface{}) ([]*models.User, error)
	SaveUser(ctx context.Context, user *models.User) (*models.User, error)
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	if _, ok := datastoreFactories[name]; ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
