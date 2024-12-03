package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
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

	fmt.Printf("The total sum is %v", res)
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
	splittedNumbers := strings.Split(input, ",")

	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		panic(err)
	}

	var numbers []int

	for i := range splittedNumbers {
		convertedInt, err := strconv.Atoi(reg.ReplaceAllString(splittedNumbers[i], ""))

		if err != nil {
			panic(err)
		}

		numbers = append(numbers, convertedInt)
	}

	return numbers[0] * numbers[1]
}

func parseFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return "", err
	} else {
		scanner := bufio.NewScanner(file)

		var content string

		for scanner.Scan() {
			content += scanner.Text()
		}

		return content, nil
	}
}
