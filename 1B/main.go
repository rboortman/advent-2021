package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := []int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		depths = append(depths, i)
	}

	sliding_window := []int{}
	for i, d := range depths {
		if i < 2 {
			continue
		}
		sliding_window = append(sliding_window, d+depths[i-1]+depths[i-2])
	}

	increases := 0
	for i, d := range sliding_window {
		if i == 0 {
			continue
		}
		if d > sliding_window[i-1] {
			increases++
		}
	}

	fmt.Print(increases)
}
