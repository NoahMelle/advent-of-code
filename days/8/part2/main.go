package main

import (
	"bufio"
	"fmt"
	"os"
)

type Position struct {
	posX, posY int
}

func main() {
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)
	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}

	antennas := make(map[rune][]Position)
	antinodes := make(map[string]bool)

	for yIndex, line := range lines {
		for xIndex, char := range line {
			if char != '.' {
				antennas[char] = append(antennas[char], Position{xIndex, yIndex})
			}
		}
	}

	for _, positions := range antennas {
		combinations := getAllAntennaCombinations(positions)
		for _, combination := range combinations {
			xDiff := combination[0].posX - combination[1].posX
			yDiff := combination[0].posY - combination[1].posY

			tempAntennaPos := Position{combination[1].posX, combination[1].posY}
			lastValidPos := Position{combination[1].posX, combination[1].posY}

			for isValidPosition(tempAntennaPos, len(lines[0]), len(lines)) {
				antinodes[fmt.Sprintf("%d, %d", tempAntennaPos.posX, tempAntennaPos.posY)] = true

				tempAntennaPos = Position{tempAntennaPos.posX + xDiff, tempAntennaPos.posY + yDiff}

				if isValidPosition(tempAntennaPos, len(lines[0]), len(lines)) {
					lastValidPos = tempAntennaPos
				}
			}

			tempAntennaPos = lastValidPos

			for isValidPosition(tempAntennaPos, len(lines[0]), len(lines)) {
				antinodes[fmt.Sprintf("%d, %d", tempAntennaPos.posX, tempAntennaPos.posY)] = true

				tempAntennaPos = Position{tempAntennaPos.posX - xDiff, tempAntennaPos.posY - yDiff}
			}
		}
	}

	fmt.Println(len(antinodes))
}

func getAllAntennaCombinations(positions []Position) [][2]Position {
	var combinations [][2]Position
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			combinations = append(combinations, [2]Position{positions[i], positions[j]})
		}
	}
	return combinations
}

func parseFile(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, []rune(scanner.Text()))
	}
	return content, scanner.Err()
}

func isValidPosition(pos Position, width, height int) bool {
	return pos.posX >= 0 && pos.posX < width && pos.posY >= 0 && pos.posY < height
}
