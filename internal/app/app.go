package app

import "dockerized-api/internal/dependencies"

type (
	Application interface {
		Run()
	}

	application struct {
		deps dependencies.Dependencies
	}
)

func NewApplication() Application {
	return &application{
		deps: dependencies.NewDependencies(),
	}
}

func (app *application) Run() {
	appServer := app.deps.AppServer()
	appServer.Run()
}
