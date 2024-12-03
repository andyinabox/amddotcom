package internal

import (
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

type Context struct {
	BuildData BuildData
	SiteData  map[string]string
	Pages     map[string]MarkdownFile
}

type Builder interface {
	Build() error
}

type ContextService interface {
	GetContext() (*Context, error)
}

type CompileService interface {
	Compile(*Context) error
}
