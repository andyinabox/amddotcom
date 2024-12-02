package main

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/amddotcom"
	"github.com/amddotcom/internal/buildservice"
	"github.com/amddotcom/internal/contextservice"
)

func main() {
	var contextService amddotcom.ContextService
	var buildService amddotcom.BuildService

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
		amddotcom.ContentFS,
		&contextservice.Config{
			ContentPath:      "content",
			SiteDataFileName: "data.json",
		},
	)

	buildService = buildservice.New(
		amddotcom.SrcFS,
		&buildservice.Config{
			SrcPath:    "src",
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
