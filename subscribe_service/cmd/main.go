package main

import (
	"context"
	"flag"
	"log"

	"go_subs_service/internal/app"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./config/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	a, err := app.NewApp(ctx, pathConfig)

	err = a.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
