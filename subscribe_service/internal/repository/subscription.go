package repository

import (
	"context"
	"go_subs_service/internal/pkg/db"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SubcribeRepository interface {
	CreateSubcription(ctx context.Context) (int64, error)
	CancelSubcription(ctx context.Context) error
	GetSubcriptionList(ctx context.Context) error
}

type subsRepository struct {
	driver neo4j.DriverWithContext
}

func NewSubsRepository(client *db.Client) SubcribeRepository {
	return &subsRepository{
		driver: client.DB(),
	}
}

type Subcription struct {
	id         int64
	ownId      string
	userSubsId string
}

func (s *subsRepository) CreateSubcription(ctx context.Context) (int64, error) {

	return 2, nil
}

func (s *subsRepository) CancelSubcription(ctx context.Context) error {
	return nil
}

func (s *subsRepository) GetSubcriptionList(ctx context.Context) error {
	return nil
}
