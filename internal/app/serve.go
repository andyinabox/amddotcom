package app

import (
	"context"
	"path/filepath"

	"github.com/amddotcom/internal"
	"github.com/charmbracelet/log"
)

func (b *App) Serve(ctx context.Context, port int) (err error) {

	snippet := b.srv.GetLiveReloadSnippet()
	log.Debug("set livereload snippet", "snippet", snippet)
	ctx = context.WithValue(ctx, internal.LiveReloadSnippetKey, snippet)

	// make sure to build once
	err = b.Build(ctx)
	if err != nil {
		return
	}

	refreshChan, err := b.srv.Serve(ctx)
	if err != nil {
		return
	}

	changesChan, err := b.wch.Watch(ctx, []string{
		b.cnf.ContentPath,
		b.cnf.SrcPath,
		b.cnf.AssetsPath,
	})
	if err != nil {
		return
	}

	// loop over changes in range loop
	for change := range changesChan {
		log.Debug("changed", "file", change)
		err = b.Build(ctx)
		if err != nil {
			log.Error("error building after changes", "error", err)
		}
		// don't send the whole path, just the file name
		log.Info("refresh browser")
		refreshChan <- filepath.Base(change)
	}

	return
}

// func (a *App) buildForReload(ctx context.Context, liveReloadSnippet string) (err error) {
// 	log.Info("building...")

// 	rc, err := a.ctx.GetRenderContext(ctx)
// 	if err != nil {
// 		return
// 	}

// 	rc.LiveReloadSnippet = template.HTML(liveReloadSnippet)

// 	err = a.cmp.Compile(ctx, *rc)
// 	return

// }
