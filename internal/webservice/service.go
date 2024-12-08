package webservice

const liveReloadPath = "/_livereload"

type Config struct {
	WebRoot string // path to directory to be served
	Port    int64
}

type Service struct {
	cnf *Config
}

func New(cnf *Config) *Service {
	return &Service{cnf}
}
