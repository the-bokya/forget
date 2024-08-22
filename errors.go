package main

import "fmt"

type cantMakeDirError struct {
	dir string
}

func (err *cantMakeDirError) Error() string {
	return fmt.Sprintf("Can't make directory %v\n", err.dir)
}
