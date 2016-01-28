package main

import (
	"DBWorker/app"
	"flag"
	"fmt"
	"os"
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
	fmt.Printf("no of args:%d", len(os.Args))

	inputFile := (*file)[:len()]
	db := app.DefaultDBConfig()

	switch *operation {
	case setup:
		app.DB_Setup(utils.File{
			name: *file,
			dir:  *dir,
		}, flag.Args()...)

	case teardown:
	case up:
	case down:
	case migration:
		tmpFile := *file
		ipFile := tmpFile(tmpFile, ".sql", "", -1)
		createdFile, err := db.DBCreateMigration(*dir, ipFile)

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
