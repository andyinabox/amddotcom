package unamica

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
)

// BuildSingle will perform a build based on a given file path. It will use the location
// of the file to determine if the entire project should be built, or only a single part
func (a *App) BuildSingle(ctx context.Context, absPath string) error {

	if !filepath.IsAbs(absPath) {
		return fmt.Errorf("renderer expected absolute path: %s", absPath)
	}

	// if it's in the assets directory, we can simply re-copy assets
	if strings.Contains(absPath, a.cnf.AssetsPath) {
		log.Info("copy assets dir")
		return a.sc.Renderer().CopyAssets(ctx)
	}

	// the rest of the options require a render context
	rc, err := a.sc.ContentParser().GetRenderContext(ctx)
	if err != nil {
		return err
	}

	// if the file comes from the source dir we can just render that file
	if strings.Contains(absPath, a.cnf.SrcPath) {
		log.Infof("render file %s", absPath)
		return a.sc.Renderer().RenderSingle(ctx, *rc, absPath)
	}

	// otherwise return everything
	log.Info("render all...")
	return a.sc.Renderer().RenderAll(ctx, *rc)
}
