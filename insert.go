package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var insertQuery string = `
INSERT INTO messages (message) VALUES(?)
`

var endCommand string = "/done"

func (app *Application) insertMessage() error {
	message, err := insertInput()
	if err != nil {
		return &insertError{err.Error()}
	}
	err = app.insertDB(message)
	if err != nil {
		return &insertError{err.Error()}
	}
	return nil
}

func (app *Application) insertDB(message string) error {
	_, err := app.db.Exec(insertQuery, message)
	if err != nil {
		return err
	}
	return nil
}

func insertInput() (string, error) {
	fmt.Printf("Type away! When you are done, type %q on a new line: \n", endCommand)
	bold()
	defer unbold()
	reader := bufio.NewReader(os.Stdin)
	input := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		line = strings.TrimSpace(line)
		if line == "/done" {
			break
		}
		input = append(input, line)
	}
	return strings.Join(input, "\n"), nil
}
