package buildservice

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/amddotcom/internal"
)

func (s *Service) Build(ctx *internal.Context) (err error) {

	files, err := os.ReadDir(s.conf.SrcPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.Name() == "index.html.tmpl" {
			copyErr := s.parseHtmlTemplate(file, ctx)
			if copyErr != nil {
				fmt.Printf("error: %s\n", copyErr)
			}
		} else {
			copyErr := s.copyFile(file)
			if copyErr != nil {
				fmt.Printf("error: %s\n", copyErr)
			}
		}
	}

	return
}

func (s *Service) copyFile(f fs.DirEntry) (err error) {

	var b []byte
	b, err = os.ReadFile(filepath.Join(s.conf.SrcPath, f.Name()))
	if err != nil {
		return
	}

	outFile := filepath.Join(s.conf.OutputPath, f.Name())
	fmt.Printf("copy file %s\n", outFile)
	err = os.WriteFile(outFile, b, os.ModePerm)
	return
}

func (s *Service) parseHtmlTemplate(f fs.DirEntry, ctx *internal.Context) (err error) {
	var b []byte
	b, err = os.ReadFile(filepath.Join(s.conf.SrcPath, f.Name()))
	if err != nil {
		return
	}

	tmpl, err := template.New(f.Name()).Parse(string(b))
	if err != nil {
		return
	}

	outFile := filepath.Join(s.conf.OutputPath, strings.Replace(f.Name(), ".tmpl", "", 1))
	fmt.Printf("parse html file %s\n", outFile)

	out, err := os.Create(outFile)
	if err != nil {
		return
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, f.Name(), ctx)

	return
}
