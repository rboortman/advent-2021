package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type Display struct {
	input  [][]string
	output []string
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func get_mapping(input [][]string) [10]string {
	digit_map := make(map[string]string)
	top_bar := difference(input[1], input[0])[0]

	digit_map["a"] = top_bar

	diff20 := difference(input[2], input[0])
	diff32a := difference(input[3], append(input[2], top_bar))
	diff42a := difference(input[4], append(input[2], top_bar))
	diff52a := difference(input[5], append(input[2], top_bar))

	if len(diff32a) == 1 {
		digit_map["g"] = diff32a[0]

		diff03 := difference(input[0], input[3])
		diff23 := difference(input[2], input[3])
		if len(diff03) == 1 {
			digit_map["c"] = diff03[0]
			digit_map["f"] = difference(input[0], diff03)[0]

			if len(diff42a) == 1 {
				diff24 := difference(input[2], input[4])
				digit_map["b"] = diff24[0]
				digit_map["d"] = difference(diff20, diff24)[0]
				digit_map["e"] = difference(diff52a, diff32a)[0]
			} else {
				diff25 := difference(input[2], input[5])
				digit_map["b"] = diff25[0]
				digit_map["d"] = difference(diff20, diff25)[0]
				digit_map["e"] = difference(diff42a, diff32a)[0]
			}
		} else {
			digit_map["b"] = diff23[0]
			digit_map["d"] = difference(diff20, diff23)[0]

			if len(diff42a) == 1 {
				diff04 := difference(input[0], input[4])
				digit_map["c"] = diff04[0]
				digit_map["f"] = difference(input[0], diff04)[0]
				digit_map["e"] = difference(diff52a, diff32a)[0]
			} else {
				diff05 := difference(input[0], input[5])
				digit_map["c"] = diff05[0]
				digit_map["f"] = difference(input[0], diff05)[0]
				digit_map["e"] = difference(diff42a, diff32a)[0]
			}
		}
	} else {
		digit_map["g"] = diff42a[0]

		diff04 := difference(input[0], input[4])
		diff24 := difference(input[2], input[4])
		if len(diff04) == 1 {
			digit_map["c"] = diff04[0]
			digit_map["f"] = difference(input[0], diff04)[0]

			diff25 := difference(input[2], input[5])
			digit_map["b"] = diff25[0]
			digit_map["d"] = difference(diff20, diff25)[0]
			digit_map["e"] = difference(diff32a, diff42a)[0]
		} else {
			digit_map["b"] = diff24[0]
			digit_map["d"] = difference(diff20, diff24)[0]

			diff05 := difference(input[0], input[5])
			digit_map["c"] = diff05[0]
			digit_map["f"] = difference(input[0], diff05)[0]
			digit_map["e"] = difference(diff32a, diff42a)[0]
		}
	}

	d0 := []string{digit_map["a"], digit_map["b"], digit_map["c"], digit_map["e"], digit_map["f"], digit_map["g"]}
	d1 := []string{digit_map["c"], digit_map["f"]}
	d2 := []string{digit_map["a"], digit_map["c"], digit_map["d"], digit_map["e"], digit_map["g"]}
	d3 := []string{digit_map["a"], digit_map["c"], digit_map["d"], digit_map["f"], digit_map["g"]}
	d4 := []string{digit_map["b"], digit_map["c"], digit_map["d"], digit_map["f"]}
	d5 := []string{digit_map["a"], digit_map["b"], digit_map["d"], digit_map["f"], digit_map["g"]}
	d6 := []string{digit_map["a"], digit_map["b"], digit_map["d"], digit_map["e"], digit_map["f"], digit_map["g"]}
	d7 := []string{digit_map["a"], digit_map["c"], digit_map["f"]}
	d8 := []string{digit_map["a"], digit_map["b"], digit_map["c"], digit_map["d"], digit_map["e"], digit_map["f"], digit_map["g"]}
	d9 := []string{digit_map["a"], digit_map["b"], digit_map["c"], digit_map["d"], digit_map["f"], digit_map["g"]}

	sort.Strings(d0)
	sort.Strings(d1)
	sort.Strings(d2)
	sort.Strings(d3)
	sort.Strings(d4)
	sort.Strings(d5)
	sort.Strings(d6)
	sort.Strings(d7)
	sort.Strings(d8)
	sort.Strings(d9)

	mapping_array := [10]string{strings.Join(d0, ""), strings.Join(d1, ""), strings.Join(d2, ""), strings.Join(d3, ""), strings.Join(d4, ""), strings.Join(d5, ""), strings.Join(d6, ""), strings.Join(d7, ""), strings.Join(d8, ""), strings.Join(d9, "")}

	return mapping_array
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
		output := []string{}

		for _, d := range strings.Split(input_output[0], ",") {
			for _, in := range strings.Fields(d) {
				to_append := strings.Split(in, "")
				sort.Strings(to_append)
				input = append(input, to_append)
			}
		}
		for _, d := range strings.Split(input_output[1], ",") {
			for _, in := range strings.Fields(d) {
				to_append := strings.Split(in, "")
				sort.Strings(to_append)
				output = append(output, strings.Join(to_append, ""))
			}
		}

		sort.SliceStable(input, func(i, j int) bool {
			return len(input[i]) < len(input[j])
		})
		displays = append(displays, Display{input: input, output: output})
	}

	counter := 0
	for _, display := range displays {
		mapping := get_mapping(display.input)
		display_digits := 0
		for i, d_string := range display.output {
			for j, dd_string := range mapping {
				if d_string == dd_string {
					display_digits += j * int(math.Pow10(3-i))
				}
			}
		}
		counter += display_digits
	}

	fmt.Println(counter)
}
