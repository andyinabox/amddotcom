package buildservice

import (
	"embed"
)

type Config struct {
	SrcPath    string
	OutputPath string
}

type Service struct {
	fs   embed.FS
	conf *Config
}

func New(fs embed.FS, conf *Config) *Service {
	return &Service{fs, conf}
}
