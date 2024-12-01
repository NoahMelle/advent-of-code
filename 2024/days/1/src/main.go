package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	var parseResult []string = parseFile(filePath)
	var firstCol, secondCol []int = splitParsedFile(parseResult)
	sort.Ints(firstCol)
	sort.Ints(secondCol)

	var differences = make([]int, 0)

	for i := 0; i < len(firstCol); i++ {
		differences = append(differences, calcDifferenceBetweenNumbers(firstCol[i], secondCol[i]))
	}

	fmt.Printf("The total differences between the lists are %v", sumArray(differences))
}

func splitParsedFile(arr []string) ([]int, []int) {
	var res1, res2 = make([]int, 0), make([]int, 0)

	for i := 0; i < len(arr); i++ {
		splitRes := strings.Split(arr[i], "   ")
		res1 = append(res1, handleStringToIntConversion(splitRes[0]))
		res2 = append(res2, handleStringToIntConversion(splitRes[1]))
	}

	return res1, res2
}

func handleStringToIntConversion(toBeConverted string) int {
	res, err := strconv.Atoi(toBeConverted)

	if err != nil {
		panic(err)
	}

	return res
}

func calcDifferenceBetweenNumbers(num1 int, num2 int) int {
	return int(math.Abs(float64(num1 - num2)))
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
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
