package amdvanilla

import (
	"embed"
	"errors"
	"fmt"
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
		fmt.Println(file.Name())
	}

	return nil, errors.New("not implemented")
}

func main() {
	_, err := LoadContent()

	if err != nil {
		panic(err)
	}
}
