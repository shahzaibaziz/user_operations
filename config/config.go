package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	MongoDBHost     = "mongo.db.host"
	MongoDBUsername = "mongo.db.username"
	MongoDBPassword = "mongo.db.password"
	MongoDBName     = "mongo_db_name"
)

func init() {
	// env var for db
	_ = viper.BindEnv(MongoDBHost, "MONGO_DB_HOSTS")
	_ = viper.BindEnv(MongoDBUsername, "MONGO_DB_USERNAME")
	_ = viper.BindEnv(MongoDBPassword, "MONGO_DB_PASSWORD")

	viper.SetDefault(MongoDBName, "userdb")
	viper.SetDefault(MongoDBHost, "mongodb:27017")
}
