package service

import (
	"context"
	// "fmt"
	l "post-service/pkg/logger"
	pbp "post-service/protos/post-service"
	pbu "post-service/protos/user-service"
	"post-service/storage"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"

	grpcClient "post-service/service/grpc_client"
)

// PostService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// mongo
func NewPostServiceMongo(db *mongo.Collection, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}

func (s *PostService) Create(ctx context.Context, req *pbp.Post) (*pbp.Post, error) {
	post, err := s.storage.Post().Create(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

func (s *PostService) Update(ctx context.Context, req *pbp.Post) (*pbp.Post, error) {
	post, err := s.storage.Post().Update(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

func (s *PostService) Delete(ctx context.Context, req *pbp.GetRequest) (*pbp.CheckResponse, error) {
	user, err := s.storage.Post().Delete(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *PostService) GetPost(ctx context.Context, req *pbp.GetRequest) (*pbp.PostResponse, error) {
	post, err := s.storage.Post().GetPost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	user, err := s.client.UserService().Get(ctx, &pbu.UserRequest{UserId: post.OwnerId})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &pbp.PostResponse{
		Id:        post.Id,
		Content:   post.Content,
		Title:     post.Title,
		Likes:     post.Likes,
		Dislikes:  post.Dislikes,
		Views:     post.Views,
		Category:  post.Category,
		OwnerId:   post.OwnerId,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Owner: &pbp.Owner{
			Id:        user.Id,
			Name:      user.Name,
			LastName:  user.LastName,
			Username:  user.Username,
			Email:     user.Email,
			Bio:       user.Bio,
			Website:   user.Website,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

func (s *PostService) GetAllPosts(ctx context.Context, req *pbp.GetAllPostsRequest) (*pbp.GetPostsByOwnerIdResponse, error) {
	posts, err := s.storage.Post().GetAllPosts(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPostsByOwnerId(ctx context.Context, req *pbp.GetPostsByOwnerIdRequest) (*pbp.GetPostsByOwnerIdResponse, error) {
	posts, err := s.storage.Post().GetPostsByOwnerId(&pbp.GetPostsByOwnerIdRequest{OwnerId: req.OwnerId})
	if err != nil {
		return nil, err
	}

	return posts, nil
}
