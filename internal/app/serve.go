package app

import (
	"context"
	"fmt"
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
		fmt.Printf("changed: %s", change)
		err = b.Build(ctx)
		if err != nil {
			fmt.Printf("build error: %s", err)
		}
		refreshChan <- struct{}{}
	}

	return
}
