package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	code := [][]string{}
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		code = append(code, chars)
	}

	char_map := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}

	autocomplete := [][]string{}
	for _, line := range code {
		expected_stack := []string{}
		for _, char := range line {
			if val, ok := char_map[char]; ok {
				expected_stack = append(expected_stack, val)
			} else if expected_stack[len(expected_stack)-1] == char {
				expected_stack = expected_stack[:len(expected_stack)-1]
			} else {
				expected_stack = []string{}
				break
			}
		}
		if len(expected_stack) < 1 {
			continue
		}
		autocomplete = append(autocomplete, reverseArray(expected_stack))
	}

	points_map := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	all_points := []int{}
	for _, line := range autocomplete {
		points := 0
		for _, char := range line {
			points *= 5
			points += points_map[char]
		}
		all_points = append(all_points, points)
	}

	sort.Ints(all_points)

	fmt.Println(len(all_points))
	fmt.Println(all_points[len(all_points)/2])
}
