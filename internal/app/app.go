package app

import (
	"github.com/amddotcom/internal"
	"github.com/amddotcom/internal/compileservice"
	"github.com/amddotcom/internal/contextservice"
	"github.com/amddotcom/internal/watchservice"
	"github.com/amddotcom/internal/webservice"
)

type Config struct {
	SrcPath           string
	OutputPath        string
	ContentPath       string
	AssetsPath        string
	SiteDataFileName  string
	DateFormat        string
	OutputContextJSON bool
}

type App struct {
	cmp internal.CompileService
	ctx internal.ContextService
	srv internal.WebService
	wch internal.WatchService
	cnf *Config
}

func New(cnf *Config) *App {
	return &App{
		cmp: compileservice.New(&compileservice.Config{
			SrcPath:           cnf.SrcPath,
			OutputPath:        cnf.OutputPath,
			OutputContextJSON: cnf.OutputContextJSON,
		}),
		ctx: contextservice.New(&contextservice.Config{
			ContentPath:      cnf.ContentPath,
			SiteDataFileName: cnf.SiteDataFileName,
			DateFormat:       cnf.DateFormat,
		}),
		srv: webservice.New(&webservice.Config{
			WebRoot: cnf.OutputPath,
		}),
		wch: watchservice.New(&watchservice.Config{}),
		cnf: cnf,
	}
}
