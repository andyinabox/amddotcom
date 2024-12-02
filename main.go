package amddotcom

import (
	"embed"
	"html/template"
)

//go:embed content/*
var ContentFS embed.FS

//go:embed src/*
var SrcFS embed.FS

type SiteData struct {
	Title string `json:"title"`
}

type MarkdownFile struct {
	Raw    []byte
	Parsed template.HTML
}

type Context struct {
	SiteData SiteData
	Pages    map[string]MarkdownFile
}

type ContextService interface {
	GetContext() (*Context, error)
}

type BuildService interface {
	Build(*Context) error
}
