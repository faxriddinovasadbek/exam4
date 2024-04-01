package postgres

import (
	pbu "user-service/protos/user-service"

	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetUserByEmail(email *pbu.ByEmail) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT 
		id, 
		name, 
		last_name,
		username,
		email,
		password,
		bio,
		website,
		created_at,
		updated_at
	FROM 
		users 
	WHERE 
		email = $1
	AND 
		deleted_at IS NULL
	`
	err := r.db.QueryRow(query, email.Email).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.Website,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) CheckUniques(req *pbu.CheckUniquesRequest) (*pbu.CheckUniquesResponse, error) {
	var exists int
	query := fmt.Sprintf("SELECT count(1) from users WHERE %s = $1 ", req.Field)
	err := r.db.QueryRow(query, req.Value).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if exists == 1 {
		return &pbu.CheckUniquesResponse{
			Check: true,
		}, nil
	}
	return &pbu.CheckUniquesResponse{
		Check: false,
	}, nil
}

func (r *userRepo) GetUserByRefreshToken(refresh *pbu.RefreshToken) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT 
		id, 
		name, 
		last_name,
		username,
		email,
		password,
		bio,
		website,
		refresh_token,
		created_at,
		updated_at,
	FROM 
		users 
	WHERE 
		refresh_token = $1
	AND 
		deleted_at IS NULL
	`
	err := r.db.QueryRow(query, refresh.Token).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.Bio,
		&user.Website,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Create(user *pbu.User) (*pbu.User, error) {
	var respUser pbu.User
	query := `
		INSERT INTO users (
			id, 
			name, 
			last_name,
			username,
			email,
			password,
			bio,
			website,
			refresh_token
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		RETURNING 
			id, 
			name, 
			last_name,
			username,
			email,
			password,
			bio,
			website,
			refresh_token,
			created_at,
			updated_at
	`
	id := uuid.New().String()
	err := r.db.QueryRow(
		query,
		id,
		user.Name,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.Bio,
		user.Website,
		user.RefreshToken,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.LastName,
		&respUser.Username,
		&respUser.Email,
		&respUser.Password,
		&respUser.RefreshToken,
		&respUser.Bio,
		&respUser.Website,
		&respUser.CreatedAt,
		&respUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &respUser, nil
}

func (r *userRepo) Update(user *pbu.User) (*pbu.User, error) {
	var respUser pbu.User
	query := `
		UPDATE 
			users 
		SET 
			updated_at = CURRENT_TIMESTAMP,
			name = $1,
			last_name = $2, 
			username = $3,
			email = $4,
			password = $5,
			bio = $6,
			website = $7,
			refresh_token = $8
		WHERE 
			id = $9
		AND
			deleted_at IS NULL
		RETURNING
			id, 
			name, 
			last_name,
			username,
			email,
			password,
			bio,
			website,
			refresh_token,
			created_at,
			updated_at
	`
	err := r.db.QueryRow(
		query,
		user.Name,
		user.LastName,
		user.Username,
		user.Email,
		user.Password,
		user.Bio,
		user.Website,
		user.RefreshToken,
		user.Id,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.LastName,
		&respUser.Username,
		&respUser.Email,
		&respUser.Password,
		&respUser.RefreshToken,
		&respUser.Bio,
		&respUser.Website,
		&respUser.CreatedAt,
		&respUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &respUser, nil
}

func (r *userRepo) Delete(user_id *pbu.UserRequest) (*pbu.CheckUniquesResponse, error) {
	queryDel := `
	UPDATE
		users
	SET
		deleted_at = CURRENT_TIMESTAMP
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	num, err := r.db.Exec(queryDel, user_id.UserId)
	if err != nil {
		return nil, err
	}

	son, err := num.RowsAffected()
	if err != nil {
		return nil, err
	}

	if son == 0 {
		return &pbu.CheckUniquesResponse{Check: false}, nil
	}

	return &pbu.CheckUniquesResponse{Check: true}, nil
}

func (r *userRepo) Get(user_id *pbu.UserRequest) (*pbu.User, error) {
	var user pbu.User
	query := `
	SELECT 
		id, 
		name, 
		last_name,
		username,
		email,
		password,
		bio,
		website,
		created_at,
		updated_at
	FROM 
		users 
	WHERE 
		id = $1
	AND 
		deleted_at IS NULL
	`
	err := r.db.QueryRow(query, user_id.UserId).Scan(
		&user.Id,
		&user.Name,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Bio,
		&user.Website,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetAll(req *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	offset := req.Limit * (req.Page - 1)
	query := `
		SELECT 
			id, 
			name, 
			last_name,
			username,
			email,
			password,
			bio,
			website,
			created_at,
			updated_at
		FROM 
			users
		WHERE 
			deleted_at IS NULL 	
		LIMIT $1 
		OFFSET $2
	`
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	var allUsers pbu.GetAllUsersResponse
	for rows.Next() {
		var user pbu.User
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.LastName,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Bio,
			&user.Website,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
			return nil, err
		}
		allUsers.AllUsers = append(allUsers.AllUsers, &user)
	}
	return &allUsers, nil
}
