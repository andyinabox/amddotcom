package contentparser

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/amddotcom/pkg/unamica"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func (p *Parser) loadMdFile(entry fs.DirEntry) (data *unamica.MarkdownFile, err error) {
	var md []byte

	// read raw markdown
	md, err = os.ReadFile(filepath.Join(p.cnf.ContentPath, entry.Name()))
	if err != nil {
		return
	}

	// convert md to html
	parsed := string(blackfriday.Run(md))

	// strip html
	sanitizer := bluemonday.StrictPolicy()
	clean := sanitizer.Sanitize(parsed)

	data = &unamica.MarkdownFile{
		Source: string(md),
		Text:   clean,
		HTML:   template.HTML(parsed),
	}

	return
}
