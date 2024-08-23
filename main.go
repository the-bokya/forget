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
	listFlag := flag.Bool("list", false, "List all current messages by id")
	deleteFlag := flag.Int("delete", -1, "Delete a message by id")
	flag.Parse()
	db, err := InitDB()
	if err != nil {
		displayError(err)
		return
	}
	defer db.Close()
	app := Application{db}
	if *addFlag {
		err = app.insertMessage()
		if err != nil {
			displayError(err)
			return
		}
		return
	}
	if *listFlag {
		err = app.list()
		if err != nil {
			displayError(err)
			return
		}
		return
	}
	if *deleteFlag > -1 {
		err = app.deleteMessage(*deleteFlag)
		if err != nil {
			displayError(err)
			return
		} else {
			fmt.Printf("Deleted id %d successfully\n", *deleteFlag)
		}
		return
	}
	message, err := app.getMessage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bold()
	defer unbold()
	fmt.Println(message)
}
