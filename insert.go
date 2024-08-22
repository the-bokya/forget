package main

import (
	"bufio"
	"os"
	"strings"
)

func (app *Application) insertMessage() error {
	messageReader := bufio.NewReader(os.Stdin)
	message, err := messageReader.ReadString('\n')
	message = strings.TrimSpace(message)
	if err != nil {
		return err
	}
	insertQuery := `
	INSERT INTO messages (message) VALUES(?)
	`
	_, err = app.db.Exec(insertQuery, message)
	if err != nil {
		return err
	}
	return nil
}
