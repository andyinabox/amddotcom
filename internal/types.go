package internal

import (
	"html/template"
	"time"
)

type BuildData struct {
	Time          time.Time
	TimeFormatted string
}

type MarkdownFile struct {
	Source string
	Text   string
	HTML   template.HTML
	Meta   map[string]string
}

type Context struct {
	BuildData BuildData
	SiteData  map[string]string
	Pages     map[string]MarkdownFile
	JSONData  template.JS
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
