package app

import "context"

func (b *App) Serve(ctx context.Context, port int) (err error) {
	return b.srv.Serve(ctx, port)
}
