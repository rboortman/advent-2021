package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func if_err(err error) {
	if err != nil {
		panic(err)
	}
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

	wrong_chars := []string{}
	for _, line := range code {
		expected_stack := []string{}
		for _, char := range line {
			if val, ok := char_map[char]; ok {
				expected_stack = append(expected_stack, val)
			} else if expected_stack[len(expected_stack)-1] == char {
				expected_stack = expected_stack[:len(expected_stack)-1]
			} else {
				wrong_chars = append(wrong_chars, char)
				break
			}
		}
	}

	points_map := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	points := 0
	for _, char := range wrong_chars {
		points += points_map[char]
	}

	fmt.Println(points)
}
