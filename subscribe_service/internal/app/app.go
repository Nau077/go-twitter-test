package app

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine

	serviceProvider *serviceProvider
	pathConfig      string
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
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

func (a *App) initServer(ctx context.Context) error {
	a.subscription = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) {
	gin.SetMode(gin.ReleaseMode)
	a.routes.NewSubsHTTPHandler(r.router, r.subsService)
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	go func() {
		defer wg.Done()

		a.router.Run(port)
		port := a.serviceProvider.GetConfig().HTTP.GetPort()
		if errHTTP := http.ListenAndServe(":"+port, a.router); errHTTP != nil {
			return errHTTP
		}
	}()

	log.Printf("run http server on host %s", a.serviceProvider.GetConfig().HTTP.GetAddress())

	return nil
}
