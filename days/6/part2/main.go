package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const filePath string = "./my_input_file.txt"

	lines, initialGuardX, initialGuardY, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}

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

	loopCount := 0

	for yIndex, line := range lines {
		for xIndex, place := range line {
			if place == "." {
				guardX := initialGuardX
				guardY := initialGuardY

				currentDir := 0

				gridWithObstacle := make([][]string, len(lines))
				for i := range lines {
					gridWithObstacle[i] = make([]string, len(lines[i]))
					copy(gridWithObstacle[i], lines[i])
				}

				gridWithObstacle[yIndex][xIndex] = "O"

				visitedPositions := make(map[[3]int]bool)
				for guardX >= 0 && guardX < totalX && guardY >= 0 && guardY < totalY {
					if visitedPositions[[3]int{guardX, guardY, currentDir}] {
						loopCount++
						break
					}

					visitedPositions[[3]int{guardX, guardY, currentDir}] = true

					currentDir, guardX, guardY = moveGuardWithObstacle(directions, currentDir, gridWithObstacle, guardX, guardY, totalX, totalY)
				}
			}
		}
	}

	fmt.Println("Exited for loop.")

	fmt.Printf("\nThe obstacles can be places at %v distinct places.", loopCount)
}

func moveGuardWithObstacle(directions []struct{ dx, dy int }, currentDir int, grid [][]string, posX int, posY int, totalX int, totalY int) (int, int, int) {
	newX := posX + directions[currentDir].dx
	newY := posY + directions[currentDir].dy

	if newX < 0 || newX >= totalX || newY < 0 || newY >= totalY {
		return currentDir, newX, newY
	}

	if grid[newY][newX] == "#" || grid[newY][newX] == "O" {
		currentDir = (currentDir + 1) % 4

		return currentDir, posX, posY
	}
	return currentDir, newX, newY
}

func parseFile(path string) ([][]string, int, int, error) {
	var guardX int
	var guardY int

	file, err := os.Open(path)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()
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
			}

			row = append(row, symbol)
		}

		grid = append(grid, row)

		currentRow++
	}

	return grid, guardX, guardY, scanner.Err()
}
