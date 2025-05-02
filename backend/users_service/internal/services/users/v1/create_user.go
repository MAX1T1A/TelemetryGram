package users

import (
	"context"
	usersmodelsv1 "users_service/internal/models/users/v1"
)

func (s *service) CreateUser(ctx context.Context, user usersmodelsv1.UserRegisterRequest) error {

	err := s.usersRepositoryV1.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil

}
