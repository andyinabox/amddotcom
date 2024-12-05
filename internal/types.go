package internal

import (
	"context"
	"html/template"
)

// structs

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

// interfaces
type App interface {
	Build(context.Context) error
	Serve(ctx context.Context, port int) error
}

type ContextService interface {
	GetRenderContext(context.Context) (*RenderContext, error)
}

type CompileService interface {
	Compile(context.Context, RenderContext) error
}

type WebService interface {
	Serve(ctx context.Context, port int) error
}

type WatchService interface {
	Watch(context.Context) error
}
