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
	Data  ContentData     `json:"data"`
	BioXS ContentMarkdown `json:"bio-xs"`
	BioSM ContentMarkdown `json:"bio-sm"`
	BioMD ContentMarkdown `json:"bio-md"`
	BioLG ContentMarkdown `json:"bio-lg"`
	BioXL ContentMarkdown `json:"bio-xl"`
}

type Context struct {
	Content     Content
	ContentJSON template.JS
}

func LoadContent() (*Content, error) {

	files, err := ContentDir.ReadDir("content")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		switch file.Name() {
		case "data.json":
		case "bio-xs.md":
		case "bio-sm.md":
		case "bio-md.md":
		case "bio-lg.md":
		case "bio-xl.md":
		}
	}

	return nil, nil
}
