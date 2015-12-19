package app

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"db_worker/utils"
	"time"
)

func (db *DBConfig) DBCreateMigration(dir, fileName string) (string, error) {
	useDir := func() (string, error) {

		if dir != "" {
			return utils.GetCurrDir()
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
