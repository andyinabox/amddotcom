package unamica

import (
	"context"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func (a *App) Serve(ctx context.Context, port int) (err error) {

	snippet := a.sc.Server().GetLiveReloadSnippet()
	log.Debug("set livereload snippet", "snippet", snippet)
	ctx = context.WithValue(ctx, LiveReloadSnippetKey, snippet)

	// make sure to build once
	err = a.Build(ctx)
	if err != nil {
		return
	}

	refreshChan, err := a.sc.Server().Serve(ctx)
	if err != nil {
		return
	}

	changesChan, err := a.sc.Watcher().Watch(ctx, []string{
		a.cnf.ContentPath,
		a.cnf.SrcPath,
		a.cnf.AssetsPath,
	})
	if err != nil {
		return
	}

	// loop over changes in range loop
	for change := range changesChan {
		log.Debug("changed", "file", change)
		err = a.BuildSingle(ctx, change)
		if err != nil {
			log.Error("error building after changes", "error", err)
		}

		// only send refresh if there is a livereload connection
		// NOTE: Server should really do this under the hood
		if a.sc.Server().HasLiveReloadConnection() {
			log.Info("refresh browser")
			refreshChan <- filepath.Base(change)
		}
	}

	return
}
