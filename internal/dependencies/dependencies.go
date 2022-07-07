package dependencies

import "dockerized-api/internal/server"

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		server server.Server
	}
)

func NewDependencies() Dependencies {
	return &dependencies{}
}

func (d *dependencies) AppServer() server.Server {
	if d.server == nil {
		appServer, err := server.NewServer()
		if err != nil {
			panic("Error while initialize web server!") // temporary
		}

		d.server = appServer
	}
	return d.server
}
