package dependencies

import "dockerized-api/internal/server"

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		appServer server.Server
	}
)

func NewDependencies() Dependencies {
	return &dependencies{}
}

func (d *dependencies) AppServer() server.Server {
	if d.appServer == nil {
		appServer, err := server.NewAppServer()
		if err != nil {
			panic("SERVER INITIALIZE ERROR!") // temporary
		}

		d.appServer = appServer
	}
	return d.appServer
}
