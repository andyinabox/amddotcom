package servicecontainer

import (
	"github.com/amddotcom/pkg/unamica"
	"github.com/amddotcom/pkg/unamica/contentparser"
	"github.com/amddotcom/pkg/unamica/renderer"
	"github.com/amddotcom/pkg/unamica/server"
	"github.com/amddotcom/pkg/unamica/watcher"
)

type Container struct {
	renderer unamica.Renderer
	parser   unamica.ContentParser
	watcher  unamica.Watcher
	server   unamica.Server
	cnf      *unamica.Config
}

func New(cnf *unamica.Config) *Container {
	return &Container{
		renderer: renderer.New(&renderer.Config{
			SrcPath:           cnf.SrcPath,
			OutputPath:        cnf.OutputPath,
			AssetsPath:        cnf.AssetsPath,
			OutputContextJSON: cnf.OutputContextJSON,
		}),
		parser: contentparser.New(&contentparser.Config{
			ContentPath:      cnf.ContentPath,
			SiteDataFileName: cnf.SiteDataFileName,
			DateFormat:       cnf.DateFormat,
		}),
		watcher: watcher.New(),
		server: server.New(&server.Config{
			WebRoot: cnf.OutputPath,
			Port:    cnf.Port,
		}),
		cnf: cnf,
	}
}

func (c *Container) Renderer() unamica.Renderer {
	return c.renderer
}

func (c *Container) ContentParser() unamica.ContentParser {
	return c.parser
}

func (c *Container) Watcher() unamica.Watcher {
	return c.watcher
}

func (c *Container) Server() unamica.Server {
	return c.server
}
