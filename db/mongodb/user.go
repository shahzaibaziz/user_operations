package mongodb

import (
	"context"

	wraperrors "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/shahzaibaziz/user_operations/models"
)

const (
	// UserCollection mongodb  collections.
	UserCollection = "users"
	// Mongodb field names.
	pk = "_id"
	// Filter map keys.
	idRef   = "id"
	nameRef = "name"
	// Options map keys.
	limitRef = "limit"
)

// SaveUser save user object into database
func (c *client) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	if _, err := c.dbc.Collection(UserCollection).ReplaceOne(
		ctx,
		bson.D{{Key: pk, Value: user.ID}},
		user,
		options.Replace().SetUpsert(true)); err != nil {
		return nil, wraperrors.Wrap(err, "unable to save user")
	}

	log(ctx).Infof("saved user, %+v", user)

	return user, nil
}

// FindUserByFilter find one user bas on given filter
func (c *client) FindUserByFilter(ctx context.Context, filter map[string]interface{}) (*domain.User, error) {
	var user *domain.User
	if err := c.dbc.Collection(UserCollection).FindOne(ctx, applyFilter(filter)).Decode(&user); err != nil {
		if wraperrors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, wraperrors.Wrap(err, "unable to find subscription")
	}
	log(ctx).Infof("found %v user", user)

	return user, nil
}

// FindUser find one user  by id
func (c *client) FindUser(ctx context.Context, id string) (*domain.User, error) {
	var user *domain.User
	if err := c.dbc.Collection(UserCollection).FindOne(ctx, bson.D{{Key: pk, Value: id}}).Decode(&user); err != nil {
		if wraperrors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, wraperrors.Wrap(err, "unable to find user")
	}

	log(ctx).Infof("found user, %+v", user)

	return user, nil
}

// GetUsers find all users base on given filters
func (c *client) GetUsers(ctx context.Context, filter map[string]interface{}) ([]*domain.User, error) {
	users := make([]*domain.User, 0)

	cur, err := c.dbc.Collection(UserCollection).Find(ctx, applyFilter(filter), applyOptions(filter))
	if err != nil {
		return nil, wraperrors.Wrap(err, "unable to find users")
	}
	if err := cur.All(ctx, &users); err != nil {
		return nil, wraperrors.Wrap(err, "unable to decode user records")
	}

	log(ctx).Infof("found %v users", len(users))

	return users, nil
}

// DeleteUser delete user using id
func (c *client) DeleteUser(ctx context.Context, id string) error {
	if _, err := c.dbc.Collection(UserCollection).DeleteOne(
		ctx,
		bson.D{{Key: pk, Value: id}},
	); err != nil {
		return wraperrors.Wrap(err, "unable to delete service instance")
	}

	log(ctx).Infof("deleted user %+v", id)

	return nil
}

func applyFilter(filter map[string]interface{}) map[string]interface{} {
	if val, ok := filter[idRef].(string); ok {
		// update filter with the correct field name and value.
		filter[pk] = val
		delete(filter, idRef)
	}

	return filter
}

func applyOptions(filter map[string]interface{}) *options.FindOptions {
	opts := options.Find().SetSort(bson.D{{nameRef, 1}})
	if value, ok := filter[limitRef].(*int64); ok {
		filter["limit"] = value
		opts.SetLimit(*value)
		delete(filter, limitRef)
	}

	return opts
}
