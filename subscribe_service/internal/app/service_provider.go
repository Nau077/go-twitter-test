package app

import (
	"context"
	"go_subs_service/internal/config"
	"go_subs_service/internal/pkg/db"
	"go_subs_service/internal/repository"
	"go_subs_service/internal/services/subs_service"
	"log"
)

type ServiceProvider struct {
	db             *db.Client
	configPath     string
	config         *config.Config
	subsRepository repository.SubcribeRepository
	subsService    *subs_service.Service
}

func newServiceProvider(configPath string) *ServiceProvider {
	return &ServiceProvider{
		configPath: configPath,
	}
}

func (s *ServiceProvider) GetDB(ctx context.Context) *db.Client {
	if s.db == nil {
		cfg := s.GetConfig().GetDBConfig()

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("cant connect to db err %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *ServiceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}
		s.config = cfg
	}

	return s.config
}

func (s *ServiceProvider) GetSubsRepository(ctx context.Context) repository.SubcribeRepository {
	if s.subsRepository == nil {
		s.subsRepository = repository.NewSubsRepository(s.GetDB(ctx))
	}

	return s.subsRepository
}

func (s *ServiceProvider) GetSubsService(ctx context.Context) *subs_service.Service {
	if s.subsService == nil {
		s.subsService = subs_service.NewService(s.GetSubsRepository(ctx))
	}

	return s.subsService
}
