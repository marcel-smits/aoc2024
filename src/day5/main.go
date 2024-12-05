package day05

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func p1(path string) int {
	res := 0

	graph := map[int][]int{}
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		src, dst := parseRule(text)
		graph[src] = append(graph[src], dst)
	}

	updates := [][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		updates = append(updates, parseUpdates(text))
	}

	for _, update := range updates {
		path := map[int]bool{}

	CHECK:
		for _, page := range update {
			for _, neighbor := range graph[page] {
				if _, ok := path[neighbor]; ok {
					break CHECK
				}
			}
			path[page] = true
		}

		if len(path) == len(update) {
			res += update[len(update)/2]
		}
	}

	return res
}

func p2(path string) int {
	res := 0

	graph := map[int][]int{}
	input, _ := os.Open(path)
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		src, dst := parseRule(text)
		graph[src] = append(graph[src], dst)
	}

	updates := [][]int{}
	for scanner.Scan() {
		text := scanner.Text()
		updates = append(updates, parseUpdates(text))
	}

	for _, update := range updates {
		path := map[int]int{}
		fixed := false

		for i := 0; i < len(update); i++ {
			page := update[i]
			path[page] = i
			for _, next := range graph[page] {
				if j, ok := path[next]; ok && path[next] < path[page] {
					fixed = true
					update[i], update[j] = update[j], update[i]
					path[next], path[page] = path[page], path[next]
					i = path[page]
				}
			}
		}

		if fixed {
			res += update[len(update)/2]
		}
	}

	return res
}

func parseRule(rule string) (int, int) {
	parts := strings.Split(rule, "|")
	
	p1, _ := strconv.Atoi(parts[0])
	p2, _ := strconv.Atoi(parts[1])

	return p1,p2
}

func parseUpdates(update string) []int  {
	res := []int{}

	updates := strings.Split(update, ",")

	for _, segment := range updates {
		num, _ := strconv.Atoi(segment)
		res = append(res, num)
	}

	return res
}
