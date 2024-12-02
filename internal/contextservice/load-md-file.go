package contextservice

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/amddotcom"
	"github.com/russross/blackfriday/v2"
)

func (s *Service) loadMdFile(entry fs.DirEntry) (data *amddotcom.MarkdownFile, err error) {
	var b []byte

	b, err = s.fs.ReadFile(filepath.Join(s.conf.ContentPath, entry.Name()))
	if err != nil {
		return
	}

	parsed := blackfriday.Run(b)

	data = &amddotcom.MarkdownFile{
		Raw:    b,
		Parsed: template.HTML(string(parsed)),
	}

	return
}
