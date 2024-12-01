package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func parseFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return nil, err
	} else {
		scanner := bufio.NewScanner(file)

		var content []string

		for scanner.Scan() {
			content = append(content, scanner.Text())
		}

		return content, nil
	}
}
