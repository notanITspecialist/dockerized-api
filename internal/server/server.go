package server

import (
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
	}
)

func NewServer() (Server, error) {
	port := ":8000"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	server := &server{
		listener: listener,
		server:   &http.Server{Handler: mux.NewRouter()},
	}

	return server, nil
}

func (s *server) Run() {
	err := s.server.Serve(s.listener)
	if err != nil {
		panic(err)
	}
}

func (s *server) initRoutes(r *mux.Router) {
}
