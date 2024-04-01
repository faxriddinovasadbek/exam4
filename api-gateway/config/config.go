package config

import (
	"os"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	// another services
	UserServiceHost    string
	UserServicePort    int
	PostServiceHost    string
	PostServicePort    int
	CommentServiceHost string
	CommentServicePort int

	// context timeout in seconds
	CtxTimeout int

	// casbin
	AuthConfigPath string
	CSVFilePath    string

	// JWT
	SigningKey        string
	AccessTokenTimout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	// user serrvice bn connect
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 1111))

	// post service bn connect
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 2222))

	// comment service bn connect
	c.CommentServiceHost = cast.ToString(getOrReturnDefault("COMMENT_SERVICE_HOST", "localhost"))
	c.CommentServicePort = cast.ToInt(getOrReturnDefault("COMMENT_SERVICE_PORT", 3333))

	// Jwt
	c.SigningKey = cast.ToString(getOrReturnDefault("SIGNING_KEY", "test-key"))
	c.AccessTokenTimout = cast.ToInt(getOrReturnDefault("ACCESS_TOKEN_TIMOUT", 555555))

	// casbin
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATH", "auth.conf"))
	c.CSVFilePath = cast.ToString(getOrReturnDefault("CSV_FILE_PATH", "casbin_rules.csv"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
