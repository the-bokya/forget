package main

import (
	"os"
	"path"
)

func InitPath() (string, error) {
	dir := os.Getenv("FORGET_PATH")
	if dir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dir = path.Join(homeDir, ".config/forget/")
	}
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, os.FileMode(0755))
			if err != nil {
				return "", &cantMakeDirError{dir}
			}
		} else {
			return "", err
		}
	}
	dir = path.Join(dir, "forget.db")
	return dir, nil
}
