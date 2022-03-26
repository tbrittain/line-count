package main

import (
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
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

	navigateDir(dir)
}

func navigateDir(dir []os.FileInfo) {
	for _, file := range dir {
		fmt.Println("file: ", file.Name())
		fmt.Println("file.IsDir(): ", file.IsDir())

		if file.IsDir() {
			fmt.Println("inside if")
			newDir, err := fs.ReadDir(file.Name())
			handleError(err)
			fmt.Println("newDir: ", newDir)
			navigateDir(newDir)
		} else {
			fmt.Println(file.Name())
		}
	}
}

//func lineCounter(r io.Reader) (int, error) {
//  buf := make([]byte, 32*1024)
//  count := 0
//  lineSep := []byte{'\n'}
//
//  for {
//    c, err := r.Read(buf)
//    count += bytes.Count(buf[:c], lineSep)
//
//    switch {
//    case err == io.EOF:
//      return count, nil
//
//    case err != nil:
//      return count, err
//    }
//  }
//}

func handleError(err error) {
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
