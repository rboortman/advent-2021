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
	horizontal := 0
	vertical := 0
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		res := strings.Fields(scanner.Text())
		i, err := strconv.Atoi(res[1])
		if err != nil {
			panic(err)
		}

		switch res[0] {
		case "up":
			vertical -= i
		case "down":
			vertical += i
		case "forward":
			horizontal += i
		}
	}

	fmt.Printf("horizontal: %d\nvertical: %d\ntogether: %d\n", horizontal, vertical, horizontal*vertical)
}
