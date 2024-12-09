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

	checkedNumbers := make(map[int]bool)
	differentNumbers := make(map[int]bool)

	for _, num := range blocks {
		if !differentNumbers[num] && num > 0 {
			differentNumbers[num] = false
		}
	}

	// Extract keys and sort them in descending order
	keys := make([]int, 0, len(differentNumbers))
	for key := range differentNumbers {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, num := range keys {
		// fmt.Println("Key: ", num)
		lastItem := blocks[len(blocks)-1]

		if lastItem == -1 {
			blocks = blocks[:len(blocks)-1]
			continue
		}

		lastItemBlockLength := 0
		firstPossibleBlockStartingIndex := -1
		continuousNullCount := 0
		encounteredFirstOccurence := false

		for j := len(blocks) - 1; j >= 0; j-- {
			if blocks[j] == num {
				lastItemBlockLength++
				if !encounteredFirstOccurence {
					encounteredFirstOccurence = true
				}
			} else if encounteredFirstOccurence {
				if continuousNullCount >= lastItemBlockLength {
					break
				}
			}
		}

		for j, v := range blocks {
			if v == -1 {
				if firstPossibleBlockStartingIndex == -1 {
					firstPossibleBlockStartingIndex = j
				}
				continuousNullCount++

				if continuousNullCount == lastItemBlockLength {
					break
				}
			} else {
				firstPossibleBlockStartingIndex = -1
				continuousNullCount = 0
			}
		}

		// fmt.Println(continuousNullCount, lastItemBlockLength, blocks)

		checkedNumbers[num] = true

		if firstPossibleBlockStartingIndex != -1 {
			for i := 0; i < lastItemBlockLength; i++ {
				blocks[firstPossibleBlockStartingIndex+i] = num
			}

			lastNumberstartingPos := -1

			for j := len(blocks) - 1; j >= 0; j-- {
				if blocks[j] == num {
					lastNumberstartingPos = j
					break
				}
			}

			for j := lastNumberstartingPos; j > lastNumberstartingPos-lastItemBlockLength; j-- {
				blocks[j] = -1
			}
		}

		for len(blocks) > 0 && blocks[len(blocks)-1] == -1 {
			blocks = blocks[:len(blocks)-1]
		}

		// fmt.Printf("First possible starting index: %v, number: %v, last block length: %v\n", firstPossibleBlockStartingIndex, num, lastItemBlockLength)
	}

	// fmt.Println(blocks)

	sum := 0

	for index, num := range blocks {
		if num != -1 {
			sum += index * num
		}
	}

	fmt.Println(sum)
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
