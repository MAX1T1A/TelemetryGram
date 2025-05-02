package app

import (
	"context"
	"log"
	"users_service/internal/config"
	"users_service/internal/db"
	usersRepositoryV1 "users_service/internal/repositories/users/v1"
	usersServiceV1 "users_service/internal/services/users/v1"
)

type serviceProvider struct {
	postgresDB db.PostgresClient

	grpcConfig config.GRPCConfig

	usersRepositoryV1 usersRepositoryV1.Repository
	usersServiceV1    usersServiceV1.Service
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PostgresDB(ctx context.Context) db.PostgresClient {
	if s.postgresDB == nil {
		config, err := config.GetPostgresConfig()
		if err != nil {
			log.Fatalf("не удалось получить конфигурацию Postgres: %s", err.Error())
		}

		db, err := db.NewPostgresClient(ctx, config)
		if err != nil {
			log.Fatalf("не удалось подключиться к Postgres: %s", err.Error())
		}

		s.postgresDB = db
	}

	return s.postgresDB
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("не удалось получить конфигурацию grpc: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) UsersRepositoryV1() usersRepositoryV1.Repository {
	if s.usersRepositoryV1 == nil {
		s.usersRepositoryV1 = usersRepositoryV1.NewRepository(s.PostgresDB(context.Background()))
	}
	return s.usersRepositoryV1
}

func (s *serviceProvider) UsersServiceV1() usersServiceV1.Service {
	if s.usersServiceV1 == nil {
		s.usersServiceV1 = usersServiceV1.NewService(s.UsersRepositoryV1())
	}
	return s.usersServiceV1
}

// func (s *serviceProvider) TracksImpl() *tracks.Implementation {
// 	if s.tracksImpl == nil {
// 		s.tracksImpl = tracks.NewImplementation(s.TracksServiceV2())
// 	}

// 	return s.tracksImpl
// }
