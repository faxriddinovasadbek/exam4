package v1

import (
	"api-gateway/api/handlers/tokens"
	"api-gateway/config"
	"api-gateway/pkg/logger"
	"api-gateway/services"

	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     tokens.JWTHandler
	enforcer       *casbin.Enforcer
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     tokens.JWTHandler
	Enforcer       *casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
		enforcer:       c.Enforcer,
	}
}
