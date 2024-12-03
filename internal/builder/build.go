package builder

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (b *Builder) Build() (err error) {
	ctx, err := b.ctx.GetContext()
	if err != nil {
		return
	}

	err = b.cmp.Compile(ctx)
	if err != nil {
		return
	}

	data, err := json.Marshal(ctx)
	if err != nil {
		return
	}

	err = os.WriteFile(filepath.Join(b.cnf.OutputPath, "context.json"), data, os.ModePerm)

	return
}
