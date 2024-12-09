package compileservice

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/amddotcom/pkg/unamica"
	"github.com/charmbracelet/log"
)

func (s *Service) Compile(ctx context.Context, rc unamica.RenderContext) (err error) {

	files, err := os.ReadDir(s.cnf.SrcPath)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.Name() == "index.html.tmpl" {
			copyErr := s.parseHtmlTemplate(file, rc)
			if copyErr != nil {
				log.Error(copyErr)
			}
		} else {
			copyErr := s.copyFile(file)
			if copyErr != nil {
				log.Error(copyErr)
			}
		}
	}

	err = s.copyAssets()
	if err != nil {
		return
	}

	if s.cnf.OutputContextJSON {
		var data []byte
		data, err = json.Marshal(rc)
		if err != nil {
			return
		}

		contextFilePath := filepath.Join(s.cnf.OutputPath, "context.json")
		log.Debug("output context file", "file", contextFilePath)
		err = os.WriteFile(contextFilePath, data, os.ModePerm)
	}

	return
}
