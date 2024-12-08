package app

import (
	"context"
	"path/filepath"

	"github.com/amddotcom/internal"
	"github.com/charmbracelet/log"
)

func (a *App) Serve(ctx context.Context, port int) (err error) {

	snippet := a.srv.GetLiveReloadSnippet()
	log.Debug("set livereload snippet", "snippet", snippet)
	ctx = context.WithValue(ctx, internal.LiveReloadSnippetKey, snippet)

	// make sure to build once
	err = a.Build(ctx)
	if err != nil {
		return
	}

	refreshChan, err := a.srv.Serve(ctx)
	if err != nil {
		return
	}

	changesChan, err := a.wch.Watch(ctx, []string{
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
		err = a.Build(ctx)
		if err != nil {
			log.Error("error building after changes", "error", err)
		}
		// don't send the whole path, just the file name
		log.Info("refresh browser")
		refreshChan <- filepath.Base(change)
	}

	return
}
