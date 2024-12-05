package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	const filePath string = "./my_input_file.txt"

	rules, updates, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	validSequences := [][]int{}
	res := 0

	for _, update := range updates {
		if checkIfSequenceIsValid(update, rules) {
			validSequences = append(validSequences, update)
		}
	}

	fmt.Printf("Valid sequences: %v", validSequences)

	for _, validSequence := range validSequences {
		median := median(validSequence)
		res += median
	}

	fmt.Println("The median is ", res)

	duration := time.Since(start)
	fmt.Printf("\nExecuted the program in %s\n", duration)
}

func checkIfSequenceIsValid(sequence []int, rules [][]int) bool {
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		indexX := indexOf(sequence, x)
		indexY := indexOf(sequence, y)

		if indexX != -1 && indexY != -1 {
			if indexX >= indexY {
				return false
			}
		}
	}
	return true
}

func median(data []int) int {
	dataCopy := make([]int, len(data))
	copy(dataCopy, data)

	var median int
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}

func indexOf(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func parseFile(path string) ([][]int, [][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var rules, updates [][]int
	scanner := bufio.NewScanner(file)
	isUpdateSection := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isUpdateSection = true
			continue
		}

		var numbers []int
		var separator string

		if isUpdateSection {
			separator = ","
		} else {
			separator = "|"
		}

		parts := strings.Split(line, separator)
		for _, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return nil, nil, err
			}
			numbers = append(numbers, num)
		}

		if isUpdateSection {
			updates = append(updates, numbers)
		} else {
			rules = append(rules, numbers)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}
