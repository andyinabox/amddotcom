package watchservice

type Config struct{}

type Service struct {
	cnf *Config
}

func New(cnf *Config) *Service {
	return &Service{cnf}
}
