package postgres

import (
	pbc "comment-service/protos/comment-service"

	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

// NewCommentRepo ...
func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{db: db}
}

func (c *commentRepo) CreateCommment(comment *pbc.Comment) (*pbc.Comment, error) {
	query := `
	INSERT INTO comments (
		content,
		post_id,
		user_id
	)
	VALUES ($1, $2, $3)
	RETURNING
		id,
		content,
		post_id,
		user_id,
		created_at,
		updated_at
	`
	createdComment := pbc.Comment{}
	if err := c.db.QueryRow(query, comment.Content, comment.PostId, comment.OwnerId).Scan(
		&createdComment.Id,
		&createdComment.Content,
		&createdComment.PostId,
		&createdComment.OwnerId,
		&createdComment.CreatedAt,
		&createdComment.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &createdComment, nil
}

func (c *commentRepo) UpdateComment(comment *pbc.Comment) (*pbc.Comment, error) {
	query := `
	UPDATE
		comments
	SET
		content = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $2
	RETURNING
		id,
		content,
		post_id,
		user_id,
		created_at,
		updated_at
	`
	var updatedComment pbc.Comment
	if err := c.db.QueryRow(query, comment.Content, comment.Id).Scan(
		&updatedComment.Id,
		&updatedComment.Content,
		&updatedComment.PostId,
		&updatedComment.OwnerId,
		&updatedComment.CreatedAt,
		&updatedComment.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &updatedComment, nil
}

func (c *commentRepo) DeleteComment(commendId string) error {
	query := `
	DELETE FROM 
		comments
	WHERE
		id = $1
	`
	_, err := c.db.Exec(query, commendId)
	if err != nil {
		return err
	}
	return nil
}

func (c *commentRepo) GetComment(commentId string) (*pbc.Comment, error) {
	query := `
	SELECT 
		id, 
		content, 
		post_id, 
		user_id,
		created_at, 
		updated_at
	FROM 
		comments 
	WHERE
		id = $1
	`
	var commentResponse pbc.Comment
	if err := c.db.QueryRow(query, commentId).Scan(
		&commentResponse.Id,
		&commentResponse.Content,
		&commentResponse.PostId,
		&commentResponse.OwnerId,
		&commentResponse.CreatedAt,
		&commentResponse.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &commentResponse, nil
}

func (c *commentRepo) GetAllComment(page, limit int64) ([]*pbc.Comment, error) {
	offset := limit * (page - 1)
	query := `
	SELECT
		id,
		content,
		post_id,
		user_id,
		created_at,
		updated_at
	FROM
		comments
	LIMIT $1
	OFFSET $2
	`
	rows, err := c.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	var comments []*pbc.Comment
	for rows.Next() {
		var comment pbc.Comment
		if err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.PostId,
			&comment.OwnerId,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c *commentRepo) GetAllCommentsByPostId(postId string) ([]*pbc.Comment, error) {
	query := `
	SELECT 
		id, 
		content, 
		post_id, 
		user_id,
		created_at, 
		updated_at
	FROM 
		comments 
	WHERE 
		post_id = $1
	`
	rows, err := c.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	var comments []*pbc.Comment
	for rows.Next() {
		var comment pbc.Comment
		if err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.PostId,
			&comment.OwnerId,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c *commentRepo) GetAllCommentsByOwnerId(ownerId string) ([]*pbc.Comment, error) {
	query := `
	SELECT 
		id, 
		content, 
		post_id, 
		user_id,
		created_at, 
		updated_at
	FROM
		comments
	WHERE 
		user_id = $1
	`
	rows, err := c.db.Query(query, ownerId)
	if err != nil {
		return nil, err
	}
	var comments []*pbc.Comment
	for rows.Next() {
		var comment pbc.Comment
		if err := rows.Scan(
			&comment.Id,
			&comment.Content,
			&comment.PostId,
			&comment.OwnerId,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (c *commentRepo) GetCommentsById(id string) (*pbc.Comment, error) {
	query := `
	SELECT 
		id, 
		content, 
		post_id, 
		user_id,
		created_at, 
		updated_at
	FROM
		comments
	WHERE 
		id = $1
	`
	var comment pbc.Comment
	if err := c.db.QueryRow(query, id).Scan(
		&comment.Id,
		&comment.Content,
		&comment.PostId,
		&comment.OwnerId,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &comment, nil
}
