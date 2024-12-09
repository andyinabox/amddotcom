package renderer

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/amddotcom/pkg/unamica"
	"github.com/charmbracelet/log"
)

func (r *Renderer) parseHtmlTemplate(f fs.DirEntry, rc unamica.RenderContext) (err error) {
	var b []byte
	b, err = os.ReadFile(filepath.Join(r.cnf.SrcPath, f.Name()))
	if err != nil {
		return
	}

	tmpl, err := template.New(f.Name()).Parse(string(b))
	if err != nil {
		return
	}

	outFile := filepath.Join(r.cnf.OutputPath, strings.Replace(f.Name(), ".tmpl", "", 1))
	log.Debug("parse html file", "file", outFile)

	out, err := os.Create(outFile)
	if err != nil {
		return
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, f.Name(), rc)

	return
}
