package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type line struct {
	from coord
	to   coord
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func convert_to_line(s string) line {
	coords := strings.Split(s, " -> ")
	from := strings.Split(coords[0], ",")
	to := strings.Split(coords[1], ",")

	from_x, err := strconv.Atoi(from[0])
	if_err(err)
	from_y, err := strconv.Atoi(from[1])
	if_err(err)
	to_x, err := strconv.Atoi(to[0])
	if_err(err)
	to_y, err := strconv.Atoi(to[1])
	if_err(err)
	return line{from: coord{x: from_x, y: from_y}, to: coord{x: to_x, y: to_y}}
}

func main() {
	file, err := os.Open("./input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []line{}
	for scanner.Scan() {
		l := convert_to_line(scanner.Text())
		lines = append(lines, l)
	}

	highest_x := -1
	highest_y := -1
	for _, l := range lines {
		if l.from.x > highest_x {
			highest_x = l.from.x
		}
		if l.to.x > highest_x {
			highest_x = l.to.x
		}
		if l.from.y > highest_y {
			highest_y = l.from.y
		}
		if l.to.y > highest_y {
			highest_y = l.to.y
		}
	}

	board := make([][]int, highest_y+1)
	for i := range board {
		board_line := make([]int, highest_x+1)
		for j := range board_line {
			board_line[j] = 0
		}
		board[i] = board_line
	}

	for _, l := range lines {
		if l.from.x == l.to.x {
			min_y := int(math.Min(float64(l.from.y), float64(l.to.y)))
			max_y := int(math.Max(float64(l.from.y), float64(l.to.y)))
			for i := min_y; i <= max_y; i++ {
				board[i][l.from.x]++
			}
		} else if l.from.y == l.to.y {
			min_x := int(math.Min(float64(l.from.x), float64(l.to.x)))
			max_x := int(math.Max(float64(l.from.x), float64(l.to.x)))
			for i := min_x; i <= max_x; i++ {
				board[l.from.y][i]++
			}
		} else {
			diff := int(math.Abs(float64(l.from.x) - float64(l.to.x)))
			sign_x := 1
			if l.from.x > l.to.x {
				sign_x = -1
			}
			sign_y := 1
			if l.from.y > l.to.y {
				sign_y = -1
			}
			for i := 0; i <= diff; i++ {
				board[l.from.y+(i*sign_y)][l.from.x+(i*sign_x)]++
			}
		}
	}

	counter := 0
	for _, ys := range board {
		for _, xs := range ys {
			if xs > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)
	// fmt.Printf("highest_x: %d\nhighest_y: %d\n", highest_x, highest_y)
}
