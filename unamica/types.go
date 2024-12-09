package unamica

import (
	"context"
	"html/template"
)

// structs

type ContextKey string

const LiveReloadSnippetKey ContextKey = "LiveReloadSnippet"

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
	BuildData         BuildData
	SiteData          map[string]string
	Pages             map[string]MarkdownFile
	LiveReloadSnippet template.HTML
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
	Serve(ctx context.Context) (refresh chan<- string, err error)
	GetLiveReloadSnippet() string
}

type WatchService interface {
	Watch(str context.Context, paths []string) (changes <-chan string, err error)
}
