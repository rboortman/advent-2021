package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x     int
	y     int
	value int
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func is_low_point(field [][]Point, location Point) bool {
	is_lowest := true

	if is_lowest && location.x > 0 {
		is_lowest = field[location.y][location.x-1].value > location.value
	}
	if is_lowest && location.x < len(field[0])-1 {
		is_lowest = field[location.y][location.x+1].value > location.value
	}
	if is_lowest && location.y > 0 {
		is_lowest = field[location.y-1][location.x].value > location.value
	}
	if is_lowest && location.y < len(field)-1 {
		is_lowest = field[location.y+1][location.x].value > location.value
	}

	return is_lowest
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	field := [][]Point{}
	low_points := []Point{}
	for scanner.Scan() {
		points_strings := strings.Split(scanner.Text(), "")
		row := []Point{}
		y := len(field)
		for _, ch := range points_strings {
			d, err := strconv.Atoi(ch)
			if_err(err)
			row = append(row, Point{x: len(row), y: y, value: d})
		}
		field = append(field, row)
	}

	for _, row := range field {
		for _, loc := range row {
			if is_low_point(field, loc) {
				low_points = append(low_points, loc)
			}
		}
	}

	sum := len(low_points)
	for _, loc := range low_points {
		sum += loc.value
	}

	fmt.Println(sum)
}
