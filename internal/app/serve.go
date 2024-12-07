package app

import (
	"context"
	"path/filepath"

	"github.com/charmbracelet/log"
)

func (b *App) Serve(ctx context.Context, port int) (err error) {

	// make sure to build once
	err = b.Build(ctx)
	if err != nil {
		return
	}

	refreshChan, err := b.srv.Serve(ctx, port)
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
