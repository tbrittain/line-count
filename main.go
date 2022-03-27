package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"os"
	"path/filepath"
	"sort"
)

// https://golangbot.com/webassembly-using-go/

var storer *memory.Storage
var fs billy.Filesystem

func main() {
	storer = memory.NewStorage()
	fs = memfs.New()

	repository := "https://github.com/tbrittain/portfolio-site"

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:      repository,
		Progress: os.Stdout,
		Tags:     git.NoTags,
		Depth:    1,
	})
	handleErrorFatal(err)

	w, err := r.Worktree()
	handleErrorFatal(err)

	var results []FileResult
	path := "./"

	walkFileTreeAndAppendResults(w.Filesystem, &results, path)
	jsonResults, err := json.Marshal(results)
	handleErrorFatal(err)

	fmt.Println(string(jsonResults))

	aggregatedResults := aggregateResults(results)
	jsonAggregatedResults, err := json.Marshal(aggregatedResults)
	handleErrorFatal(err)

	fmt.Println(string(jsonAggregatedResults))
}

func walkFileTreeAndAppendResults(fileSystem billy.Filesystem, results *[]FileResult, path string) {
	dir, err := fileSystem.ReadDir(path)
	handleErrorFatal(err)

	for _, file := range dir {
		fileName := file.Name()

		if file.IsDir() {
			walkFileTreeAndAppendResults(fileSystem, results, path+fileName+"/")
		} else {
			fileExtension := filepath.Ext(fileName)
			lang, ok := validateFileExtension(fileExtension)
			if !ok {
				continue
			}

			openedFile, err := fs.Open(path + fileName)
			if err != nil {
				fmt.Println("error opening file: ", err)
				return
			}

			numLines, err := countLines(openedFile)
			result := FileResult{
				Filename: fileName,
				Filetype: lang,
				NumLines: numLines,
				Size:     file.Size(),
				Error:    err,
			}
			*results = append(*results, result)
		}
	}
}

func countLines(file billy.File) (int, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return 0, err
	}

	lineCounter := 0
	for range bytes.Split(buf.Bytes(), []byte("\n")) {
		lineCounter++
	}

	return lineCounter, nil
}

func validateFileExtension(fileExtension string) (*Language, bool) {
	lang, ok := LanguageMap[fileExtension]
	if !ok {
		return nil, false
	}

	return &lang, true
}

func aggregateResults(results []FileResult) []AggregatedResult {
	var aggregatedResults []AggregatedResult

	for _, result := range results {
		found := false
		for i, aggregatedResult := range aggregatedResults {
			if aggregatedResult.Filetype.Name == result.Filetype.Name {
				aggregatedResults[i].NumLines += result.NumLines
				aggregatedResults[i].Size += result.Size
				found = true
			}
		}

		if !found {
			aggregatedResults = append(aggregatedResults, AggregatedResult{
				Filetype: result.Filetype,
				NumLines: result.NumLines,
				Size:     result.Size,
			})
		}
	}

	sort.Slice(aggregatedResults, func(i, j int) bool {
		return aggregatedResults[i].NumLines > aggregatedResults[j].NumLines
	})

	return aggregatedResults
}

func handleErrorFatal(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

type FileResult struct {
	Filename string    `json:"filename"`
	Filetype *Language `json:"filetype"`
	NumLines int       `json:"num_lines"`
	Size     int64     `json:"size"`
	Error    error     `json:"error"`
}

type AggregatedResult struct {
	Filetype *Language `json:"filetype"`
	NumLines int       `json:"num_lines"`
	Size     int64     `json:"size"`
}
