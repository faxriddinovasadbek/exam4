package storage

import (
	"post-service/storage/mongodb"
	"post-service/storage/postgres"
	"post-service/storage/repo"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

type storageMongo struct {
	db       *mongo.Collection
	postRepo repo.PostStorageI
}

func (s storageMongo) Post() repo.PostStorageI {
	return s.postRepo
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewPostRepo(db)}
}
func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{db, mongodb.NewPostRepoMongo(db)}
}
