package compileservice

import (
	"path/filepath"

	"github.com/charmbracelet/log"
	cp "github.com/otiai10/copy"
)

func (s *Service) copyAssets() error {
	assetsDestDir := filepath.Join(s.cnf.OutputPath, filepath.Base(s.cnf.AssetsPath))
	log.Debug("copying assets", "from", s.cnf.AssetsPath, "to", assetsDestDir)
	return cp.Copy(s.cnf.AssetsPath, assetsDestDir)
}
