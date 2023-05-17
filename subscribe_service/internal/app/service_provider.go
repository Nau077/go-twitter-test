package app

import (
	"context"
	"go_subs_service/internal/config"
	"go_subs_service/internal/pkg/db"
	"go_subs_service/internal/repository"
	"go_subs_service/internal/services/subs_service"
	"log"
)

type serviceProvider struct {
	db             *db.Client
	configPath     string
	config         *config.Config
	subsRepository repository.SubcribeRepository
	subsService    *subs_service.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) *db.Client {
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

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Fatalf("failed to get config: %s", err.Error())
		}
		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetSubsRepository(ctx context.Context) repository.SubcribeRepository {
	if s.subsRepository == nil {
		s.subsRepository = repository.NewSubsRepository(s.GetDB(ctx))
	}

	return s.subsRepository
}

func (s *serviceProvider) GetSubsService(ctx context.Context) *Subs.Service {
	if s.SubsService == nil {
		s.SubsService = Subs.NewService(s.GetSubsRepository(ctx))
	}

	return s.SubsService
}
