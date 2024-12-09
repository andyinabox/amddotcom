package app

import (
	"github.com/amddotcom/unamica"
	"github.com/amddotcom/unamica/compileservice"
	"github.com/amddotcom/unamica/contextservice"
	"github.com/amddotcom/unamica/watchservice"
	"github.com/amddotcom/unamica/webservice"
)

type Config struct {
	SrcPath           string
	OutputPath        string
	ContentPath       string
	AssetsPath        string
	SiteDataFileName  string
	DateFormat        string
	OutputContextJSON bool
	Port              int64
}

type App struct {
	cmp unamica.CompileService
	ctx unamica.ContextService
	srv unamica.WebService
	wch unamica.WatchService
	cnf *Config
}

func New(cnf *Config) *App {
	return &App{
		cmp: compileservice.New(&compileservice.Config{
			SrcPath:           cnf.SrcPath,
			OutputPath:        cnf.OutputPath,
			AssetsPath:        cnf.AssetsPath,
			OutputContextJSON: cnf.OutputContextJSON,
		}),
		ctx: contextservice.New(&contextservice.Config{
			ContentPath:      cnf.ContentPath,
			SiteDataFileName: cnf.SiteDataFileName,
			DateFormat:       cnf.DateFormat,
		}),
		srv: webservice.New(&webservice.Config{
			WebRoot: cnf.OutputPath,
			Port:    cnf.Port,
		}),
		wch: watchservice.New(&watchservice.Config{}),
		cnf: cnf,
	}
}
