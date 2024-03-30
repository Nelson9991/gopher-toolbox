package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

type FileReader struct {
	InputPath string
}

func (fm FileReader) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading the file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading the file: %v", err)
	}

	return lines, nil
}

func New(inputPath string) FileReader {
	return FileReader{
		InputPath: inputPath,
	}
}
