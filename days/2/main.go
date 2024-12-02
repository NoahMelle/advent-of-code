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

	splittedRes := splitParsedFile(lines)
	var validLines [][]int

	for _, line := range splittedRes {
		validRes := checkIfSequenceIsValid(line)
		if validRes {
			validLines = append(validLines, line)
		}
	}

	fmt.Print(len(validLines))
}

func splitParsedFile(parseInput []string) [][]int {
	var res [][]int

	for x := range parseInput {
		var newArr []int
		splittedInput := strings.Split(parseInput[x], " ")
		for y := range splittedInput {
			newArr = append(newArr, handleStringToIntConversion(splittedInput[y]))
		}
		res = append(res, newArr)
	}

	return res
}

func checkIfSequenceIsValid(sequence []int) bool {
	var valid bool = true
	var prev int = sequence[0]
	var descending bool = sequence[0] > sequence[1]

	for i := range sequence {
		difference := int(math.Abs(float64(prev - sequence[i])))
		if (difference > 3 || difference < 1) && i != 0 {
			fmt.Printf("The character %v made the sequence %v fail because it the difference between the previous number was too large or small \n", sequence[i], sequence)
			valid = false
			break
		}

		if prev < sequence[i] && descending {
			fmt.Printf("The character %v made the sequence %v fail because it wasn't descending \n", sequence[i], sequence)
			valid = false
			break
		}

		if prev > sequence[i] && !descending {
			fmt.Printf("The character %v made the sequence %v fail because it wasn't ascending \n", sequence[i], sequence)
			valid = false
			break
		}

		prev = sequence[i]
	}

	return valid
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

func handleStringToIntConversion(toBeConverted string) int {
	res, err := strconv.Atoi(toBeConverted)
	if err != nil {
		panic(err)
	}

	return res
}
