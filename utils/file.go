package utils

import "os"

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Mkdir(path string) {
	if PathExists(path) {
		return
	}

	_ = os.MkdirAll(path, os.ModePerm)
}
