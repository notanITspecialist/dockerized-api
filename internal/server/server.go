package server

import (
	"dockerized-api/internal/endpoints"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

type (
	Server interface {
		Run()
	}

	server struct {
		listener net.Listener
		server   *http.Server

		orderEndpoints endpoints.OrderEndpoints
	}
)

func NewServer(
	orderEndpoints endpoints.OrderEndpoints,
) (Server, error) {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		return nil, err
	}

	router := mux.NewRouter()
	server := &server{
		listener: listener,
		server:   &http.Server{Handler: router},

		orderEndpoints: orderEndpoints,
	}
	server.initRoutes(router)

	return server, nil
}

func (s *server) Run() {
	err := s.server.Serve(s.listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) initRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/orders", s.orderEndpoints.GetOrders).Methods(http.MethodGet)
}
