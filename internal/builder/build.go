package builder

import (
	"context"
	"fmt"
	"path/filepath"

	cp "github.com/otiai10/copy"
)

func (b *Builder) Build(ctx context.Context) (err error) {
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
	fmt.Printf("copying assets from %s to %s\n", b.cnf.AssetsPath, assetsDestDir)
	err = cp.Copy(b.cnf.AssetsPath, assetsDestDir)

	return
}
