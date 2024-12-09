package app

import (
	"context"

	"github.com/charmbracelet/log"
)

func (a *App) Build(ctx context.Context) (err error) {
	log.Info("building...")

	rc, err := a.ctx.GetRenderContext(ctx)
	if err != nil {
		return
	}

	return a.cmp.Compile(ctx, *rc)
}
