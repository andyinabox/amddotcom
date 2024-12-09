package renderer

type Config struct {
	SrcPath           string
	OutputPath        string
	AssetsPath        string
	OutputContextJSON bool
}

type Renderer struct {
	cnf *Config
}

func New(cnf *Config) *Renderer {
	return &Renderer{cnf}
}
