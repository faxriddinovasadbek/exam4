package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment      string // develop, staging, production
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
	// connect fields
	PostServiceHost    string
	PostServicePort    int
	CommentServiceHost string
	CommentServicePort int

	MongoDatabase string
	MongoHost     string
	MongoPort     int
	
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "db"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "examdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "asadbek"))

	// mongo
	c.MongoHost = cast.ToString(getOrReturnDefault("MONGODB_HOST", "mongodb"))
	c.MongoPort = cast.ToInt(getOrReturnDefault("MONGODB_PORT", 27017))
	c.MongoDatabase = cast.ToString(getOrReturnDefault("MONGO_DATABASE", "examdb"))


	// connect to post-service
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post-service"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_HOST", "2222"))

	// connect to comment-service
	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "comment-service"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", "3333"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":1111"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
