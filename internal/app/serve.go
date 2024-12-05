package app

import "context"

func (b *App) Serve(ctx context.Context, port int) (err error) {

	// changes, errs := b.wch.Watch(ctx, []string{
	// 	b.cnf.ContentPath,
	// 	b.cnf.SrcPath,
	// 	b.cnf.AssetsPath,
	// })

	return b.srv.Serve(ctx, port)
}
