package main

import (
	"fmt"
	"os"
)

var app = Application{}

func init() {
	fmt.Println("Meow")
	os.Setenv("FORGET_PATH", "/tmp/forget_test")
	db, err := InitDB()
	if err != nil {
		fmt.Println("Problem setting up data")
	}
	app.db = db

}
