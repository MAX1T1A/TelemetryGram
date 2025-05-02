package users

import (
	"context"
	pbusersv1 "users_service/pkg/proto/users/v1"
)

func (i *Implementation) CreateUser(ctx context.Context, req *pbusersv1.UserRegisterRequest) *pbusersv1.UserRegisterResponse {
	return nil
}
