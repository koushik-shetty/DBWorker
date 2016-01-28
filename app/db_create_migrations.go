package app

import (
	"DBWorker/lib"
	"DBWorker/utils"
	"bitbucket.org/liamstask/goose/lib/goose"
	"time"
)

func (db *DBConfig) DBCreateMigration(file FileOper) (string, *lib.Error) {
	folder := ""
	err := lib.EmptyError()
	if dir == "" {
		folder, err = file.GetCurrentDir()
		if err != nil {
			return "", err
		}
	}

	createdFile, err := goose.CreateMigration(file.Name, "sql", folder, time.Now())

	if err != nil {
		return "", lib.ToLibError(err, lib.FileCreateError, "create migration file")
	}

	return createdFile, nil
}
