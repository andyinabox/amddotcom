package internal

import (
	"context"
	"html/template"
)

type BuildData struct {
	Time          string
	TimeFormatted string
}

type MarkdownFile struct {
	Source string
	Text   string
	HTML   template.HTML
}

type RenderContext struct {
	BuildData BuildData
	SiteData  map[string]string
	Pages     map[string]MarkdownFile
}

type Builder interface {
	Build(context.Context) error
	Serve(context.Context) error
}

type ContextService interface {
	GetRenderContext(context.Context) (*RenderContext, error)
}

type CompileService interface {
	Compile(context.Context, RenderContext) error
}

type WebService interface {
	Serve(context.Context) error
}

type WatchService interface {
	Watch(context.Context) error
}
