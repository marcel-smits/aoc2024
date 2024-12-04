package day04

import (
	"bufio"
	"os"
	"strings"
)

func p1(path string) int {
	res := 0
	input := parseFile(path)

	for rowId, row := range input {
		for colId := range row {
			if input[rowId][colId] == "X" {
				res += findMatch(input, "XMAS", rowId, colId)
			}
		}
	}

	return res
}

func p2(path string) int {
	res := 0
	input := parseFile(path)

	for rowId, row := range input {
		for colId := range row {
			if input[rowId][colId] == "A" && findMatch2(input, "MAS", rowId, colId) {
				res += 1
			}
		}
	}

	return res
}

func findMatch(input [][]string, term string, rowId, colId int) int {
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var res int

	for _, direction := range directions {
		currentRow, currentCol := rowId, colId
		var match bool
		for _, char := range term[1:] {
			currentRow += direction[0]
			currentCol += direction[1]

			if currentRow < 0 || currentRow >= len(input) || currentCol < 0 || currentCol >= len(input[0]) || input[currentRow][currentCol] != string(char) {
				match = false
				break
			}
			match = true
		}

		if match {
			res++
		}
	}

	return res
}

func findMatch2(input [][]string, word string, rowId, colId int) bool {
	directions := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	var found bool
	var wordCount int

	for _, dir := range directions {
		halfLength := len(word) / 2
		currRow := rowId - (dir[0] * halfLength)
		currCol := colId - (dir[1] * halfLength)
		for _, char := range word {
			if currRow < 0 || currRow >= len(input) || currCol < 0 || currCol >= len(input[0]) || input[currRow][currCol] != string(char) {
				found = false
				break
			}
			found = true
			currRow += dir[0]
			currCol += dir[1]
		}
		if found {
			wordCount += 1
		}
	}
	return wordCount == 2
}

func parseFile(path string) [][]string {
	file, _ := os.Open(path)

	defer file.Close()

	var parsed [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parsed = append(parsed, strings.Split(line, ""))
	}

	return parsed
}
