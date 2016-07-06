package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/koushik-shetty/DBWorker/app"
	"github.com/koushik-shetty/DBWorker/utils"
)

const (
	setup     = "setup"
	teardown  = "teardown"
	up        = "up"
	down      = "down"
	migration = "genmigration"
)

var (
	configFile = flag.String("config", "", "config file for the application to run")
	operation  = flag.String("work", "", `-work can have : ["setup","teardown","up","down","genmigration"] values`)
	file       = flag.String("f", "", "Script file to use for the operation")
	dir        = flag.String("dir", "", "Script file directory")
	tokens     = flag.String("tokens", "", "token to replace in the file. Should be the last flag. Values should be token:tokenvalue pairs, seperated by space. ")
)

func main() {
	flag.Parse()
	config, err := app.LoadConfig(*configFile)
	if err != nil {
		fmt.Printf("Error reading config file: %v", err)
		return
	}
	if !strings.HasSuffix(*file, ".sql") {
		fmt.Println("Error : only sql files are allowed")
		return
	}

	switch *operation {
	case setup:
		dbc := config.DBConfig()
		db, err := app.NewDatabase(dbc, dbc.DBSetupString())
		if err != nil {
			fmt.Printf("Failed to connect to db : %v", err)
			return
		}
		fileData := utils.NewFile(*file, *dir)
		e := db.DB_Setup(fileData, utils.ToPairs(flag.Args()))
		if !e.IsEmpty() {
			fmt.Printf("error seting up database: %v", e)
		}
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
