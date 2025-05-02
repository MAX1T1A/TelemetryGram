package users

import (
	usersservicev1 "users_service/internal/services/users/v1"
	pbusersv1 "users_service/pkg/proto/users/v1"
)

type Implementation struct {
	pbusersv1.UnimplementedUsersServiceV1Server
	usersServiceV1 usersservicev1.Service
}

func NewImplementation(usersServiceV1 usersservicev1.Service) *Implementation {
	return &Implementation{
		usersServiceV1: usersServiceV1,
	}
}
