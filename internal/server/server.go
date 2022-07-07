package server

type (
	Server interface {
		Run()
	}

	server struct{}
)

func NewServer() (Server, error) {
	return &server{}, nil
}

func (s *server) Run() {
	panic("SSDsDfs")
}
