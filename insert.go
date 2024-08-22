package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (app *Application) insertMessage() error {
	fmt.Printf("\x1b[1m")
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
