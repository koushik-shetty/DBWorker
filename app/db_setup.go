package app

import (
	//	"database/sql"
	"errors"
	"fmt"

	"DBWorker/lib"
	"DBWorker/utils"

	//	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	//	"bitbucket.org/liamstask/goose/lib/goose"
)

type DbSetup interface {
	VerfiyTokens()
}

func (db *Database) DB_Setup(file utils.FileOper, tokenPairs utils.Pairs) (err *lib.Error) {
	//prepare file
	fileContents, err := file.FormatContents(tokenPairs)
	fmt.Printf("contents: %v\n", fileContents)
	if err != nil {
		fmt.Printf("file contents err:%v", err)
		return err
	}
	id := 13
	name := ""
	e := db.QueryRow("select name from sc.test_table where id=$1", id).Scan(&name)
	if e != nil {
		fmt.Printf("\ngot err:%v\n", e)
	}
	fmt.Printf("\ngot value : %v\n", name)
	//run it against psql client

	//handle the results
	return nil
}

func (d *Database) ExecuteAgainstDB(file string) error {
	return errors.New("")
}
