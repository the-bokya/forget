package main

import (
	"database/sql"
	"flag"
	"fmt"
)

type Application struct {
	db *sql.DB
}

func main() {
	addFlag := flag.Bool("add", false, "Add new message")
	flag.Parse()
	db, err := InitDB()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	app := Application{db}
	if *addFlag {
		app.insertMessage()
		return
	}
	message, err := app.getMessage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(message)
}
