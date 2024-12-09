package contextservice

type Config struct {
	ContentPath      string
	SiteDataFileName string
	DateFormat       string
}

type Service struct {
	conf *Config
}

func New(conf *Config) *Service {
	return &Service{conf}
}
