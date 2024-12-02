package day01

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Main(path string) {
	p1(path)
	p2(path)
}

func p1(path string) float64 {
	var left []int
	var right []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		line := strings.Split(temp, "   ")

		t, e := strconv.Atoi(line[0])
		if e == nil {
			left = append(left, t)
		}

		t, e = strconv.Atoi(line[1])
		if e == nil {
			right = append(right, t)
		}
	}
	sort.Ints(left)
	sort.Ints(right)

	var sum float64 = 0

	for i := 0; i < len(left); i++ {
		j := float64(left[i] - right[i])
		sum += math.Abs(j)
	}
	return sum
}

func p2(path string) int {
	var left []int
	var right []int

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp := scanner.Text()
		line := strings.Split(temp, "   ")

		t, e := strconv.Atoi(line[0])
		if e == nil {
			left = append(left, t)
		}

		t, e = strconv.Atoi(line[1])
		if e == nil {
			right = append(right, t)
		}
	}

	counts := make(map[int]int)

	for _, num := range right {
		counts[num]++
	}

	sum := 0

	for i := 0; i < len(left); i++ {
		sum += left[i] * counts[left[i]]
	}

	return sum
}
