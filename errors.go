package main

import "fmt"

type (
	cantMakeDirError struct {
		dir string
	}

	insertError struct {
		errorMessage string
	}

	listError struct {
		errorMessage string
	}
)

func (err *cantMakeDirError) Error() string {
	return fmt.Sprintf("Can't make directory %v\n", err.dir)
}

func (err *insertError) Error() string {
	return fmt.Sprintf("There was an error inserting %v\n", err.errorMessage)
}

func (err *listError) Error() string {
	return fmt.Sprintf("There was an error listing %v\n", err.errorMessage)
}
