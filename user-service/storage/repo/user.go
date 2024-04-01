package repo

import (
	pb "user-service/protos/user-service"
)

// UserStorageI ...
type UserStorageI interface {
	Create(user *pb.User) (*pb.User, error)
	Update(request *pb.User) (*pb.User, error)
	Delete(request *pb.UserRequest) (*pb.CheckUniquesResponse, error)
	Get(request *pb.UserRequest) (*pb.User, error)
	GetAll(request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error)
	CheckUniques(req *pb.CheckUniquesRequest) (*pb.CheckUniquesResponse, error)
	GetUserByEmail(email *pb.ByEmail) (*pb.User, error)
	GetUserByRefreshToken(refreshToken *pb.RefreshToken) (*pb.User, error)
}
