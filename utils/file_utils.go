package utils

import (
	"os"
	"path/filepath"
)

type GetData func(int) ([]byte,err)

func GetCurrentDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

type FileContents struct {
	noOfTokens int
	tokens     []string
	data       GetData
}

func GetFileContents(filename string) (*FileContents, error) {
	file, err := os.Open(filename)

	getter := func (n int) []byte{
		buf ;= make([]byte,n)
		_, err := file.Read(buf)
		return buf,err
	}

	contents := &FileContents{

	}
	return nil, err
}

