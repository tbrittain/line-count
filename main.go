package main

import (
	"bytes"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"os"
	"path/filepath"
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

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	w, err := r.Worktree()
	handleErrorFatal(err)

	var results []FileResult
	path := "./"

	walkFileTreeAndAppendResults(w.Filesystem, &results, path)
	fmt.Printf("%v", results)
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
			lang, ok := testFileExtension(fileExtension)
			if !ok {
				continue
			}

			openedFile, err := fs.Open(path + fileName)
			if err != nil {
				fmt.Println("error opening file: ", err)
				return
			}

			numLines, err := countLines(openedFile)
			*results = append(*results, FileResult{file.Name(), lang, numLines, file.Size(), err})
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
	fmt.Println("lineCounter: ", lineCounter)
	return lineCounter, nil
}

func testFileExtension(fileExtension string) (*Language, bool) {
	for _, lang := range SupportedLanguages {
		for _, extension := range lang.Extensions {
			if extension == fileExtension {
				return &lang, true
			}
		}
	}
	return nil, false
}

func handleErrorFatal(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

type FileResult struct {
	Filename string
	Filetype *Language
	NumLines int
	Size     int64
	Error    error
}
