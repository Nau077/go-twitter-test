package main

import (
	"context"
	"log"

	"go_subs_service/internal"
	// "os"
	// "regexp"
	// "github.com/gin-contrib/logger"
	// "github.com/gin-gonic/gin"
	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

func main() {
	staticPath := "./static" //os.Args[1]
	ctx := context.Background()
	a, err := internal.NewApp(ctx, staticPath)

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app %s", err.Error())
	}
}
