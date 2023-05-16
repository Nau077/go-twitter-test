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

func NewSubsRepository(client db.Client) *subsRepository {
	return &subsRepository{
		driver: client.DB(),
	}
}

type Subcription struct {
	id         int64
	ownId      string
	userSubsId string
}

func CreateSubcription(ctx context.Context) (int64, error) {

	return 2, nil
}

func CancelSubcription(ctx context.Context) (int64, error) {
	return 2, nil
}

func GetSubcriptionList(ctx context.Context) error {
	return nil
}
