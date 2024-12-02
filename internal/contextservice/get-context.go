package contextservice

import (
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/amddotcom/internal"
	"github.com/iancoleman/strcase"
)

func (s *Service) GetContext() (ctx *internal.Context, err error) {

	files, err := os.ReadDir(s.conf.ContentPath)
	if err != nil {
		return
	}

	ctx = &internal.Context{
		SiteData: internal.SiteData{},
		Pages:    make(map[string]internal.MarkdownFile),
	}

	for _, file := range files {

		// load site metadata
		if file.Name() == s.conf.SiteDataFileName {
			var b []byte
			b, err = os.ReadFile(filepath.Join(s.conf.ContentPath, file.Name()))
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
			var m *internal.MarkdownFile
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
