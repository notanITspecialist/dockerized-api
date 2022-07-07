package dependencies

import (
	"dockerized-api/internal/endpoints"
	"dockerized-api/internal/server"
)

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		server server.Server

		orderEndpoints endpoints.OrderEndpoints
	}
)

func NewDependencies() Dependencies {
	return &dependencies{}
}

func (d *dependencies) AppServer() server.Server {
	if d.server == nil {
		appServer, err := server.NewServer(d.OrderEndpoints())
		if err != nil {
			panic(err)
		}

		d.server = appServer
	}
	return d.server
}

func (d *dependencies) OrderEndpoints() endpoints.OrderEndpoints {
	if d.orderEndpoints == nil {
		orderEndpoints := endpoints.NewOrderEndpoints()
		d.orderEndpoints = orderEndpoints
	}
	return d.orderEndpoints
}
