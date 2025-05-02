package users

import (
	"context"
	pbusersv1 "users_service/pkg/proto/users/v1"
)

func (i *Implementation) AuthenticateUser(ctx context.Context, req *pbusersv1.UserAuthenticateRequest) *pbusersv1.UserAuthenticateResponse {
	return nil
}
