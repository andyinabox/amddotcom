package contextservice

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/amddotcom/internal"
	"github.com/russross/blackfriday/v2"
)

func (s *Service) loadMdFile(entry fs.DirEntry) (data *internal.MarkdownFile, err error) {
	var b []byte

	b, err = os.ReadFile(filepath.Join(s.conf.ContentPath, entry.Name()))
	if err != nil {
		return
	}

	parsed := blackfriday.Run(b)

	data = &internal.MarkdownFile{
		Raw:    b,
		Parsed: template.HTML(string(parsed)),
	}

	return
}
