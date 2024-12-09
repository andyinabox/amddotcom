package servicecontainer

import (
	"github.com/amddotcom/pkg/unamica"
	"github.com/amddotcom/pkg/unamica/compileservice"
	"github.com/amddotcom/pkg/unamica/contextservice"
	"github.com/amddotcom/pkg/unamica/watchservice"
	"github.com/amddotcom/pkg/unamica/webservice"
)

type Container struct {
	compileService unamica.CompileService
	contextService unamica.ContextService
	watchService   unamica.WatchService
	webService     unamica.WebService
	cnf            *unamica.Config
}

func New(cnf *unamica.Config) *Container {
	return &Container{
		compileService: compileservice.New(&compileservice.Config{
			SrcPath:           cnf.SrcPath,
			OutputPath:        cnf.OutputPath,
			AssetsPath:        cnf.AssetsPath,
			OutputContextJSON: cnf.OutputContextJSON,
		}),
		contextService: contextservice.New(&contextservice.Config{
			ContentPath:      cnf.ContentPath,
			SiteDataFileName: cnf.SiteDataFileName,
			DateFormat:       cnf.DateFormat,
		}),
		watchService: watchservice.New(&watchservice.Config{}),
		webService: webservice.New(&webservice.Config{
			WebRoot: cnf.OutputPath,
			Port:    cnf.Port,
		}),
		cnf: cnf,
	}
}

func (c *Container) CompileService() unamica.CompileService {
	return c.compileService
}

func (c *Container) ContextService() unamica.ContextService {
	return c.contextService
}

func (c *Container) WatchService() unamica.WatchService {
	return c.watchService
}

func (c *Container) WebService() unamica.WebService {
	return c.webService
}
