package unamica

import (
	"context"

	"github.com/charmbracelet/log"
)

func (a *App) Build(ctx context.Context) (err error) {
	log.Info("building...")

	rc, err := a.sc.ContextService().GetRenderContext(ctx)
	if err != nil {
		return
	}

	return a.sc.CompileService().Compile(ctx, *rc)
}
