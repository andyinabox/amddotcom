package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/russross/blackfriday/v2"
)

type jsonData map[string]json.RawMessage
type markdownData struct {
	Raw    string
	Parsed string
}
type outputDataFormat struct {
	JsonFiles     map[string]jsonData
	MarkdownFiles map[string]markdownData
}

func main() {
	var inputPath, outputPath, outputFileName string
	var pretty bool

	flag.StringVar(&inputPath, "i", "", "input dir path")
	flag.StringVar(&outputPath, "o", "publish", "output dir path")
	flag.StringVar(&outputFileName, "n", "content.json", "output file name")
	flag.BoolVar(&pretty, "pretty", true, "pretty print output json")

	flag.Parse()

	err := parseContent(inputPath, outputPath, outputFileName, pretty)
	if err != nil {
		panic(err)
	}

}

func parseContent(inputPath, outputPath, outputFileName string, pretty bool) (err error) {
	if inputPath == "" {
		err = errors.New("no input path provided")
		return
	}

	// get input file info
	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		return
	}

	// make sure it's a directory
	if !fileInfo.IsDir() {
		err = fmt.Errorf("path %s is not a directory", inputPath)
		return
	}

	// traverse directory
	files, err := os.ReadDir(inputPath)
	if err != nil {
		return
	}

	output := outputDataFormat{
		JsonFiles:     make(map[string]jsonData),
		MarkdownFiles: make(map[string]markdownData),
	}

	var fn, fpath string

	for _, file := range files {

		fn = file.Name()
		fpath = filepath.Join(inputPath, fn)

		// skip directories for now
		if file.IsDir() {
			continue
		}

		if filepath.Ext(fn) == ".json" {
			fmt.Printf("json file: %s\n", fn)
			data, err := parseJsonFile(fpath)
			if err != nil {
				fmt.Printf("error %s: %s\n", fn, err)
				continue
			}
			name := getFileName(fn)
			output.JsonFiles[name] = data
		}

		if filepath.Ext(fn) == ".md" {
			fmt.Printf("md file: %s\n", fn)
			data, err := parseMarkdownFile(fpath)
			if err != nil {
				fmt.Printf("error %s: %s\n", fn, err)
				continue
			}
			name := getFileName(fn)
			output.MarkdownFiles[name] = data
		}

	}

	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return
	}

	var b []byte
	if pretty {
		b, err = json.MarshalIndent(output, "", "\t")
		if err != nil {
			return
		}

	} else {
		b, err = json.Marshal(output)
		if err != nil {
			return
		}

	}

	err = os.WriteFile(filepath.Join(outputPath, outputFileName), b, fs.ModePerm)
	return
}

func getFileName(p string) string {
	base := filepath.Base(p)
	parts := strings.Split(base, ".")
	name := parts[0]
	return strcase.ToCamel(name)
}

func parseMarkdownFile(p string) (data markdownData, err error) {
	var b []byte

	b, err = os.ReadFile(p)
	if err != nil {
		return
	}

	parsed := blackfriday.Run(b)

	data = markdownData{
		Raw:    string(b),
		Parsed: string(parsed),
	}

	return
}

func parseJsonFile(p string) (data jsonData, err error) {
	var b []byte
	b, err = os.ReadFile(p)
	if err != nil {
		return
	}
	data = jsonData{}
	err = json.Unmarshal(b, &data)

	return
}
