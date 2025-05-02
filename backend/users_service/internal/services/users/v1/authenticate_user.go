package users

import (
	"context"

	usersmodelsv1 "users_service/internal/models/users/v1"
)

func (s *service) AuthenticateUser(ctx context.Context, user usersmodelsv1.UserAuthenticateRequest) (*usersmodelsv1.UserAuthenticateResponse, error) {

	userAuthenticateResponse, err := s.usersRepositoryV1.AuthenticateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return userAuthenticateResponse, nil

}
