package renderer

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func (r *Renderer) copyFile(f fs.DirEntry) (err error) {

	var b []byte
	b, err = os.ReadFile(filepath.Join(r.cnf.SrcPath, f.Name()))
	if err != nil {
		return
	}

	outFile := filepath.Join(r.cnf.OutputPath, f.Name())
	log.Debug("copy file", "file", outFile)
	err = os.WriteFile(outFile, b, os.ModePerm)
	return
}
