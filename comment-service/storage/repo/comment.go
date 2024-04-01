package repo

import (
	pbc "comment-service/protos/comment-service"
)

// CommentStorageI ...
type CommentStorageI interface {
	CreateCommment(*pbc.Comment) (*pbc.Comment, error)
	UpdateComment(*pbc.Comment) (*pbc.Comment, error)
	DeleteComment(commentId string) error
	GetComment(commentId string) (*pbc.Comment, error)
	GetAllComment(page, limit int64) ([]*pbc.Comment, error)
	GetAllCommentsByPostId(postId string) ([]*pbc.Comment, error)
	GetAllCommentsByOwnerId(ownerId string) ([]*pbc.Comment, error)
	GetCommentsById(id string) (*pbc.Comment, error)
}
