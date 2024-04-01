package mongodb

import (
	"context"
	"fmt"
	"time"

	pb "comment-service/protos/comment-service"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	mdb *mongo.Collection
}

func (u *userRepo) CreateCommment(comment *pb.Comment) (*pb.Comment, error) {
	id := uuid.New().String()
	update := bson.M{
		"id":        id,
		"content":   comment.Content,
		"postid":    comment.PostId,
		"userid":    comment.OwnerId,
		"createdat": time.Now().String(),
		"updatedat": time.Now().String(),
	}

	_, err := u.mdb.InsertOne(context.Background(), update)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"id": id}
	var comment1 pb.Comment
	err = u.mdb.FindOne(context.Background(), filter).Decode(&comment1)
	if err != nil {
		return nil, err
	}
	return &comment1, nil
}

func (u *userRepo) GetComment(commentId string) (*pb.Comment, error) {
	filter := bson.M{"id": commentId}
	var comment pb.Comment
	err := u.mdb.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (u *userRepo) GetAllComment(page, limit int64) ([]*pb.Comment, error) {
	filter := bson.M{}
	options := options.Find().SetSkip(page).SetLimit(limit)
	cursor, err := u.mdb.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var comments []*pb.Comment

	for cursor.Next(context.Background()) {
		var comment pb.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}

		pbcomment := &pb.Comment{
			Id:        comment.Id,
			Content:   comment.Content,
			PostId:    comment.PostId,
			OwnerId:   comment.OwnerId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		// fmt.Println(pbcomment)
		comments = append(comments, pbcomment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (u *userRepo) UpdateComment(comment *pb.Comment) (*pb.Comment, error) {
	update := bson.M{
		"$set": bson.M{
			"id":        comment.Id,
			"content":   comment.Content,
			"postid":    comment.PostId,
			"userid":    comment.OwnerId,
			"updatedat": time.Now().String(),
		},
	}
	filter := bson.M{"id": comment.Id}
	updateResult, err := u.mdb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("post with id %v not found", comment.Id)
	}
	updatedUser, err := u.GetComment(comment.Id)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *userRepo) DeleteComment(commendId string) error {
	filter := bson.M{"id": commendId}
	result, err := u.mdb.DeleteOne(context.Background(), filter)

	if err != nil {
		return  err
	}
	if result.DeletedCount == 0 {
		return  nil
	}
	return err
}

func (u *userRepo) GetAllCommentsByPostId(postId string) ([]*pb.Comment, error) {
	filter := bson.M{"id": postId}
	cursor, err := u.mdb.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var comments []*pb.Comment

	for cursor.Next(context.Background()) {
		var comment pb.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}

		pbcomment := &pb.Comment{
			Id:        comment.Id,
			Content:   comment.Content,
			PostId:    comment.PostId,
			OwnerId:   comment.OwnerId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		comments = append(comments, pbcomment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (u *userRepo) GetAllCommentsByOwnerId(ownerId string) ([]*pb.Comment, error) {
	filter := bson.M{"id": ownerId}
	cursor, err := u.mdb.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var comments []*pb.Comment

	for cursor.Next(context.Background()) {
		var comment pb.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}

		pbcomment := &pb.Comment{
			Id:        comment.Id,
			Content:   comment.Content,
			PostId:    comment.PostId,
			OwnerId:   comment.OwnerId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}

		comments = append(comments, pbcomment)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (u *userRepo)GetCommentsById(id string) (*pb.Comment, error) {
	filter := bson.M{"id": id}
	var comment pb.Comment
	err := u.mdb.FindOne(context.Background(), filter).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}


func NewPostRepoMongo(db *mongo.Collection) *userRepo {
	return &userRepo{mdb: db}
}
