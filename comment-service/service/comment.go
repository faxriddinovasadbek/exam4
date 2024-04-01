package service

import (
	l "comment-service/pkg/logger"
	pbc "comment-service/protos/comment-service"
	pbp "comment-service/protos/post-service"
	pbu "comment-service/protos/user-service"
	"comment-service/storage"
	"context"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"

	grpcClient "comment-service/service/grpc_client"
)

// CommentService ...
type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewCommentService ...
func NewCommentService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// mongo
func NewPostServiceMongo(db *mongo.Collection, log l.Logger, client grpcClient.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}

func (s *CommentService) GetAllUsers(ctx context.Context, req *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentsResponse, error) {
	var response pbc.GetAllCommentsResponse
	users, err := s.client.UserService().GetAll(ctx, &pbu.GetAllUsersRequest{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	for _, user := range users.AllUsers {
		userResp := pbc.Users{
			Id:           user.Id,
			Name:         user.Name,
			LastName:     user.LastName,
			Email:        user.Email,
			Bio:          user.Bio,
			Website:      user.Website,
			Password:     user.Password,
			RefreshToken: user.RefreshToken,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		}
		posts, err := s.client.PostService().GetPostsByOwnerId(ctx, &pbp.GetPostsByOwnerIdRequest{
			OwnerId: user.Id,
		})
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		for _, post := range posts.Posts {
			postResp := pbc.Posts{
				Id:        post.Id,
				UserId:    post.OwnerId,
				Content:   post.Content,
				Title:     post.Title,
				Likes:     post.Likes,
				Dislikes:  post.Dislikes,
				Views:     post.Views,
				Category:  post.Category,
				CreatedAt: post.CreatedAt,
				UpdatedAt: post.UpdatedAt,
			}
			comments, err := s.storage.Comment().GetAllCommentsByPostId(postResp.Id)
			if err != nil {
				s.logger.Error(err.Error())
				return nil, err
			}
			for _, comment := range comments {
				commentResp := pbc.Comments{
					Id:        comment.Id,
					Content:   comment.Content,
					PostId:    comment.PostId,
					OwnerId:   comment.OwnerId,
					CreatedAt: comment.CreatedAt,
					UpdatedAt: comment.UpdatedAt,
				}
				postResp.AllComments = append(postResp.AllComments, &commentResp)
			}
			userResp.AllPosts = append(userResp.AllPosts, &postResp)
		}
		response.AllUsers = append(response.AllUsers, &userResp)
	}
	return &response, nil
}

func (s *CommentService) GetPostById(ctx context.Context, req *pbc.GetPostByIdRequest) (*pbc.GetPostByIdResponse, error) {
	var response pbc.GetPostByIdResponse
	post, err := s.client.PostService().GetPost(ctx, &pbp.GetRequest{
		PostId: req.PostId,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	response.Post = &pbc.Post{
		Id:        post.Id,
		UserId:    post.OwnerId,
		Content:   post.Content,
		Title:     post.Title,
		Likes:     post.Likes,
		Dislikes:  post.Dislikes,
		Views:     post.Views,
		Category:  post.Category,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	owner, err := s.client.UserService().Get(ctx, &pbu.UserRequest{
		UserId: post.Owner.Id,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	userData := pbc.User{
		Id:       owner.Id,
		Name:     owner.Name,
		LastName: owner.LastName,
	}
	response.PostWriter = &userData
	comments, err := s.storage.Comment().GetAllCommentsByPostId(response.Post.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	for _, comment := range comments {
		var commentData pbc.Comment
		commentData.Id = comment.Id
		commentData.Content = comment.Content
		commentData.OwnerId = comment.OwnerId
		commentData.PostId = comment.PostId
		commentData.CreatedAt = comment.CreatedAt
		commentData.UpdatedAt = comment.UpdatedAt
		response.Comments = append(response.Comments, &commentData)
	}

	return &response, nil
}

func (s *CommentService) GetUserById(ctx context.Context, req *pbc.GetUserByIdRequest) (*pbc.GetUserByIdResponse, error) {
	var response pbc.GetUserByIdResponse
	owner, err := s.client.UserService().Get(ctx, &pbu.UserRequest{
		UserId: req.OwnerId,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	response.OwnerInfo = &pbc.User{
		Id:           owner.Id,
		Name:         owner.Name,
		LastName:     owner.LastName,
		Email:        owner.Email,
		Bio:          owner.Bio,
		Website:      owner.Website,
		Password:     owner.Password,
		RefreshToken: owner.RefreshToken,
		CreatedAt:    owner.CreatedAt,
		UpdatedAt:    owner.UpdatedAt,
	}
	posts, err := s.client.PostService().GetPostsByOwnerId(ctx, &pbp.GetPostsByOwnerIdRequest{
		OwnerId: response.OwnerInfo.Id,
	})
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	for _, post := range posts.Posts {
		comments, err := s.storage.Comment().GetAllCommentsByPostId(post.Id)
		if err != nil {
			s.logger.Error(err.Error())
			return nil, err
		}
		allCommentRes := []*pbc.Comments{}
		commentRes := pbc.Comments{}
		for _, com := range comments {
			commentRes.Id = com.Id
			commentRes.Content = com.Content
			commentRes.PostId = com.PostId
			commentRes.OwnerId = com.OwnerId
			commentRes.CreatedAt = com.CreatedAt
			commentRes.UpdatedAt = com.UpdatedAt
			commentWriterInfo, err := s.client.UserService().Get(ctx, &pbu.UserRequest{
				UserId: commentRes.OwnerId,
			})
			if err != nil {
				s.logger.Error(err.Error())
				return nil, err
			}

			commentRes.CommentWriter = &pbc.User{
				Id:           commentWriterInfo.Id,
				Name:         commentWriterInfo.Name,
				LastName:     commentWriterInfo.LastName,
				Username:     commentWriterInfo.Username,
				Email:        commentWriterInfo.Email,
				Bio:          commentWriterInfo.Bio,
				Website:      commentWriterInfo.Website,
				RefreshToken: commentWriterInfo.RefreshToken,
				CreatedAt:    commentWriterInfo.CreatedAt,
				UpdatedAt:    commentWriterInfo.UpdatedAt,
			}
			allCommentRes = append(allCommentRes, &commentRes)
		}
		response.AllPosts = append(response.AllPosts, &pbc.Posts{
			Id:          post.Id,
			UserId:      post.OwnerId,
			Content:     post.Content,
			Title:       post.Title,
			Likes:       post.Likes,
			Dislikes:    post.Dislikes,
			Views:       post.Views,
			Category:    post.Category,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			AllComments: allCommentRes,
		})
	}

	return &response, nil
}

func (s *CommentService) CreateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	comment, err := s.storage.Comment().CreateCommment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) UpdateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	comment, err := s.storage.Comment().UpdateComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, req *pbc.IdRequst) (*pbc.DeleteResponse, error) {
	if err := s.storage.Comment().DeleteComment(req.Id); err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pbc.DeleteResponse{}, nil
}

func (s *CommentService) GetComment(ctx context.Context, req *pbc.IdRequst) (*pbc.Comment, error) {
	comment, err := s.storage.Comment().GetComment(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) GetAllComment(ctx context.Context, req *pbc.GetAllCommentsRequest) (*pbc.GetAllCommentResponse, error) {
	comments, err := s.storage.Comment().GetAllComment(req.Page, req.Limit)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pbc.GetAllCommentResponse{
		AllComments: comments,
	}, nil
}

func (s *CommentService) GetCommentsByPostId(ctx context.Context, req *pbc.IdRequst) (*pbc.GetAllCommentResponse, error) {
	comments, err := s.storage.Comment().GetAllCommentsByPostId(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pbc.GetAllCommentResponse{
		AllComments: comments,
	}, nil
}

func (s *CommentService) GetCommentsByOwnerId(ctx context.Context, req *pbc.IdRequst) (*pbc.GetAllCommentResponse, error) {
	comments, err := s.storage.Comment().GetAllCommentsByOwnerId(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &pbc.GetAllCommentResponse{
		AllComments: comments,
	}, nil
}

func (s *CommentService) GetCommentById(ctx context.Context, req *pbc.IdRequst) (*pbc.Comment, error) {
	comments, err := s.storage.Comment().GetCommentsById(req.Id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comments, nil
}
