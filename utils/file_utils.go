package utils

import (
	"DBWorker/lib"
	"os"
	"path/filepath"
	"regexp"
)

type GetData func(int) ([]byte, err)

type DiskOper interface {
	GetCurrentDir() (string, *lib.Error)
	GetApplicationName() (string, *lib.Error)
}

type DiskIO struct {
	filename string
	dir      string
}

func GetApplicationDir() *Dir {
	return &DiskIO{
		filename: filepath.Abs(filepath.Base(os.Args[0])),
		dir:      filepath.Abs(filepath.Dir(os.Args[0])),
	}
}

func (d *Dir) GetCurrentDir() (string, *lib.Error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	return
}

type FileContents struct {
	NoOfTokens int
	Tokens     []string
	Data       []byte
}

func GetFileContents(filename string) (*FileContents, *lib.Error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, *lib.ToError
	}
	contents := &FileContents{}
	return nil, err
}

func (fc *FileContents) GetAllTokens(src string, reg regexp.Regexp) {
	tokens := reg.FindAllString(src, -1)
	fc.noOfTokens = len(tokens)
	fc.tokens = tokens
}
