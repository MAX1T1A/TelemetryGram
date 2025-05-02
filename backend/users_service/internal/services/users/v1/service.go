package users

import (
	"context"
	usersmodelsv1 "users_service/internal/models/users/v1"
	usersrepositoryv1 "users_service/internal/repositories/users/v1"
)

type Service interface {
	CreateUser(ctx context.Context, user usersmodelsv1.UserRegisterRequest) error
	AuthenticateUser(ctx context.Context, user usersmodelsv1.UserAuthenticateRequest) (*usersmodelsv1.UserAuthenticateResponse, error)
}

var _ Service = (*service)(nil)

type service struct {
	usersRepositoryV1 usersrepositoryv1.Repository
}

func NewService(usersRepositoryV1 usersrepositoryv1.Repository) Service {
	return &service{
		usersRepositoryV1: usersRepositoryV1,
	}
}
