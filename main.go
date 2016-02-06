package main

import (
	"DBWorker/app"
	"DBWorker/utils"
	"flag"
	"fmt"
	"strings"
)

const (
	setup     = "setup"
	teardown  = "teardown"
	up        = "up"
	down      = "down"
	migration = "genmigration"
)

var (
	operation = flag.String("work", "", `-work can have : ["setup","teardown","up","down","genmigration"] values`)
	file      = flag.String("f", "", "Script file to use for the operation")
	dir       = flag.String("dir", "", "Script file directory")
	tokens    = flag.String("tokens", "", "token to replace in the file. Should be the last flag. Values should be token:tokenvalue pairs, seperated by space. ")
)

func main() {
	flag.Parse()

	if !strings.HasSuffix(*file, ".sql") {
		fmt.Println("Error : only sql files are allowed")
		return
	}

	dbc := app.DefaultDBConfig()
	db, err := app.NewDatabase(dbc)
	if err != nil {
		fmt.Printf("Failed to connect ot db : %v", err)
		return
	}

	switch *operation {
	case setup:
		fileData := utils.NewFile(*file, *dir)
		db.DB_Setup(fileData, utils.ToPairs(flag.Args()))
		return
	case teardown:

	case up:
	case down:

	case migration:
		tmpFile := *file
		ipFile := strings.Replace(tmpFile, ".sql", "", -1)
		fileData := utils.NewFile(ipFile, *dir)
		createdFile, err := app.DBCreateMigration(fileData)

		if err != nil {
			fmt.Println("Error creating migration:", err)
			return
		}

		fmt.Println("Created Migration file ", createdFile)

	default:
		fmt.Println("Unsupported operation:", *operation)
		return
	}
}
