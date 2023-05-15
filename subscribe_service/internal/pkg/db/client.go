package db

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type DriverType interface {
	Close()
}

type Client interface {
	Close() error
	DB() DriverType
}

type client struct {
	db        DriverType
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context) (Client, error) {
	dbUri := "neo4j://localhost"
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth("neo4j", "letmein!", ""))
	if err != nil {
		panic(err)
	}

	_, cancel := context.WithCancel(ctx)

	return &client{
		db:        driver,
		closeFunc: cancel,
	}, nil
}

// close() отменяет контекст
// у пуллера вызывает close()
func (c *client) Close() error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}

		if c.db != nil {
			c.db.Close()
		}
	}

	return nil
}

func (c *client) DB() DriverType {
	return c.db
}
