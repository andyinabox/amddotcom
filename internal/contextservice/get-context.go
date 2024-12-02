package contextservice

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/amddotcom"
	"github.com/iancoleman/strcase"
)

func (s *Service) GetContext() (ctx *amddotcom.Context, err error) {

	files, err := s.fs.ReadDir(s.conf.ContentPath)
	if err != nil {
		return
	}

	ctx = &amddotcom.Context{
		SiteData: amddotcom.SiteData{},
		Pages:    make(map[string]amddotcom.MarkdownFile),
	}

	for _, file := range files {

		// load site metadata
		if file.Name() == s.conf.SiteDataFileName {
			var b []byte
			b, err = s.fs.ReadFile(filepath.Join(s.conf.ContentPath, file.Name()))
			if err != nil {
				return
			}
			err = json.Unmarshal(b, &ctx.SiteData)

			if err != nil {
				return
			}
		}

		// load markdown pages
		if filepath.Ext(file.Name()) == ".md" {
			var m *amddotcom.MarkdownFile
			m, err = s.loadMdFile(file)
			if err != nil {
				return
			}
			ctx.Pages[getFileName(file)] = *m
		}

	}

	var b []byte
	b, err = json.Marshal(ctx)
	if err != nil {
		return
	}

	ctx.JSONData = template.JS(string(b))

	return
}

func getFileName(entry fs.DirEntry) string {
	base := filepath.Base(entry.Name())
	parts := strings.Split(base, ".")
	name := parts[0]
	return strcase.ToCamel(name)
}
