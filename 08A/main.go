package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Display struct {
	input  [][]string
	output [][]string
}

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
	displays := []Display{}
	for scanner.Scan() {
		input_output := strings.Split(scanner.Text(), " | ")
		input := [][]string{}
		output := [][]string{}

		for _, d := range strings.Split(input_output[0], ",") {
			for _, in := range strings.Fields(d) {
				input = append(input, strings.Split(in, ""))
			}
		}
		for _, d := range strings.Split(input_output[1], ",") {
			for _, in := range strings.Fields(d) {
				output = append(output, strings.Split(in, ""))
			}
		}

		displays = append(displays, Display{input: input, output: output})
	}

	counter := 0
	for _, display := range displays {
		for _, output := range display.output {
			switch len(output) {
			case 2:
				counter++
			case 3:
				counter++
			case 4:
				counter++
			case 7:
				counter++
			}
		}
	}

	fmt.Println(counter)
}
