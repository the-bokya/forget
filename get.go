package main

func (app *Application) getMessage() (string, error) {
	selectQuery := `
	SELECT message FROM messages ORDER BY RANDOM() LIMIT 1;
	`
	var message string
	err := app.db.QueryRow(selectQuery).Scan(&message)
	if err != nil {
		return "", err
	}
	return message, nil
}
