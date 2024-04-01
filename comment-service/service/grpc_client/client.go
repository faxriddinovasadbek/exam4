package grpcClient

import (
	"comment-service/config"
	pbp "comment-service/protos/post-service"
	pbu "comment-service/protos/user-service"
	"fmt"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	PostService() pbp.PostServiceClient
}

type serviceManager struct {
	cfg         config.Config
	userService pbu.UserServiceClient
	postService pbp.PostServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to user-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.UserServiceHost, cfg.UserServicePort)
	}
	// dail to post-service
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("post service dail host: %s port : %d", cfg.PostServiceHost, cfg.PostServicePort)
	}
	return &serviceManager{
		cfg:         cfg,
		userService: pbu.NewUserServiceClient(connUser),
		postService: pbp.NewPostServiceClient(connPost),
	}, nil
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}
