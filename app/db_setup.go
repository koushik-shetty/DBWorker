package app

import (
	//	"database/sql"

	"fmt"

	//	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	//	"bitbucket.org/liamstask/goose/lib/goose"

	"github.com/koushik-shetty/DBWorker/lib"
	"github.com/koushik-shetty/DBWorker/utils"
)

type DbSetup interface {
	VerfiyTokens()
}

func (db *Database) DB_Setup(file utils.FileOper, tokenPairs utils.Pairs) (err *lib.Error) {
	//prepare file
	fileContents, err := file.FormatContents(tokenPairs)
	if err != nil {
		fmt.Printf("script error:%v\n", err)
		return err
	}
	err = db.ExecuteFile(fileContents)
	if !err.IsEmpty() {
		return
	}
	return
}

func (d *Database) ExecuteFile(fileContents string) *lib.Error {
	_, err := d.Exec(fileContents)
	if err != nil {

		return lib.NewError(lib.DBError, "dbsetup", fmt.Sprintf("db exe failed with %v", err))
	}
	return nil
}
