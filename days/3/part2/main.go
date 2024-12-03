package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("error reading file: %v \n", err)
		panic(err)
	}

	doBlocks := splitTextIntoDoBlocks(lines)

	res := 0

	for i := range doBlocks {
		matches := getValidMatches(doBlocks[i])
		for i := range matches {
			res += addUpNumbers(matches[i])
		}
	}

	fmt.Printf("The total sum is %v \n", res)
	elapsed := time.Since(start)
	fmt.Printf("The program took %s to execute", elapsed)
}

func getValidMatches(input string) []string {
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(input, -1)

	return matches
}

func splitTextIntoDoBlocks(input string) []string {
	doBlocks := strings.Split(input, "do()")

	for i := range doBlocks {
		doBlocks[i] = strings.Split(doBlocks[i], "don't()")[0]
	}

	return doBlocks
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

	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}
	content := builder.String()

	return content, nil
}
