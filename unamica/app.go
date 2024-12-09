package unamica

import (
	"context"
	"html/template"
)

// misc
type ContextKey string

const LiveReloadSnippetKey ContextKey = "LiveReloadSnippet"

// constructor

func New(sc ServiceContainer, cnf *Config) *App {
	return &App{sc, cnf}
}

// structs
type App struct {
	sc  ServiceContainer
	cnf *Config
}

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
// type App interface {
// 	Build(context.Context) error
// 	Serve(ctx context.Context, port int) error
// }

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

type ServiceContainer interface {
	ContextService() ContextService
	CompileService() CompileService
	WebService() WebService
	WatchService() WatchService
}
