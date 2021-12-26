package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func most_common_digit(digits [][]int, pos int) int {
	count := 0.0
	for _, line := range digits {
		count += float64(line[pos])
	}

	fmt.Printf("most_common_digit\n  pos: %d\n  count: %f\n  half: %f\n-----------------\n", pos, count, float64(len(digits)/2))

	if count < float64(len(digits))/float64(2) {
		return 0
	} else {
		return 1
	}
}

func least_common_digit(digits [][]int, pos int) int {
	count := 0.0
	for _, line := range digits {
		count += float64(line[pos])
	}

	fmt.Printf("least_common_digit\n  pos: %d\n  count: %f\n  half: %f\n-----------------\n", pos, count, float64(len(digits)/2))

	if count < float64(len(digits))/float64(2) {
		return 1
	} else {
		return 0
	}
}

func filter_digits(digits [][]int, pos int, val int) [][]int {
	new_digits := [][]int{}

	for _, line := range digits {
		if line[pos] == val {
			new_digits = append(new_digits, line)
		}
	}

	return new_digits
}

func rec_find(digits [][]int, evaluator func(digits [][]int, pos int) int, pos int) []int {
	if len(digits) <= 1 {
		return digits[0]
	}

	d := evaluator(digits, pos)
	new_digits := filter_digits(digits, pos, d)
	return rec_find(new_digits, evaluator, (pos + 1))
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digits := [][]int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		split_digits := strings.Split(scanner.Text(), "")
		toAppend := [12]int{}
		for i, ch := range split_digits {
			d, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			toAppend[i] = d
		}
		digits = append(digits, toAppend[:])
	}

	oxy := rec_find(digits, most_common_digit, 0)
	co2 := rec_find(digits, least_common_digit, 0)

	oxy_val, err := strconv.ParseInt(strings.Trim(strings.Replace(fmt.Sprint(oxy), " ", "", -1), "[]"), 2, 64)
	if err != nil {
		panic(err)
	}
	co2_val, err := strconv.ParseInt(strings.Trim(strings.Replace(fmt.Sprint(co2), " ", "", -1), "[]"), 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(oxy)
	fmt.Println(co2)
	fmt.Println(oxy_val)
	fmt.Println(co2_val)
	fmt.Println(oxy_val * co2_val)
	// fmt.Printf("horizontal: %d\nvertical: %d\ntogether: %d\n", horizontal, vertical, horizontal*vertical)
}
