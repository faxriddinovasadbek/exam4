package grpcClient

import (
	"fmt"
	"post-service/config"
	pbc "post-service/protos/comment-service"
	pbu "post-service/protos/user-service"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	CommentService() pbc.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	userService    pbu.UserServiceClient
	commentService pbc.CommentServiceClient
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
	// dail to comment-service
	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("comment service dail host: %s port : %d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}
	return &serviceManager{
		cfg:            cfg,
		userService:    pbu.NewUserServiceClient(connUser),
		commentService: pbc.NewCommentServiceClient(connComment),
	}, nil
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}
