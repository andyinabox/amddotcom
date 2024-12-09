package unamica

import (
	"context"

	"github.com/charmbracelet/log"
)

func (a *App) Build(ctx context.Context) (err error) {
	log.Info("building...")

	rc, err := a.sc.ContentParser().GetRenderContext(ctx)
	if err != nil {
		return
	}

	return a.sc.Renderer().Render(ctx, *rc)
}
