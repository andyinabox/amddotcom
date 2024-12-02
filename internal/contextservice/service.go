package contextservice

type Config struct {
	ContentPath      string
	SiteDataFileName string
}

type Service struct {
	conf *Config
}

func New(conf *Config) *Service {
	return &Service{conf}
}
