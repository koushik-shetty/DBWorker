package utils

import (
	"DBWorker/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type GetData func(int) ([]byte, *lib.Error)

type FileOper interface {
	GetCurrentDir() (string, *lib.Error)
	GetApplicationName() (string, *lib.Error)
}

type File struct {
	name string
	dir  string
}

func GetApplicationDir() (*File, *lib.Error) {
	filename, err := filepath.Abs(filepath.Base(os.Args[0]))
	if err != nil {
		return nil, lib.ToLibError(err, lib.DirError, "application directory")
	}

	baseDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, lib.ToLibError(err, lib.DirError, "application directory")
	}

	return &File{
		name: filename,
		dir:  baseDir,
	}, nil
}

func GetCurrentDir() (string, *lib.Error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", lib.ToLibError(err, lib.DirError, "CurrentDirectory")
	}
	return dir, nil
}

type FileContents struct {
	noOfTokens int
	tokens     []string
	data       []byte
}

func NewFileContents(noOfContents int, tokens []string, data []byte) *FileContents {
	return &FileContents{
		noOfTokens: noOfContents,
		tokens:     tokens,
		data:       data,
	}
}

func GetFileContents(file File) (string, *lib.Error) {
	fileBytes, err := ioutil.ReadFile(file.dir + file.name)

	if err != nil {
		return nil, lib.ToLibError(err, lib.FileError, "get file contents")
	}
	return string(fileBytes), nil

	// regex, err := regexp.Compile(":[a-zA-Z][a-zA-Z0-9]+")
	// if err != nil {
	// 	return nil, lib.ToLibError(err, lib.RegexError, "get file contents")
	// }
	// noOfContents, tokens := GetAllTokens(string(fileBytes), regex)
	// contents := NewFileContents(noOfContents, tokens, fileBytes)
	// return contents, nil
}

func GetAllTokens(src string, reg *regexp.Regexp) (int, []string) {
	tokens := reg.FindAllString(src, -1)

	for i, value := range tokens {
		tokens[i] = value[1:]
	}
	noOfTokens := len(tokens)
	return noOfTokens, tokens
}
