package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	const filePath string = "./my_input_file.txt"

	lines, err := parseFile(filePath)

	if err != nil {
		fmt.Printf("error reading file: %v", err)
		panic(err)
	}

	blocks := []int{}
	blocksFilesLength := 0
	indexCount := 0

	for index, num := range lines {
		if index%2 == 0 {
			for i := 0; i < num; i++ {
				blocks = append(blocks, indexCount)
				blocksFilesLength++
			}
			indexCount++
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, -1)
			}
		}
	}

	differentNumbers := make(map[int]bool)

	for _, num := range blocks {
		if !differentNumbers[num] && num > 0 {
			differentNumbers[num] = false
		}
	}

	keys := make([]int, 0, len(differentNumbers))
	for key := range differentNumbers {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, num := range keys {
		lastItemBlockLength := getLastItemBlockLength(blocks, num)
		firstPossibleBlockStartingIndex := getFirstPossibleStartingIndex(blocks, lastItemBlockLength)

		if firstPossibleBlockStartingIndex != -1 {
			for i := 0; i < lastItemBlockLength; i++ {
				blocks[firstPossibleBlockStartingIndex+i] = num
			}

			lastNumberStartingPos := -1

			for j := len(blocks) - 1; j >= 0; j-- {
				if blocks[j] == num {
					lastNumberStartingPos = j
					break
				}
			}

			if lastNumberStartingPos != len(blocks)-1 {
				for j := lastNumberStartingPos; j > lastNumberStartingPos-lastItemBlockLength; j-- {
					blocks[j] = -1
				}
			} else {
				blocks = blocks[:len(blocks)-lastItemBlockLength]
			}
		}
	}

	sum := 0

	for index, num := range blocks {
		if num != -1 {
			sum += index * num
		}
	}

	fmt.Println(sum)
}

func getLastItemBlockLength(arr []int, num int) int {
	lastItemBlockLength := 0
	encounteredFirstOccurence := false

	for j := len(arr) - 1; j >= 0; j-- {
		if arr[j] == num {
			lastItemBlockLength++
			if !encounteredFirstOccurence {
				encounteredFirstOccurence = true
			}
		} else if encounteredFirstOccurence {
			return lastItemBlockLength
		}
	}

	return lastItemBlockLength
}

func getFirstPossibleStartingIndex(arr []int, targetLen int) int {
	firstPossibleStartingIndex := -1
	continuousNullCount := 0

	for j, v := range arr {
		if v == -1 {
			if firstPossibleStartingIndex == -1 {
				firstPossibleStartingIndex = j
			}
			continuousNullCount++

			if continuousNullCount == targetLen {
				break
			}
		} else {
			firstPossibleStartingIndex = -1
			continuousNullCount = 0
		}
	}

	return firstPossibleStartingIndex
}

func parseFile(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer file.Close()

	var content []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		char := scanner.Text()
		num, err := strconv.Atoi(char)
		if err != nil {
			return nil, err
		}
		content = append(content, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return content, nil
}
