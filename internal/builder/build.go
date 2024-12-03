package builder

func (b *Builder) Build() (err error) {
	ctx, err := b.ctx.GetContext()
	if err != nil {
		return
	}

	return b.cmp.Compile(ctx)
}
