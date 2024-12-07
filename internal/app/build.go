package app

import (
	"context"
	"path/filepath"

	"github.com/charmbracelet/log"
	cp "github.com/otiai10/copy"
)

func (b *App) Build(ctx context.Context) (err error) {
	log.Info("building...")

	rc, err := b.ctx.GetRenderContext(ctx)
	if err != nil {
		return
	}

	err = b.cmp.Compile(ctx, *rc)
	if err != nil {
		return
	}

	// copy assets
	assetsDestDir := filepath.Join(b.cnf.OutputPath, filepath.Base(b.cnf.AssetsPath))
	log.Debug("copying assets", "from", b.cnf.AssetsPath, "to", assetsDestDir)
	err = cp.Copy(b.cnf.AssetsPath, assetsDestDir)

	return
}
