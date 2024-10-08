package main

import (
	"fmt"
	"os"
)

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
	cantDeleteError struct {
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

func (err *cantDeleteError) Error() string {
	return fmt.Sprintf("There was an error deleting %v\n", err.errorMessage)
}

func displayError(err error) {
	fmt.Fprintf(os.Stderr, "%v", err.Error())
}
