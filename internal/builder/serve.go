package builder

import "context"

func (b *Builder) Serve(ctx context.Context) (err error) {
	return b.srv.Serve(ctx)
}
