package db

import (
	"context"
	"go_subs_service/internal/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Client struct {
	db        neo4j.DriverWithContext
	closeFunc context.CancelFunc
}

func NewClient(ctx context.Context, cfg *config.DB) (*Client, error) {
	driver, err := neo4j.NewDriverWithContext(cfg.DSN, neo4j.BasicAuth(cfg.USER, cfg.PASS, ""))
	if err != nil {
		panic(err)
	}

	_, cancel := context.WithCancel(ctx)

	return &Client{
		db:        driver,
		closeFunc: cancel,
	}, nil
}

func (c *Client) Close(ctx context.Context) error {
	if c != nil {
		if c.closeFunc != nil {
			c.closeFunc()
		}

		if c.db != nil {
			c.db.Close(ctx)
		}
	}

	return nil
}

func (c *Client) DB() neo4j.DriverWithContext {
	return c.db
}
