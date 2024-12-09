package renderer

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/amddotcom/pkg/unamica"
	"github.com/charmbracelet/log"
)

func (r *Renderer) Render(ctx context.Context, rc unamica.RenderContext) (err error) {

	files, err := os.ReadDir(r.cnf.SrcPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.Name() == "index.html.tmpl" {
			copyErr := r.parseHtmlTemplate(file, rc)
			if copyErr != nil {
				log.Error(copyErr)
			}
		} else {
			copyErr := r.copyFile(file)
			if copyErr != nil {
				log.Error(copyErr)
			}
		}
	}

	err = r.copyAssets()
	if err != nil {
		return
	}

	if r.cnf.OutputContextJSON {
		var data []byte
		data, err = json.Marshal(rc)
		if err != nil {
			return
		}

		contextFilePath := filepath.Join(r.cnf.OutputPath, "context.json")
		log.Debug("output context file", "file", contextFilePath)
		err = os.WriteFile(contextFilePath, data, os.ModePerm)
	}

	return
}
