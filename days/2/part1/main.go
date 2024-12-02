package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	validLinesCount := 0
	for _, line := range lines {
		if checkIfSequenceIsValid(parseLine(line)) {
			validLinesCount++
		}
	}

	fmt.Print(validLinesCount)
}

func parseLine(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, val := range fields {
		nums[i], _ = strconv.Atoi(val)
	}
	return nums
}

func checkIfSequenceIsValid(sequence []int) bool {
	descending := sequence[0] > sequence[1]
	for i := 1; i < len(sequence); i++ {
		diff := int(math.Abs(float64(sequence[i] - sequence[i-1])))
		if diff > 3 || diff < 1 || (descending && sequence[i] > sequence[i-1]) || (!descending && sequence[i] < sequence[i-1]) {
			return false
		}
	}
	return true
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
