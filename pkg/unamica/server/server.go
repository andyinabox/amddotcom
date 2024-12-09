package server

type Config struct {
	WebRoot string // path to directory to be served
	Port    int64
}

type Server struct {
	cnf                   *Config
	liveReloadConnections int
}

func New(cnf *Config) *Server {
	return &Server{cnf, 0}
}
