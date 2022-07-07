package dependencies

import (
	"dockerized-api/internal/endpoints"
	"dockerized-api/internal/server"
	"dockerized-api/internal/service"
)

type (
	Dependencies interface {
		AppServer() server.Server
	}

	dependencies struct {
		server server.Server

		orderEndpoints endpoints.OrderEndpoints
		orderService   service.OrderService
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
		orderEndpoints := endpoints.NewOrderEndpoints(
			d.OrderService(),
		)
		d.orderEndpoints = orderEndpoints
	}
	return d.orderEndpoints
}
func (d *dependencies) OrderService() service.OrderService {
	if d.orderService == nil {
		orderService := service.NewOrderService()
		d.orderService = orderService
	}
	return d.orderService
}
