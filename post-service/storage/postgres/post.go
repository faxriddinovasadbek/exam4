package postgres

import (
	pbp "post-service/protos/post-service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *pbp.Post) (*pbp.Post, error) {

	query := `
		INSERT INTO posts (
			id, 
			content, 
			title, 
			likes,
			dislikes,
			views,
			category,
			owner_id
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING 
			id, 
			content, 
			title, 
			likes,
			dislikes,
			views,
			category,
			owner_id,
			created_at,
			updated_at
	`

	id := uuid.New().String()
	if err := r.db.QueryRow(query, id, post.Content, post.Title, 0, 0, 0, post.Category, post.OwnerId).Scan(
		&post.Id,
		&post.Content,
		&post.Title,
		&post.Likes,
		&post.Dislikes,
		&post.Views,
		&post.Category,
		&post.OwnerId,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return post, nil
}

func (r *postRepo) Update(post *pbp.Post) (*pbp.Post, error) {
	query := `
	UPDATE
    	posts
	SET
    	updated_at = CURRENT_TIMESTAMP,
    	content = $1, 
    	title = $2, 
    	likes = $3,
    	dislikes = $4,
    	views = $5,
    	category = $6,
    	owner_id = $7
	WHERE
    	id = $8
    AND 
		deleted_at IS NULL
	RETURNING
    	id, 
    	content, 
    	title, 
    	likes,
    	dislikes,
    	views,
    	category,
    	owner_id,
    	created_at,
    	updated_at
	`
	var response pbp.Post
	if err := r.db.QueryRow(query, post.Content, post.Title, post.Likes, post.Dislikes, post.Views, post.Category, post.OwnerId, post.Id).Scan(
		&response.Id,
		&response.Content,
		&response.Title,
		&response.Likes,
		&response.Dislikes,
		&response.Views,
		&response.Category,
		&response.OwnerId,
		&response.CreatedAt,
		&response.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *postRepo) Delete(post *pbp.GetRequest) (*pbp.CheckResponse, error) {
	queryDel := `
	UPDATE
		posts
	SET
		deleted_at = CURRENT_TIMESTAMP
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	num, err := r.db.Exec(queryDel, post.PostId)
	if err != nil {
		return nil, err
	}

	son, err := num.RowsAffected()
	if err != nil {
		return nil, err
	}

	if son == 0 {
		return &pbp.CheckResponse{Chack: false}, nil
	}

	return &pbp.CheckResponse{Chack: true}, nil

}
func (r *postRepo) GetPost(post *pbp.GetRequest) (*pbp.Post, error) {
	query := `
	SELECT 
		id, 
		content, 
		title, 
		likes,
		dislikes,
		views,
		category,
		owner_id,
		created_at,
		updated_at
	FROM 
		posts
	WHERE 
		id = $1
	AND 
		deleted_at IS NULL
	`
	var respPost pbp.Post
	if err := r.db.QueryRow(query, post.PostId).Scan(
		&respPost.Id,
		&respPost.Content,
		&respPost.Title,
		&respPost.Likes,
		&respPost.Dislikes,
		&respPost.Views,
		&respPost.Category,
		&respPost.OwnerId,
		&respPost.CreatedAt,
		&respPost.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &respPost, nil
}

func (r *postRepo) GetAllPosts(post *pbp.GetAllPostsRequest) (*pbp.GetPostsByOwnerIdResponse, error) {
	offset := post.Limit * (post.Page - 1)
	query := `
	SELECT 
		id, 
		content, 
		title, 
		likes,
		dislikes,
		views,
		category,
		owner_id,
		created_at,
		updated_at
	FROM
		posts
	WHERE
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`
	rows, err := r.db.Query(query, post.Limit, offset)
	if err != nil {
		return nil, err
	}
	var posts pbp.GetPostsByOwnerIdResponse
	for rows.Next() {
		var post pbp.Post
		if err := rows.Scan(
			&post.Id,
			&post.Content,
			&post.Title,
			&post.Likes,
			&post.Dislikes,
			&post.Views,
			&post.Category,
			&post.OwnerId,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}
		posts.Posts = append(posts.Posts, &post)
	}
	return &posts, nil
}

func (r *postRepo) GetPostsByOwnerId(req *pbp.GetPostsByOwnerIdRequest) (*pbp.GetPostsByOwnerIdResponse, error) {
	query := `
	SELECT 
		id, 
		content, 
		title, 
		likes,
		dislikes,
		views,
		category,
		owner_id,
		created_at,
		updated_at
	FROM 
		posts 
	WHERE 
		owner_id = $1
	`
	rows, err := r.db.Query(query, req.OwnerId)
	if err != nil {
		return nil, err
	}
	var response pbp.GetPostsByOwnerIdResponse
	for rows.Next() {
		var post pbp.Post
		if err := rows.Scan(
			&post.Id,
			&post.Content,
			&post.Title,
			&post.Likes,
			&post.Dislikes,
			&post.Views,
			&post.Category,
			&post.OwnerId,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}
		response.Posts = append(response.Posts, &post)
	}

	return &response, nil
}
