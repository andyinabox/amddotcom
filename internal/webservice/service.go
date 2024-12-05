package webservice

type Config struct {
	WebRoot string // path to directory to be served
}

type Service struct {
	cnf *Config
}

func New(cnf *Config) *Service {
	return &Service{cnf}
}
