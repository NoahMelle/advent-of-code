package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	grid, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		panic(err)
	}

	findMiddleLetter(grid)
}

func findMiddleLetter(grid [][]string) {
	count := 0
	for y := 1; y < len(grid)-1; y++ {
		row := grid[y]
		for x := 1; x < len(grid[y])-1; x++ {
			if row[x] == "A" {
				upperLeftToLowerRight := [2]string{grid[y-1][x-1], grid[y+1][x+1]}
				upperRightToLowerLeft := [2]string{grid[y-1][x+1], grid[y+1][x-1]}
				if slices.Contains(upperLeftToLowerRight[:], "M") && slices.Contains(upperLeftToLowerRight[:], "S") && slices.Contains(upperRightToLowerLeft[:], "M") && slices.Contains(upperRightToLowerLeft[:], "S") {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}

func parseFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	return grid, scanner.Err()
}
