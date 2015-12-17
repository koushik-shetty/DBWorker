package main

import (
	"flag"
	"fmt"
)

var (
	operation = flag.String("work", "", `-work can have : ["setup","teardown","up","down"] values`)
	file      = flag.String("f", "", "Script file to use for the operation")
)

func main() {
	flag.Parse()
	fmt.Println("op:", *operation, "\nf:", *file)
}
