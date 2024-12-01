package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const filePath string = "./my_input_file.txt"

	var result []string = parseFile(filePath)

	fmt.Print(result)
}

func parseFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return make([]string, 0)
	} else {
		scanner := bufio.NewScanner(file)

		var content []string = make([]string, 0)

		for scanner.Scan() {
			content = append(content, scanner.Text())
		}

		return content
	}
}
