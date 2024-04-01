package grpcClient

import (
	"fmt"
	"user-service/config"
	pbc "user-service/protos/comment-service"
	pbp "user-service/protos/post-service"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	PostService() pbp.PostServiceClient
	CommentService() pbc.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	postService    pbp.PostServiceClient
	commentService pbc.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to post-service
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostgresUser, cfg.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.PostServiceHost, cfg.PostServicePort)
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
		postService:    pbp.NewPostServiceClient(connPost),
		commentService: pbc.NewCommentServiceClient(connComment),
	}, nil
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}
