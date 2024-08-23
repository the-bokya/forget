package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
)

type Application struct {
	db *sql.DB
}

func main() {
	addFlag := flag.Bool("add", false, "Add new message")
	listFlag := flag.Bool("list", false, "List all current messages by id")
	flag.Parse()
	db, err := InitDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
		return
	}
	defer db.Close()
	app := Application{db}
	if *addFlag {
		err = app.insertMessage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err.Error())
			return
		}
		return
	}
	if *listFlag {
		err = app.list()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err.Error())
			return
		}
		return
	}
	message, err := app.getMessage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("\x1b[1m%v\n", message)
}
