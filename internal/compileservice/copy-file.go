package compileservice

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func (s *Service) copyFile(f fs.DirEntry) (err error) {

	var b []byte
	b, err = os.ReadFile(filepath.Join(s.cnf.SrcPath, f.Name()))
	if err != nil {
		return
	}

	outFile := filepath.Join(s.cnf.OutputPath, f.Name())
	log.Debug("copy file", "file", outFile)
	err = os.WriteFile(outFile, b, os.ModePerm)
	return
}
