package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	lines, guardX, guardY, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}

	fmt.Println("\n GuardX: ", guardX)
	fmt.Println("\n GuardY: ", guardY)

	totalX := len(lines[0])
	totalY := len(lines)

	directions := []struct {
		dx, dy int
	}{
		{0, -1}, // boven
		{1, 0},  // rechts
		{0, 1},  // onder
		{-1, 0}, // links
	}

	currentDir := 0

	for guardX >= 0 && guardX < totalX && guardY >= 0 && guardY < totalY {
		currentDir, lines, guardX, guardY = moveGuard(directions, currentDir, lines, guardX, guardY, totalX, totalY)
	}

	distinctPlaces := 0

	for _, line := range lines {
		for _, place := range line {
			if place == "X" {
				distinctPlaces++
			}
		}
	}

	fmt.Printf("\nThe guard has visited %v places before leaving the area.", distinctPlaces)
}

func moveGuard(directions []struct{ dx, dy int }, currentDir int, grid [][]string, posX int, posY int, totalX int, totalY int) (int, [][]string, int, int) {
	newX := posX + directions[currentDir].dx
	newY := posY + directions[currentDir].dy

	if newX < 0 || newX >= totalX || newY < 0 || newY >= totalY {
		return currentDir, grid, newX, newY
	}

	if grid[newY][newX] == "#" {
		if currentDir < 3 {
			currentDir++
		} else {
			currentDir = 0
		}

		return currentDir, grid, posX, posY
	} else {
		grid[newY][newX] = "X"
		return currentDir, grid, newX, newY
	}
}

func parseFile(path string) ([][]string, int, int, error) {
	var guardX int
	var guardY int

	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	var possibleObstaclePositions int = 0
	var currentRow = 0
	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []string{}
		for i, symbol := range strings.Split(scanner.Text(), "") {
			if symbol == "^" {
				symbol = "X"
				guardX = i
				guardY = currentRow
			} else if symbol == "." {
				possibleObstaclePositions++
			}

			row = append(row, symbol)
		}

		grid = append(grid, row)

		currentRow++
	}

	fmt.Printf("There are %v obstacle positions", possibleObstaclePositions)

	return grid, guardX, guardY, scanner.Err()
}
