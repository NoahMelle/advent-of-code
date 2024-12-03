package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	start := time.Now()
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}
	matches := getValidMatches(lines)
	res := 0

	for i := range matches {
		res += addUpNumbers(matches[i])
	}

	fmt.Printf("The total sum is %v", res)

	elapsed := time.Since(start)
	fmt.Printf("The program took %s to execute", elapsed)
}

func getValidMatches(input string) []string {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(input, -1)

	return matches
}

func addUpNumbers(input string) int {
	var num1 int
	var num2 int

	_, err := fmt.Sscanf(input, "mul(%d,%d)", &num1, &num2)

	if err != nil {
		panic(err)
	}

	return num1 * num2
}

func parseFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var content string

	for scanner.Scan() {
		content += scanner.Text()
	}

	return content, nil
}
