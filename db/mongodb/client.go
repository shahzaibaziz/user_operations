package mongodb

import (
	"context"

	wraperrors "github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/shahzaibaziz/user_operations/config"
	"github.com/shahzaibaziz/user_operations/db"
)

type client struct {
	dbc *mongo.Database
}

// NewClient creates a configured database client.
func NewClient(option db.Option) (db.DataStore, error) {
	ctx := context.Background()

	clientOpts, err := getClientOptions()
	if err != nil {
		return nil, err
	}

	dbClient, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, wraperrors.Wrap(err, "unable to connect to database")
	}

	if err := dbClient.Ping(ctx, nil); err != nil {
		return nil, wraperrors.Wrap(err, "unable to ping database")
	}

	return &client{dbc: dbClient.Database(viper.GetString(config.MongoDBName))}, nil
}

func getClientOptions() (*options.ClientOptions, error) {
	clientOpts := options.Client()

	if creds := getCredentials(); creds != nil {
		clientOpts.SetAuth(*creds)
	}

	if err := clientOpts.Validate(); err != nil {
		return nil, wraperrors.Wrap(err, "unable to validate database client configurations")
	}

	return clientOpts, nil
}

// getCredentials provides the necessary credentials to establish a connection.
func getCredentials() *options.Credential {
	if viper.GetString(config.MongoDBUsername) != "" && viper.GetString(config.MongoDBPassword) != "" {
		return &options.Credential{
			Username: viper.GetString(config.MongoDBUsername),
			Password: viper.GetString(config.MongoDBPassword),
		}
	}

	return nil
}
