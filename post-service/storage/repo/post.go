package repo

import (
	pbp "post-service/protos/post-service"
)

// PostStorageI ...
type PostStorageI interface {
	Create(post *pbp.Post) (*pbp.Post, error)
	Update(post *pbp.Post) (*pbp.Post, error)
	Delete(req *pbp.GetRequest) (*pbp.CheckResponse, error)
	GetPost(request *pbp.GetRequest) (*pbp.Post, error)
	GetAllPosts(req *pbp.GetAllPostsRequest) (*pbp.GetPostsByOwnerIdResponse, error)
	GetPostsByOwnerId(req *pbp.GetPostsByOwnerIdRequest) (*pbp.GetPostsByOwnerIdResponse, error)
}
