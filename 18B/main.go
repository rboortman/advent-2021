package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	left        *Pair
	right       *Pair
	left_value  int
	right_value int
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func parse_calculation(line []string) ([]string, Pair) {
	line_remain := line[1:]
	p := Pair{}

	first_char := line_remain[0]

	if first_char == "[" {
		new_line, left_pair := parse_calculation(line_remain)
		line_remain = new_line
		p.left = &left_pair
	} else {
		d, err := strconv.Atoi(first_char)
		if_err(err)
		p.left_value = d
		line_remain = line_remain[1:]
	}

	second_char := line_remain[0]
	if second_char != "," {
		fmt.Println("WHELP!")
	}
	line_remain = line_remain[1:]

	third_char := line_remain[0]
	if third_char == "[" {
		new_line, right_pair := parse_calculation(line_remain)
		line_remain = new_line
		p.right = &right_pair
	} else {
		d, err := strconv.Atoi(third_char)
		if_err(err)
		p.right_value = d
		line_remain = line_remain[1:]
	}

	line_remain = line_remain[1:]

	return line_remain, p
}

func add_to_pair(p Pair, to_add int, left_add bool) Pair {
	new_pair := p

	if left_add {
		if new_pair.left == nil {
			new_pair.left_value += to_add
		} else {
			new_left := add_to_pair(*new_pair.left, to_add, left_add)
			new_pair.left = &new_left
		}
	} else {
		if new_pair.right == nil {
			new_pair.right_value += to_add
		} else {
			new_right := add_to_pair(*new_pair.right, to_add, left_add)
			new_pair.right = &new_right
		}
	}

	return new_pair
}

func explode(p Pair, level int) (Pair, int, int) {
	new_pair := p
	left_added := 0
	right_added := 0

	if level == 4 {
		if new_pair.left != nil {
			left_added = new_pair.left.left_value
			if new_pair.right != nil {
				new_pair.right.left_value += new_pair.left.right_value
			} else {
				new_pair.right_value += new_pair.left.right_value
			}
			new_pair.left = nil
		}

		if new_pair.right != nil {
			new_pair.left_value += new_pair.right.left_value
			right_added = new_pair.right.right_value
			new_pair.right = nil
		}
	} else {
		if new_pair.left != nil {
			new_left, l, r := explode(*new_pair.left, level+1)
			new_pair.left = &new_left
			left_added = l

			if new_pair.right != nil {
				new_right := add_to_pair(*new_pair.right, r, true)
				new_pair.right = &new_right
			} else {
				new_pair.right_value += r
			}
		}

		if new_pair.right != nil {
			new_right, l, r := explode(*new_pair.right, level+1)
			new_pair.right = &new_right
			right_added = r

			if new_pair.left != nil {
				new_left := add_to_pair(*new_pair.left, l, false)
				new_pair.left = &new_left
			} else {
				new_pair.left_value += l
			}
		}
	}

	return new_pair, left_added, right_added
}

func split(p Pair, level int) (Pair, bool) {
	new_pair := p
	needs_exploding := false

	if new_pair.left != nil {
		new_left, e := split(*new_pair.left, level+1)
		new_pair.left = &new_left
		needs_exploding = needs_exploding || e
	} else {
		if new_pair.left_value > 9 {
			needs_exploding = needs_exploding || level >= 4
			new_pair.left = &Pair{
				left_value:  new_pair.left_value / 2,
				right_value: new_pair.left_value/2 + new_pair.left_value%2,
			}
			new_pair.left_value = 0

			if !needs_exploding {
				new_left, e := split(*new_pair.left, level+1)
				new_pair.left = &new_left
				needs_exploding = needs_exploding || e
			}
		}
	}

	if needs_exploding {
		return new_pair, needs_exploding
	}

	if new_pair.right != nil {
		new_right, e := split(*new_pair.right, level+1)
		new_pair.right = &new_right
		needs_exploding = needs_exploding || e

	} else {
		if new_pair.right_value > 9 {
			needs_exploding = needs_exploding || level >= 4
			new_pair.right = &Pair{
				left_value:  new_pair.right_value / 2,
				right_value: new_pair.right_value/2 + new_pair.right_value%2,
			}
			new_pair.right_value = 0

			if !needs_exploding {
				new_right, e := split(*new_pair.right, level+1)
				new_pair.right = &new_right
				needs_exploding = needs_exploding || e
			}
		}
	}

	return new_pair, needs_exploding
}

func snail_math(p1, p2 Pair) Pair {
	result := Pair{
		left:  &p1,
		right: &p2,
	}

	// fmt.Println("\nto calculate:", print_pair(result))
	needs_recalculate := false

	for ok := true; ok; ok = needs_recalculate {
		after_explode, _, _ := explode(result, 1)
		after_split, e := split(after_explode, 1)
		needs_recalculate = e
		// fmt.Printf("  intermediate: recalculate = %t, pair: %s\n", needs_recalculate, print_pair(after_split))
		result = after_split
	}

	return result
}

func print_pair(p Pair) string {
	res := "["

	if p.left != nil {
		res += print_pair(*p.left)
	} else {
		res += strconv.Itoa(p.left_value)
	}

	res += ","

	if p.right != nil {
		res += print_pair(*p.right)
	} else {
		res += strconv.Itoa(p.right_value)
	}

	res += "]"

	return res
}

func calc_math(p Pair) int {
	left_value, right_value := p.left_value, p.right_value
	if p.left != nil {
		left_value = calc_math(*p.left)
	}
	if p.right != nil {
		right_value = calc_math(*p.right)
	}
	return 3*left_value + 2*right_value
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	max_result := 0
	for i, l1 := range lines {
		for j, l2 := range lines {
			if i == j {
				continue
			}
			_, p1 := parse_calculation(strings.Split(l1, ""))
			_, p2 := parse_calculation(strings.Split(l2, ""))

			calculated := snail_math(p1, p2)
			result := calc_math(calculated)
			// fmt.Printf("Calculating.. %d + %d = %d\n", i, j, result)

			if max_result < result {
				max_result = result
			}
		}
	}

	fmt.Println(max_result)
}
