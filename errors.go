package main

import "fmt"

type cantMakeDirError struct {
	dir string
}

type insertError struct {
	errorMessage string
}

func (err *cantMakeDirError) Error() string {
	return fmt.Sprintf("Can't make directory %v\n", err.dir)
}

func (err *insertError) Error() string {
	return fmt.Sprintf("There was an error inserting %v\n", err.errorMessage)
}
