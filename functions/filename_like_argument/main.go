package main

import (
	"bufio"
	"errors"
	"io"
	"os"
)

func ProccessDataInFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	var data string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// processing...
	}

	return data, nil
}

func ProccessData(reader io.Reader) (string, error) {
	if reader == nil {
		return "", errors.New("incorrect reader")
	}

	var data string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// processing...
	}

	return data, nil
}
