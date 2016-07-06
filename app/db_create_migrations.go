package app

import (
	"time"

	"bitbucket.org/liamstask/goose/lib/goose"

	"github.com/koushik-shetty/DBWorker/lib"
	u "github.com/koushik-shetty/DBWorker/utils"
)

func DBCreateMigration(file u.FileOper) (string, *lib.Error) {
	folder := file.Dir()

	if folder == "" {
		currDir, err := u.GetCurrentDir()
		folder = currDir
		if err != nil {
			return "", err
		}
	}

	createdFile, e := goose.CreateMigration(file.Name(), "sql", folder, time.Now())

	if e != nil {
		return "", lib.ToLibError(e, lib.FileCreateError, "create migration file")
	}

	return createdFile, nil
}
