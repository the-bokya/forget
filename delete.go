package main

func (app *Application) deleteMessage(id int) error {
	deleteQuery := `
	DELETE FROM messages WHERE id=?`
	_, err := app.db.Exec(deleteQuery, id)
	if err != nil {
		return &cantDeleteError{err.Error()}
	}
	return nil
}
