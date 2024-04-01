package mongodb

import (
	"context"
	"fmt"
	"time"

	pb "post-service/protos/post-service"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	mdb *mongo.Collection
}

func (u *userRepo) Create(post *pb.Post) (*pb.Post, error) {
	id := uuid.New().String()
	update := bson.M{
		"id":        id,
		"content":   post.Content,
		"title":     post.Title,
		"likes":     post.Likes,
		"dislikes":  post.Dislikes,
		"views":     post.Views,
		"category":  post.Category,
		"ownerid":   post.OwnerId,
		"createdat": time.Now().String(),
		"updatedat": time.Now().String(),
	}
	_, err := u.mdb.InsertOne(context.Background(), update)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": id}
	var post1 pb.Post
	err = u.mdb.FindOne(context.Background(), filter).Decode(&post1)
	if err != nil {
		return nil, err
	}
	return &post1, nil
}

// Get(user_id *pbu.UserRequest) (*pbu.User, error)
func (u *userRepo) GetPost(id *pb.GetRequest) (*pb.Post, error) {
	filter := bson.M{"id": id.PostId}
	var post pb.Post
	err := u.mdb.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (u *userRepo) GetAllPosts(post *pb.GetAllPostsRequest) (*pb.GetPostsByOwnerIdResponse, error) {
	filter := bson.M{}
	options := options.Find().SetSkip(post.Page).SetLimit(post.Limit)
	cursor, err := u.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts pb.GetPostsByOwnerIdResponse

	for cursor.Next(context.Background()) {
		var post pb.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		pbpost := &pb.Post{
			Id:        post.Id,
			Content:   post.Content,
			Title:     post.Title,
			Likes:     post.Likes,
			Dislikes:  post.Likes,
			Views:     post.Views,
			Category:  post.Category,
			OwnerId:   post.OwnerId,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		// fmt.Println(pbpost)
		posts.Posts = append(posts.Posts, pbpost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &posts, nil
}

func (u *userRepo) Update(req *pb.Post) (*pb.Post, error) {
	update := bson.M{
		"$set": bson.M{
			"id":        req.Id,
			"content":   req.Content,
			"title":     req.Title,
			"likes":     req.Likes,
			"dislikes":  req.Dislikes,
			"views":     req.Views,
			"category":  req.Category,
			"ownerid":   req.OwnerId,
			"updatedat": time.Now().String(),
		},
	}

	filter := bson.M{"id": req.Id}
	updateResult, err := u.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("post with id %v not found", req.Id)
	}
	updatedUser, err := u.GetPost(&pb.GetRequest{PostId: req.Id})
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *userRepo) Delete(post *pb.GetRequest) (*pb.CheckResponse, error) {
	filter := bson.M{"id": post.PostId}
	result, err := u.mdb.DeleteOne(context.Background(), filter)
	if err != nil {
		return &pb.CheckResponse{Chack: false}, err
	}
	if result.DeletedCount == 0 {
		return &pb.CheckResponse{Chack: false}, err
	}
	return &pb.CheckResponse{Chack: true}, err
}

func (u *userRepo) GetPostsByOwnerId(req *pb.GetPostsByOwnerIdRequest) (*pb.GetPostsByOwnerIdResponse, error) {
	filter := bson.M{"id": req.OwnerId}
	cursor, err := u.mdb.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts pb.GetPostsByOwnerIdResponse

	for cursor.Next(context.Background()) {
		var post pb.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		pbpost := &pb.Post{
			Id:        post.Id,
			Content:   post.Content,
			Title:     post.Title,
			Likes:     post.Likes,
			Dislikes:  post.Likes,
			Views:     post.Views,
			Category:  post.Category,
			OwnerId:   post.OwnerId,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		posts.Posts = append(posts.Posts, pbpost)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &posts, nil
}

func NewPostRepoMongo(db *mongo.Collection) *userRepo {
	return &userRepo{mdb: db}
}
