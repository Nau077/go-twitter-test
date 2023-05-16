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
	db             db.Client
	configPath     string
	config         *config.Config
	subsRepository *repository.SubcribeRepository
	subsService    *subs_service.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

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

func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.Repository {
	if s.noteRepository == nil {
		s.noteRepository = repository.(s.GetDB(ctx))
	}

	return s.noteRepository
}

// GetNoteService
func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
