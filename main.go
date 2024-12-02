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
	Raw    []byte        `json:"-"`
	Parsed template.HTML `json:"parsed"`
}

type Context struct {
	SiteData SiteData                `json:"siteData"`
	Pages    map[string]MarkdownFile `json:"pages"`
	JSONData template.JS             `json:"-"`
}

type ContextService interface {
	GetContext() (*Context, error)
}

type BuildService interface {
	Build(*Context) error
}
