package renderer

import (
	"context"
	"path/filepath"

	"github.com/amddotcom/pkg/unamica"
)

func (r *Renderer) RenderSingle(ctx context.Context, rc unamica.RenderContext, absPath string) error {

	// TODO: allow for non-html template files
	if filepath.Ext(absPath) == ".tmpl" {
		return r.parseHtmlTemplate(absPath, rc)
	}

	return r.copyFile(absPath)
}
