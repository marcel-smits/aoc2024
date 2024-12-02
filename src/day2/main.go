package day02

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func p1(path string) int {
	res := 0
	input, _ := os.ReadFile(path)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		reports := []int{}
		for _, s := range strings.Fields(s) {
			n, _ := strconv.Atoi(s)
			reports = append(reports, n)
		}

		if applyRule(reports) {
			res++
		}
	}

	return res
}

func p2(path string) int {
	res := 0
	input, _ := os.ReadFile(path)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		reports := []int{}
		for _, s := range strings.Fields(s) {
			n, _ := strconv.Atoi(s)
			reports = append(reports, n)
		}

		for i := range reports {
			if applyRule(append(slices.Clone(reports[:i]), reports[i+1:]...)) {
				res++
				break
			}
		}
	}

	return res
}

func applyRule(reports []int) bool {
	for i := 1; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]

		if getSign(diff) != getSign(reports[1]-reports[0]) || abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func getSign(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}
