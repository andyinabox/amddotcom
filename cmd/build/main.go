package main

import (
	"flag"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/amddotcom/internal/builder"
)

func main() {
	var err error

	cnf := &builder.Config{}

	flag.StringVar(&cnf.OutputPath, "o", "publish", "path to publish to")
	flag.StringVar(&cnf.ContentPath, "content", "content", "location of site content")
	flag.StringVar(&cnf.SrcPath, "src", "src", "location of site source code")
	flag.StringVar(&cnf.SiteDataFileName, "sitedatafile", "data.json", "name of file with site data")
	flag.StringVar(&cnf.DateFormat, "dateformat", "January 2, 2006", "date format for human-readable build date")

	flag.Parse()

	// get absolute paths
	cnf.OutputPath, err = filepath.Abs(cnf.OutputPath)
	if err != nil {
		panic(err)
	}
	cnf.ContentPath, err = filepath.Abs(cnf.ContentPath)
	if err != nil {
		panic(err)
	}
	cnf.SrcPath, err = filepath.Abs(cnf.SrcPath)
	if err != nil {
		panic(err)
	}

	// make sure output path exists
	err = os.MkdirAll(cnf.OutputPath, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	builder := builder.New(cnf)

	err = builder.Build()
	if err != nil {
		panic(err)
	}

}
