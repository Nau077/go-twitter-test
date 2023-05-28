package app

import (
	"context"
	"go_subs_service/internal/pkg/routes"
	"net/http"

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

	err := a.runPublicHTTP()
	if err != nil {
		return err
	}

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

func (a *App) runPublicHTTP() error {
	port := a.serviceProvider.GetConfig().HTTP.GetPort()

	a.router.Run(port)

	if errHTTP := http.ListenAndServe(":"+port, a.router); errHTTP != nil {
		return errHTTP
	}

	return nil
}
