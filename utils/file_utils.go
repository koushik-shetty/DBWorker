package utils

import (
	"DBWorker/lib"
	"io/ioutil"
	"os"
	"path/filepath"
)

type GetData func(int) ([]byte, *lib.Error)

type FileOper interface {
	Name() string
	Dir() string
}

type File struct {
	name string
	dir  string
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Dir() string {
	return f.dir
}

func NewFile(name, dir string) *File {
	return &File{
		name: name,
		dir:  dir,
	}
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

func GetFileContents(file FileOper) (string, *lib.Error) {
	fileBytes, err := ioutil.ReadFile(file.Dir() + file.Name())

	if err != nil {
		return "", lib.ToLibError(err, lib.FileError, "get file contents")
	}
	return string(fileBytes), nil
}
