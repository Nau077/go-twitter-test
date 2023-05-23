package repository

import (
	"context"
	"fmt"
	"go_subs_service/internal/pkg/db"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type SubcribeRepository interface {
	CreateSubscription(ctx context.Context, id string, subsId string) (*Subcription, error)
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
	id     string
	subsId string
}

func (s *subsRepository) CreateSubscription(ctx context.Context, id string, subsId string) (*Subcription, error) {
	result, err := neo4j.ExecuteQuery(ctx, s.driver,
		"CREATE (igor: Person {id: $id}), ($id) -[:IS_SUBSCRIBED]-> ($subsId)",
		map[string]any{
			"id":     id,
			"subsId": subsId,
		}, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}
	itemNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
	if err != nil {
		return nil, fmt.Errorf("could not find node n")
	}
	cId, err := neo4j.GetProperty[string](itemNode, id)
	if err != nil {
		return nil, err
	}

	return &Subcription{id: cId}, nil
}

func (s *subsRepository) CancelSubcription(ctx context.Context) error {
	return nil
}

func (s *subsRepository) GetSubcriptionList(ctx context.Context) error {
	return nil
}
