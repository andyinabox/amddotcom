package renderer

import (
	"context"
	"path/filepath"

	"github.com/charmbracelet/log"
	cp "github.com/otiai10/copy"
)

func (r *Renderer) CopyAssets(ctx context.Context) error {
	assetsDestDir := filepath.Join(r.cnf.OutputPath, filepath.Base(r.cnf.AssetsPath))
	log.Debug("copying assets", "from", r.cnf.AssetsPath, "to", assetsDestDir)
	return cp.Copy(r.cnf.AssetsPath, assetsDestDir)
}
