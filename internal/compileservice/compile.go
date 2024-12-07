package compileservice

import (
	"context"
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/amddotcom/internal"
	"github.com/charmbracelet/log"
)

func (s *Service) Compile(ctx context.Context, rc internal.RenderContext) (err error) {

	files, err := os.ReadDir(s.conf.SrcPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.Name() == "index.html.tmpl" {
			copyErr := s.parseHtmlTemplate(file, rc)
			if copyErr != nil {
				log.Error(copyErr)
			}
		} else {
			copyErr := s.copyFile(file)
			if copyErr != nil {
				log.Error(copyErr)
			}
		}
	}

	if s.conf.OutputContextJSON {
		var data []byte
		data, err = json.Marshal(rc)
		if err != nil {
			return
		}

		contextFilePath := filepath.Join(s.conf.OutputPath, "context.json")
		log.Debug("output context file", "file", contextFilePath)
		err = os.WriteFile(contextFilePath, data, os.ModePerm)
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
	log.Debug("copy file", "file", outFile)
	err = os.WriteFile(outFile, b, os.ModePerm)
	return
}

func (s *Service) parseHtmlTemplate(f fs.DirEntry, rc internal.RenderContext) (err error) {
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
	log.Debug("parse html file", "file", outFile)

	out, err := os.Create(outFile)
	if err != nil {
		return
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, f.Name(), rc)

	return
}
