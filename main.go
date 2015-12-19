package main

import (
	"db_worker/app"
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
)

func main() {
	flag.Parse()

	if !strings.HasSuffix(*file, ".sql") {
		fmt.Println("Error : only sql files are allowed")
		return
	}

	db := app.DefaultDBConfig()

	switch *operation {
	case setup:
	case teardown:
	case up:
	case down:
	case migration:
		ipFile := *file
		createedFile, err := db.DBCreateMigration(*dir, ipFile[:len(*file)-4])

		if err != nil {
			fmt.Println("Error creating migration:", err)
			return
		}

		fmt.Println("Created Migration file ", createedFile)
	default:
		fmt.Println("Unsupported operation:", *operation)
		return
	}
}
