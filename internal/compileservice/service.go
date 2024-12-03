package compileservice

type Config struct {
	SrcPath           string
	OutputPath        string
	OutputContextJSON bool
}

type Service struct {
	conf *Config
}

func New(conf *Config) *Service {
	return &Service{conf}
}
