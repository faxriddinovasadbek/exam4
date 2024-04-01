package storage

import (
	"comment-service/storage/mongodb"
	"comment-service/storage/postgres"
	"comment-service/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.CommentStorageI
}

type storageMongo struct {
	db       *mongo.Collection
	postRepo repo.CommentStorageI
}

func (s storageMongo) Comment() repo.CommentStorageI {
	return s.postRepo
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.postRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewCommentRepo(db)}
}
func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{db, mongodb.NewPostRepoMongo(db)}
}
