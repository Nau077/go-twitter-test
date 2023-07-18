package repository

import (
	"context"
	"go_subs_service/internal/pkg/db"
	"go_subs_service/internal/services/subs_service"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type subsRepository struct {
	driver neo4j.DriverWithContext
}

func NewSubsRepository(client *db.Client) subs_service.SubcribeRepository {
	return &subsRepository{
		driver: client.DB(),
	}
}

type Subcription struct {
	id     string
	subsId string
}

func (s *subsRepository) CreateSubscription(ctx context.Context, user1 string, user2 string) (string, error) {
	query := `
		MERGE (u1:User {name: $user1})
		MERGE (u2:User {name: $user2})
		MERGE (u1)-[:IS_FRIEND]->(u2)
	`

	params := map[string]interface{}{
		"user1": user1,
		"user2": user2,
	}

	_, err := neo4j.ExecuteQuery(ctx, s.driver, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return "", err
	} else {
		return "success create response", nil
	}

}

func (s *subsRepository) CancelSubcription(ctx context.Context, user string) (string, error) {
	query := `
		MATCH (u:User {name: $user})
		DETACH DELETE u
	`

	params := map[string]interface{}{
		"user": user,
	}

	_, err := neo4j.ExecuteQuery(ctx, s.driver, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return "", err
	} else {
		return "success delete response with " + user, nil
	}
}

func (s *subsRepository) GetSubcriptionList(ctx context.Context, user string) ([]interface{}, error) {
	query := `
		MATCH (:User {name: $user})-[:IS_FRIEND]->(friend)
		RETURN friend.name AS friendName
	`

	params := map[string]interface{}{
		"user": user,
	}

	result, err := neo4j.ExecuteQuery(ctx, s.driver, query, params, neo4j.EagerResultTransformer)
	if err != nil {
		return nil, err
	}

	var friends []interface{}
	for _, record := range result.Records {
		friends = append(friends, record)
	}

	return friends, nil
}
