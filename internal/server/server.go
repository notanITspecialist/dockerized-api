package server

type (
	Server interface {
		Start()
	}

	appServer struct{}
)

func NewAppServer() (Server, error) {
	return &appServer{}, nil
}

func (s *appServer) Start() {}
