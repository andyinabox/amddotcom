package main

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/amddotcom/internal/builder"
	"github.com/urfave/cli/v3"
)

func main() {
	cnf := &builder.Config{}

	defaultFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "o",
			Value:       "publish",
			Usage:       "path to publish to",
			Destination: &cnf.OutputPath,
		},
		&cli.StringFlag{
			Name:        "content",
			Value:       "content",
			Usage:       "location of site content",
			Destination: &cnf.ContentPath,
		},
		&cli.StringFlag{
			Name:        "src",
			Value:       "src",
			Usage:       "location of site source code",
			Destination: &cnf.SrcPath,
		},
		&cli.StringFlag{
			Name:        "assets",
			Value:       "assets",
			Usage:       "location of assets dir",
			Destination: &cnf.AssetsPath,
		},

		&cli.StringFlag{
			Name:        "sitedatafile",
			Value:       "data.json",
			Usage:       "name of file with site data",
			Destination: &cnf.SiteDataFileName,
		},
		&cli.StringFlag{
			Name:        "dateformat",
			Value:       "January 2, 2006",
			Usage:       "date format for human-readable build date",
			Destination: &cnf.DateFormat,
		},
		&cli.BoolFlag{
			Name:        "contextjson",
			Value:       false,
			Usage:       "output context.json",
			Destination: &cnf.OutputContextJSON,
		},
	}

	preparePaths := func(cnf *builder.Config) (err error) {
		// get absolute paths
		cnf.OutputPath, err = filepath.Abs(cnf.OutputPath)
		if err != nil {
			return
		}
		cnf.ContentPath, err = filepath.Abs(cnf.ContentPath)
		if err != nil {
			return
		}
		cnf.SrcPath, err = filepath.Abs(cnf.SrcPath)
		if err != nil {
			return
		}
		cnf.AssetsPath, err = filepath.Abs(cnf.AssetsPath)
		if err != nil {
			return
		}

		// make sure output path exists
		err = os.MkdirAll(cnf.OutputPath, fs.ModePerm)
		return
	}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			// build command
			{
				Name:  "build",
				Usage: "build the site",
				Flags: defaultFlags,
				Action: func(ctx context.Context, cmd *cli.Command) (err error) {

					preparePaths(cnf)

					builder := builder.New(cnf)

					err = builder.Build(ctx)
					return
				},
			},
			// serve command
			{
				Name:  "serve",
				Usage: "watch and serve the site",
				Flags: defaultFlags,
				Action: func(ctx context.Context, cmd *cli.Command) (err error) {
					return errors.New("not implemented")
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
