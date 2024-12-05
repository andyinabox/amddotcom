package contextservice

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/amddotcom/internal"
)

func (s *Service) GetRenderContext(ctx context.Context) (rc *internal.RenderContext, err error) {

	files, err := os.ReadDir(s.conf.ContentPath)
	if err != nil {
		return
	}

	now := time.Now()

	rc = &internal.RenderContext{
		BuildData: internal.BuildData{
			Time:          now.Format(time.RFC3339),
			TimeFormatted: now.Format(s.conf.DateFormat),
		},
		SiteData: make(map[string]string),
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
			err = json.Unmarshal(b, &rc.SiteData)

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
			rc.Pages[getFileName(file)] = *m
		}

	}

	return
}

func getFileName(entry fs.DirEntry) string {
	base := filepath.Base(entry.Name())
	parts := strings.Split(base, ".")
	return parts[0]
}
