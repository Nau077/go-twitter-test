package app

import (
	"context"
	"go_subs_service/internal/pkg/routes"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type App struct {
	router          *gin.Engine
	serviceProvider *ServiceProvider
	pathConfig      string
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		a.serviceProvider.db.Close(ctx)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	for err := range a.runPublicHTTP(wg) {
		if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.router = routes.NewSubsHTTPHandler(ctx, gin.Default(), a.serviceProvider.GetSubsService(ctx))

	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) <-chan error {
	resCh := make(chan error)

	go func() {
		defer func() {
			wg.Done()
			close(resCh)
		}()

		port := a.serviceProvider.GetConfig().HTTP.GetPort()
		a.router.Run(port)
		if errHTTP := http.ListenAndServe(":"+port, a.router); errHTTP != nil {
			resCh <- errHTTP
			return
		}
	}()

	log.Printf("run http server on host %s", a.serviceProvider.GetConfig().HTTP.GetAddress())

	return resCh
}
