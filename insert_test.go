package main

import (
	"testing"
)

func TestInsertDB(t *testing.T) {
	message := "This will be stored in the temp db."
	err := app.insertDB(message)
	if err != nil {
		t.Error("Couldn't insert message.")
	}
	var present bool
	err = app.db.QueryRow("SELECT EXISTS (SELECT message FROM messages WHERE message=?)", message).Scan(&present)
	if err != nil {
		t.Errorf("Error retrieving message %v", err.Error())
	}
	if !present {
		t.Error("Message inserted but not present.")
	}
}
