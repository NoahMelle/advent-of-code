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

	validUpdatedSequences := [][]int{}
	res := 0

	for _, update := range updates {
		if !checkIfSequenceIsValid(update, rules) {
			validUpdatedSequences = append(validUpdatedSequences, orderSequence(update, rules))
		}
	}

	for _, validSequence := range validUpdatedSequences {
		median := median(validSequence)
		res += median
	}

	fmt.Println(res)

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

func orderSequence(sequence []int, rules [][]int) []int {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(sequence); i++ {
			current := sequence[i]
			for _, rule := range rules {
				if rule[0] == current {
					next := rule[1]
					indexNext := indexOf(sequence, next)
					if indexNext != -1 && indexNext < i {
						// move rule[1] right after i
						sequence = move(sequence, indexNext, i+1)
						changed = true
					}
				}
			}
		}
	}

	fmt.Println("Ordered sequence:", sequence)
	return sequence
}

func move(slice []int, fromIndex, toIndex int) []int {
	element := slice[fromIndex]
	// remove element from slice
	slice = append(slice[:fromIndex], slice[fromIndex+1:]...)
	if toIndex >= len(slice) {
		// append at end of slice
		slice = append(slice, element)
	} else {
		// split slice and add element at right index,
		slice = append(slice[:toIndex], append([]int{element}, slice[toIndex:]...)...)
	}
	return slice
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
