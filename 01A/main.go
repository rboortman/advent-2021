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

	increases := 0
	for i, d := range depths {
		if i == 0 {
			continue
		}
		if d > depths[i-1] {
			increases++
		}
	}

	fmt.Print(increases)
}
