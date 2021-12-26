package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digits := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	size := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		size++
		split_digits := strings.Split(scanner.Text(), "")
		for i, ch := range split_digits {
			d, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			digits[i] += d
		}
	}

	gamma := ""
	epsilon := ""

	for _, d := range digits {
		if d > (size / 2) {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gamma_val, err := strconv.ParseInt(gamma, 2, 64)
	epsilon_val, err := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gamma_val * epsilon_val)
	// fmt.Printf("horizontal: %d\nvertical: %d\ntogether: %d\n", horizontal, vertical, horizontal*vertical)
}
