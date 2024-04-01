package mongodb

import (
	"context"
	"fmt"
	"time"

	pb "user-service/protos/user-service"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	mdb *mongo.Collection
}

func (u *userRepo) Create(user *pb.User) (*pb.User, error) {
	id := uuid.New().String()
	update := bson.M{
		"name":         user.Name,
		"lastname":     user.LastName,
		"username":     user.Username,
		"password":     user.Password,
		"email":        user.Email,
		"bio":          user.Bio,
		"website":      user.Website,
		"id":           id,
		"refreshtoken": user.RefreshToken,
		"createdat":    time.Now().String(),
		"updatedat":    time.Now().String(),
	}
	_, err := u.mdb.InsertOne(context.Background(), update)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": id}
	var user1 pb.User
	err = u.mdb.FindOne(context.Background(), filter).Decode(&user1)
	if err != nil {
		return nil, err
	}
	fmt.Println(err)
	return &user1, nil
}

// Get(user_id *pbu.UserRequest) (*pbu.User, error)
func (u *userRepo) Get(id *pb.UserRequest) (*pb.User, error) {
	filter := bson.M{"id": id.UserId}
	var user pb.User
	err := u.mdb.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetAll(req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	filter := bson.M{}
	options := options.Find().SetSkip(req.Page).SetLimit(req.Limit)
	cursor, err := u.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users pb.GetAllUsersResponse

	for cursor.Next(context.Background()) {
		var user pb.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		pbUser := &pb.User{
			Name:         user.Name,
			LastName:     user.LastName,
			Username:     user.Username,
			Password:     user.Password,
			Email:        user.Email,
			Bio:          user.Bio,
			Website:      user.Website,
			Id:           user.Id,
			RefreshToken: user.RefreshToken,
		}
		users.AllUsers = append(users.AllUsers, pbUser)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *userRepo) Update(req *pb.User) (*pb.User, error) {
	update := bson.M{
		"$set": bson.M{
			"name":         req.Name,
			"lastname":     req.LastName,
			"username":     req.Username,
			"bio":          req.Bio,
			"website":      req.Website,
			"password":     req.Password,
			"email":        req.Email,
			"id":           req.Id,
			"refreshtoken": req.RefreshToken,
		},
	}

	filter := bson.M{"id": req.Id}
	updateResult, err := u.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("user with id %v not found", req.Id)
	}
	updatedUser, err := u.Get(&pb.UserRequest{UserId: req.Id})
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *userRepo) Delete(user_id *pb.UserRequest) (*pb.CheckUniquesResponse, error) {
	filter := bson.M{"id": user_id.UserId}
	result, err := u.mdb.DeleteOne(context.Background(), filter)
	if err != nil {
		return &pb.CheckUniquesResponse{Check: false}, err
	}
	if result.DeletedCount == 0 {
		return &pb.CheckUniquesResponse{Check: false}, err
	}
	return &pb.CheckUniquesResponse{Check: true}, err
}

func (u *userRepo) CheckUniques(req *pb.CheckUniquesRequest) (*pb.CheckUniquesResponse, error) {

	count, err := u.mdb.CountDocuments(context.Background(), bson.M{req.Field: req.Value}, &options.CountOptions{})
	if err != nil {
		return &pb.CheckUniquesResponse{Check: false}, err
	}

	if count != 0 {
		return &pb.CheckUniquesResponse{Check: true}, err
	}

	return &pb.CheckUniquesResponse{Check: false}, nil
}



func (u *userRepo) GetUserByEmail(req *pb.ByEmail) (*pb.User, error) {
	var user pb.User
	filter := bson.M{"email": req.Email}
	err := u.mdb.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	pbUser := &pb.User{
		Name:         user.Name,
		LastName:     user.LastName,
		Username:     user.Username,
		Password:     user.Password,
		Email:        user.Email,
		Bio:          user.Bio,
		Website:      user.Website,
		Id:           user.Id,
		RefreshToken: user.RefreshToken,
	}
	return pbUser, nil
}

func (r *userRepo) GetUserByRefreshToken(refresh *pb.RefreshToken) (*pb.User, error) {
	return nil, nil
}

func NewUserRepoMongo(db *mongo.Collection) *userRepo {
	return &userRepo{mdb: db}
}
