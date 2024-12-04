package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	grid, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		panic(err)
	}

	xmasCount := findXMAS(grid)
	fmt.Println(xmasCount)
}

func findXMAS(grid [][]string) int {
	const target = "XMAS"
	directions := []struct {
		dx, dy int
	}{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	count := 0
	for y, row := range grid {
		for x := range row {
			for _, dir := range directions {
				if checkDirection(grid, x, y, dir.dx, dir.dy, target) {
					count++
				}
			}
		}
	}
	return count
}

func checkDirection(grid [][]string, startX, startY, dx, dy int, target string) bool {
	for i := 0; i < len(target); i++ {
		newX := startX + i*dx
		newY := startY + i*dy
		if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[newY]) || grid[newY][newX] != string(target[i]) {
			return false
		}
	}
	return true
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
