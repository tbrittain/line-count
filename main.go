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
)

// use https://github.com/go-git/go-git
// https://github.com/go-git/go-git/pull/446/files#diff-15808dd1f39f7d3198c9803a02fc1222b866ad5705b5aea887bb6a89ad572223
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
	})

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	w, err := r.Worktree()
	handleError(err)

	dir, err := w.Filesystem.ReadDir(".")
	handleError(err)

	var results []result
	path := "./"

	navigateDir(dir, &results, path)
	fmt.Printf("%v", results)
}

func navigateDir(dir []os.FileInfo, results *[]result, path string) {
	for _, file := range dir {
		fileName := file.Name()

		if file.IsDir() {

			newDir, err := fs.ReadDir(fileName)
			handleError(err)

			navigateDir(newDir, results, path+fileName+"/")
		} else {
			openedFile, err := fs.OpenFile(path+fileName, os.O_RDONLY, 0666)
			if err != nil {
				fmt.Println("error opening file: ", err)
				return
			}

			numLines, err := countLines(openedFile)
			*results = append(*results, result{file.Name(), numLines, file.Size(), err})
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

func handleError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}

type result struct {
	Name     string
	NumLines int
	Size     int64
	Error    error
}
