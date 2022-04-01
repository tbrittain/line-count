package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

var storer *memory.Storage
var fs billy.Filesystem

func main() {
	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.PathParameters["url"] == "" {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "No url parameter passed",
		}, nil
	}

	repoUrl := request.PathParameters["url"]
	pattern := `((git|ssh|http(s)?)|(git@[\w\.]+))(:(\/\/)?)([\w\.@\:\/\-~]+)(\.git)(\/)?`
	matched, err := regexp.MatchString(pattern, repoUrl)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	if matched == false {
		return &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid URL",
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	result, err := getRepositorySummary(repoUrl)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	body, err := json.Marshal(result)
	if err != nil {
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func getRepositorySummary(repository string) ([]AggregatedResult, error) {
	storer = memory.NewStorage()
	fs = memfs.New()

	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL:      repository,
		Progress: os.Stdout,
		Tags:     git.NoTags,
		Depth:    1,
	})
	if err != nil {
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	var results []FileResult
	path := "./"

	err = walkFileTreeAndAppendResults(w.Filesystem, &results, path)
	if err != nil {
		return nil, err
	}

	aggregatedResults := aggregateResults(results)
	return aggregatedResults, nil
}

func walkFileTreeAndAppendResults(fileSystem billy.Filesystem, results *[]FileResult, path string) error {
	dir, err := fileSystem.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range dir {
		fileName := file.Name()

		if file.IsDir() {
			err := walkFileTreeAndAppendResults(fileSystem, results, path+fileName+"/")
			if err != nil {
				return err
			}
		} else {
			fileExtension := filepath.Ext(fileName)
			lang, ok := validateFileExtension(fileExtension)
			if !ok {
				continue
			}

			openedFile, err := fs.Open(path + fileName)
			if err != nil {
				return err
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

	return nil
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
