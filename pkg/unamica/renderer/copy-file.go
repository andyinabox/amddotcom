package renderer

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func (r *Renderer) copyFile(absPath string) (err error) {

	var b []byte
	b, err = os.ReadFile(absPath)
	if err != nil {
		return
	}

	outFile := filepath.Join(r.cnf.OutputPath, filepath.Base(absPath))
	log.Debug("copy file", "from", absPath, "to", outFile)
	err = os.WriteFile(outFile, b, os.ModePerm)
	return
}
