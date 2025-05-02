package users

import (
	"context"
	"users_service/internal/db"
	usersmodelsv1 "users_service/internal/models/users/v1"
)

type Repository interface {
	CreateUser(ctx context.Context, user usersmodelsv1.UserRegisterRequest) error
	AuthenticateUser(ctx context.Context, user usersmodelsv1.UserAuthenticateRequest) (*usersmodelsv1.UserAuthenticateResponse, error)
}

var _ Repository = (*repository)(nil)

type repository struct {
	db db.PostgresClient
}

func NewRepository(db db.PostgresClient) *repository {
	return &repository{
		db: db,
	}
}
