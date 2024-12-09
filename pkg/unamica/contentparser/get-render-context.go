package contentparser

import (
	"context"
	"encoding/json"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amddotcom/pkg/unamica"
)

func (p *Parser) GetRenderContext(ctx context.Context) (rc *unamica.RenderContext, err error) {

	files, err := os.ReadDir(p.cnf.ContentPath)
	if err != nil {
		return
	}

	now := time.Now()

	rc = &unamica.RenderContext{
		BuildData: unamica.BuildData{
			Time:          now.Format(time.RFC3339),
			TimeFormatted: now.Format(p.cnf.DateFormat),
		},
		SiteData: make(map[string]string),
		Pages:    make(map[string]unamica.MarkdownFile),
	}

	for _, file := range files {

		// load site metadata
		if file.Name() == p.cnf.SiteDataFileName {
			var b []byte
			b, err = os.ReadFile(filepath.Join(p.cnf.ContentPath, file.Name()))
			if err != nil {
				return
			}
			err = json.Unmarshal(b, &rc.SiteData)

			if err != nil {
				return
			}
		}

		// load markdown pages
		if filepath.Ext(file.Name()) == ".md" {
			var m *unamica.MarkdownFile
			m, err = p.loadMdFile(file)
			if err != nil {
				return
			}
			rc.Pages[getFileName(file)] = *m
		}

	}

	snippet := ctx.Value(unamica.LiveReloadSnippetKey)
	if snippet != nil {
		rc.LiveReloadSnippet = template.HTML(snippet.(string))
	}

	return
}

func getFileName(entry fs.DirEntry) string {
	base := filepath.Base(entry.Name())
	parts := strings.Split(base, ".")
	return parts[0]
}
