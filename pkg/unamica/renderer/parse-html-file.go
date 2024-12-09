package renderer

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/amddotcom/pkg/unamica"
	"github.com/charmbracelet/log"
)

func (r *Renderer) parseHtmlTemplate(absPath string, rc unamica.RenderContext) (err error) {

	fn := filepath.Base(absPath)

	var b []byte
	b, err = os.ReadFile(absPath)
	if err != nil {
		return
	}

	tmpl, err := template.New(fn).Parse(string(b))
	if err != nil {
		return
	}

	outFile := filepath.Join(r.cnf.OutputPath, strings.Replace(fn, ".tmpl", "", 1))
	log.Debug("parse html file", "file", outFile)

	out, err := os.Create(outFile)
	if err != nil {
		return
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, fn, rc)

	return
}
