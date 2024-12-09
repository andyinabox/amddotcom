package contentparser

type Config struct {
	ContentPath      string
	SiteDataFileName string
	DateFormat       string
}

type Parser struct {
	cnf *Config
}

func New(cnf *Config) *Parser {
	return &Parser{cnf}
}
