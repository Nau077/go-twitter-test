package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type SubcribeRepository interface {
	CreateSubcription(ctx context.Context) (int64, error)
	CancelSubcription(ctx context.Context) error
	GetSubcriptionList(ctx context.Context) error
}

type repository struct {
	db *sqlx.DB
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
