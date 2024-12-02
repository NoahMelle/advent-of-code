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
	lines, err := parseFile("./my_input_file.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	validLinesCount := 0
	for _, line := range lines {
		if checkIfSequenceIsValid(parseLine(line), true) {
			validLinesCount++
		}
	}

	fmt.Println(validLinesCount)
}

func parseFile(path string) ([]string, error) {
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
	return lines, scanner.Err()
}

func parseLine(line string) []int {
	fields := strings.Fields(line)
	nums := make([]int, len(fields))
	for i, val := range fields {
		nums[i], _ = strconv.Atoi(val)
	}
	return nums
}

func removeFromSlice(slice []int, s int) []int {
	newSlice := make([]int, len(slice)-1)
	copy(newSlice, slice[:s])
	copy(newSlice[s:], slice[s+1:])
	return newSlice
}

func checkIfSequenceIsValid(sequence []int, allowBadValue bool) bool {
	descending := sequence[0] > sequence[1]
	for i := 1; i < len(sequence); i++ {
		diff := int(math.Abs(float64(sequence[i] - sequence[i-1])))
		if diff > 3 || diff < 1 || (descending && sequence[i] > sequence[i-1]) || (!descending && sequence[i] < sequence[i-1]) {
			if allowBadValue {
				return validateWithOneRemoval(sequence)
			}
			return false
		}
	}
	return true
}

func validateWithOneRemoval(sequence []int) bool {
	for i := range sequence {
		if checkIfSequenceIsValid(removeFromSlice(sequence, i), false) {
			return true
		}
	}
	return false
}
