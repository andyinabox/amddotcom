package renderer

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/amddotcom/pkg/unamica"
	"github.com/charmbracelet/log"
)

func (r *Renderer) RenderAll(ctx context.Context, rc unamica.RenderContext) error {

	files, err := os.ReadDir(r.cnf.SrcPath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// render files
	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			absPath := filepath.Join(r.cnf.SrcPath, file.Name())
			renderErr := r.RenderSingle(ctx, rc, absPath)
			if renderErr != nil {
				log.Error("error rendering file", "error", err, "file", absPath)
			}
		}()
	}

	// copy assets
	wg.Add(1)
	go func() {
		defer wg.Done()
		copyErr := r.copyAssets()
		if copyErr != nil {
			log.Error("error copying assets", "error", err)
		}
	}()

	// output context json
	if r.cnf.OutputContextJSON {

		wg.Add(1)

		go func() {
			defer wg.Done()
			data, ctxErr := json.Marshal(rc)
			if ctxErr != nil {
				log.Error(ctxErr)
			}

			contextFilePath := filepath.Join(r.cnf.OutputPath, "context.json")
			log.Debug("output context file", "file", contextFilePath)
			ctxErr = os.WriteFile(contextFilePath, data, os.ModePerm)
			if ctxErr != nil {
				log.Error(ctxErr)
			}
		}()

	}

	wg.Wait()

	return nil
}
