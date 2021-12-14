package main

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) CreateUser(ctx context.Context, user User) (User, error) {
	err := repo.db.QueryRowxContext(
		ctx,
		"INSERT INTO users (username, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Username, user.FirstName, user.LastName, user.Email, user.Phone,
	).Scan(&user.ID)
	return user, err
}

func (repo *userRepository) GetUser(ctx context.Context, id int64) (User, error) {
	var user User
	err := repo.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}

func (repo *userRepository) UpdateUser(ctx context.Context, id int64, user User) (User, error) {
	_, err := repo.db.ExecContext(
		ctx,
		"UPDATE users SET username=$1, first_name=$2, last_name=$3, email=$4, phone=$5 WHERE id = $6",
		user.Username, user.FirstName, user.LastName, user.Email, user.Phone, id,
	)
	if err != nil {
		return user, err
	}
	return repo.GetUser(ctx, id)
}

func (repo *userRepository) DeleteUser(ctx context.Context, id int64) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
