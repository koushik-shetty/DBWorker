package utils

import (
	"os"
	"path/filepath"
)

func GetCurrDir() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}
