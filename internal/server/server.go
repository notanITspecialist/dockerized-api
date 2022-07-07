package server

import "fmt"

type (
	Server interface {
		Start()
	}

	server struct{}
)

func NewServer() (Server, error) {
	return &server{}, nil
}

func (s *server) Start() {
	fmt.Println("Start SERVER!")
}
