package app

import (
	"DBWorker/utils"
	"bitbucket.org/liamstask/goose/lib/goose"
	"time"
)

func (db *DBConfig) DBCreateMigration(dir, fileName string) (string, error) {
	useDir := func() (string, error) {

		if dir != "" {
			return utils.GetCurrentDir()
		}
		return dir, nil
	}

	dir, err := useDir()

	if err != nil {
		return "", err
	}

	createdFile, err := goose.CreateMigration(fileName, "sql", dir, time.Now())
	return createdFile, err
}
