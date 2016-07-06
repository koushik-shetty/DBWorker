package utils

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/koushik-shetty/DBWorker/lib"
)

type FileOper interface {
	Name() string
	Dir() string
	FormatContents(pairSring Pairs) (fileContents string, err *lib.Error)
}

type File struct {
	name string
	dir  string
}

func NewFile(name, dir string) *File {
	return &File{
		name: name,
		dir:  dir,
	}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Dir() string {
	return f.dir
}

func (file *File) FormatContents(pairs Pairs) (fileContents string, err *lib.Error) {
	fileContents, err = GetFileContents(file)
	if err != nil {
		return
	}
	if err = pairs.Verify(); err != nil {
		return "", err
	}

	tokens := pairs.ToTokens()
	if err != nil {
		return "", err
	}
	return tokens.stringInterpolate(fileContents), nil
}

func GetFileContents(file FileOper) (string, *lib.Error) {
	fileBytes, err := ioutil.ReadFile(path.Join(file.Dir() + file.Name()))

	if err != nil {
		return "", lib.ToLibError(err, lib.FileError, "get file contents")
	}
	return string(fileBytes), nil
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
