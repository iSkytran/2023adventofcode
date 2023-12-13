package utilities

import (
	"bufio"
	"os"
)

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func OpenFile(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	ErrorCheck(err)
	return bufio.NewScanner(file), file
}
