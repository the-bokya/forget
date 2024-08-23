package main

import "fmt"

type message struct {
	id      int
	message string
}

func (app *Application) list() error {
	bold()
	fmt.Printf("id\tmessage\n")
	fmt.Printf("--\t-------\n")
	unbold()
	messages, err := app.listDB()
	if err != nil {
		return &listError{err.Error()}
	}
	for _, messages := range messages {
		fmt.Printf("%d\t%q\n", messages.id, messages.message)
	}
	return nil

}

func (app *Application) listDB() ([]message, error) {
	messages := make([]message, 0)
	listQuery := `
	SELECT id, message FROM messages ORDER BY id DESC`
	rows, err := app.db.Query(listQuery)
	if err != nil {
		return nil, err
	}
	for {
		if !rows.Next() {
			break
		}
		currentMessage := message{}
		err = rows.Scan(&currentMessage.id, &currentMessage.message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, currentMessage)
	}
	return messages, nil
}
