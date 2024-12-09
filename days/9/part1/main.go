package main

import (
	"bufio"
	"fmt"
	"os"
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

	for len(blocks) > blocksFilesLength {
		lastItem := blocks[len(blocks)-1]

		if lastItem == -1 {
			blocks = blocks[:len(blocks)-1]
			continue
		}

		lastItemBlockLength := 0

		for i := len(blocks) - 1; i >= 0 && blocks[i] == lastItem; i-- {
			if blocks[i] == lastItem {
				lastItemBlockLength++
			}
		}

		firstNullIndex := -1

		for i, v := range blocks {
			if v == -1 {
				firstNullIndex = i
				break
			}
		}

		blocks[firstNullIndex] = lastItem
		blocks = blocks[:len(blocks)-1]
	}
	sum := 0

	for index, num := range blocks {
		sum += index * num
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
