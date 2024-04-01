package service

import (
	"context"
	l "user-service/pkg/logger"

	// pbp "user-service/protos/post-service"
	pbu "user-service/protos/user-service"
	"user-service/storage"

	grpcClient "user-service/service/grpc_client"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// mongo
func NewUserServiceMongo(db *mongo.Collection, log l.Logger, client grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}


func (s *UserService) GetUserByRefreshToken(ctx context.Context, refresh *pbu.RefreshToken) (*pbu.User, error) {
	user, err := s.storage.User().GetUserByRefreshToken(refresh)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email *pbu.ByEmail) (*pbu.User, error) {
	user, err := s.storage.User().GetUserByEmail(email)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) CheckUniques(ctx context.Context, req *pbu.CheckUniquesRequest) (*pbu.CheckUniquesResponse, error) {
	check, err := s.storage.User().CheckUniques(req)
	if err != nil {
		return nil, err
	}

	return check, nil
}

func (s *UserService) Create(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pbu.UserRequest) (*pbu.CheckUniquesResponse, error) {
	user, err := s.storage.User().Delete(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pbu.UserRequest) (*pbu.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	users, err := s.storage.User().GetAll(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return users, nil
}

