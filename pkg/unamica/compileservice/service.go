package compileservice

type Config struct {
	SrcPath           string
	OutputPath        string
	AssetsPath        string
	OutputContextJSON bool
}

type Service struct {
	cnf *Config
}

func New(cnf *Config) *Service {
	return &Service{cnf}
}
