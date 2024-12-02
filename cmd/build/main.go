package main

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/amddotcom/internal"
	"github.com/amddotcom/internal/buildservice"
	"github.com/amddotcom/internal/contextservice"
)

func main() {
	var contextService internal.ContextService
	var buildService internal.BuildService

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(cwd, "publish")
	err = os.MkdirAll(outputPath, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	contextService = contextservice.New(
		&contextservice.Config{
			ContentPath:      filepath.Join(cwd, "content"),
			SiteDataFileName: "data.json",
		},
	)

	buildService = buildservice.New(
		&buildservice.Config{
			SrcPath:    filepath.Join(cwd, "src"),
			OutputPath: outputPath,
		},
	)

	ctx, err := contextService.GetContext()
	if err != nil {
		panic(err)
	}

	err = buildService.Build(ctx)
	if err != nil {
		panic(err)
	}

}
