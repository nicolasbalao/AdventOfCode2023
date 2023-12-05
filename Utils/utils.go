package utils

import (
	"bufio"
	"os"
)

func FileScanner(path string) *bufio.Scanner {

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	return fileScanner
}
