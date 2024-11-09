package utils

import (
	"bufio"
	"log"
	"os"
)

func MustReadInput(path string) []string {
	lines, err := ReadInput(path)
	if err != nil {
		log.Fatal("Could not ReadInput", err)
	}
	return lines
}

func ReadInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
