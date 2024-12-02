package contextservice

import (
	"embed"
)

type Config struct {
	ContentPath      string
	SiteDataFileName string
}

type Service struct {
	fs   embed.FS
	conf *Config
}

func New(fs embed.FS, conf *Config) *Service {
	return &Service{fs, conf}
}
