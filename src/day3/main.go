package day03

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func p1(path string) int {
	res := 0
	input, _ := os.ReadFile(path)
	regexExp := `mul\((\d{1,3}),(\d{1,3})\)`
	regex := regexp.MustCompile(regexExp)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		matches := regex.FindAllStringSubmatch(s, -1)

		for _, match := range matches {

			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			res += x * y
		}
	}

	return res
}

func p2(path string) int {
	res := 0
	enabled := true
	input, _ := os.ReadFile(path)
	re := regexp.MustCompile(`(don't|do|\b\w*mul\((\d+),(\d+)\))`)
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\r\n") {
		matches := re.FindAllStringSubmatch(s, -1)

		for _, match := range matches {
			if match[1] != "don't" && match[1] != "do" {
				if enabled {
					x, _ := strconv.Atoi(match[2])
					y, _ := strconv.Atoi(match[3])
					res += x * y
					fmt.Printf("mul(%d, %d) = %d\n", x, y, x*y)
				} else {
					fmt.Printf("Skipping mul(%s, %s) due to disabled state\n", match[2], match[3])
				}
			}

			switch match[1] {
			case "don't":
				enabled = false
				fmt.Println("Multiplications disabled")
			case "do":
				enabled = true
				fmt.Println("Multiplications enabled")
			}
		}
	}

	return res
}
