package app

import (
	"context"

	"github.com/charmbracelet/log"
)

func (b *App) Build(ctx context.Context) (err error) {
	log.Info("building...")

	rc, err := b.ctx.GetRenderContext(ctx)
	if err != nil {
		return
	}

	return b.cmp.Compile(ctx, *rc)
}
