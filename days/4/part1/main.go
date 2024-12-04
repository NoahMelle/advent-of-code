package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	const filePath string = "./my_input_file.txt"

	grid, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		panic(err)
	}

	xmasCount := findXMAS(grid)
	fmt.Println(xmasCount)

	duration := time.Since(start)
	fmt.Printf("\nExecuted the program in %s", duration)
}

func findXMAS(grid [][]rune) int {
	const target = "XMAS"
	const targetLen = len(target)

	directions := []struct {
		dx, dy int
	}{
		{1, 0},   // rechts
		{-1, 0},  // links
		{0, 1},   // onder
		{0, -1},  // boven
		{1, 1},   // rechtsonder
		{-1, -1}, // linksboven
		{1, -1},  // rechtsboven
		{-1, 1},  // linksonder
	}

	count := 0
	for y, row := range grid {
		for x := range row {
			if row[x] == 'X' {
				for _, dir := range directions {
					// check out of bounds
					if x+(targetLen-1)*dir.dx < 0 || x+(targetLen-1)*dir.dx >= len(row) ||
						y+(targetLen-1)*dir.dy < 0 || y+(targetLen-1)*dir.dy >= len(grid) {
						continue
					}
					if checkDirection(grid, x, y, dir.dx, dir.dy, target) {
						count++
					}
				}
			}
		}
	}
	return count
}

func checkDirection(grid [][]rune, startX, startY, dx, dy int, target string) bool {
	for i := 1; i < len(target); i++ {
		newX := startX + i*dx
		newY := startY + i*dy
		if grid[newY][newX] != rune(target[i]) {
			return false
		}
	}
	return true
}

func parseFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid, scanner.Err()
}
