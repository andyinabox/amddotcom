package amdvanilla

import (
	"embed"
	"html/template"
)

//go:embed content/*
var ContentDir embed.FS

//go:embed src/*
var SrcDir embed.FS

type ContentData struct {
	Name string `json:"name"`
}

type ContentMarkdown struct {
	Raw    string        `json:"raw"`
	Parsed template.HTML `json:"parsed"`
}

type Content struct {
	Data        ContentData     `json:"data"`
	BioXS       ContentMarkdown `json:"bio-xs"`
	BioSM       ContentMarkdown `json:"bio-sm"`
	BioMD       ContentMarkdown `json:"bio-md"`
	BioLG       ContentMarkdown `json:"bio-lg"`
	BioXL       ContentMarkdown `json:"bio-xl"`
	ContentJSON template.JS     `json:"-"`
}
