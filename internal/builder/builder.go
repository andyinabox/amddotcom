package builder

import (
	"github.com/amddotcom/internal"
	"github.com/amddotcom/internal/compileservice"
	"github.com/amddotcom/internal/contextservice"
)

type Config struct {
	SrcPath          string
	OutputPath       string
	ContentPath      string
	SiteDataFileName string
	DateFormat       string
}

type Builder struct {
	cmp internal.CompileService
	ctx internal.ContextService
	cnf *Config
}

func New(cnf *Config) *Builder {
	return &Builder{
		cmp: compileservice.New(&compileservice.Config{
			SrcPath:    cnf.SrcPath,
			OutputPath: cnf.OutputPath,
		}),
		ctx: contextservice.New(&contextservice.Config{
			ContentPath:      cnf.ContentPath,
			SiteDataFileName: cnf.SiteDataFileName,
			DateFormat:       cnf.DateFormat,
		}),
		cnf: cnf,
	}
}
