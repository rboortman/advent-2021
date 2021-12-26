package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func get_basin(field [][]Point, location Point) ([][]Point, []Point) {
	basin := []Point{location}
	new_field := field
	new_field[location.y][location.x].value = -1

	if location.x > 0 && new_field[location.y][location.x-1].value < 9 && new_field[location.y][location.x-1].value >= 0 {
		updated_field, to_add := get_basin(new_field, new_field[location.y][location.x-1])
		new_field = updated_field
		basin = append(basin, to_add...)
	}
	if location.x < len(new_field[0])-1 && new_field[location.y][location.x+1].value < 9 && new_field[location.y][location.x+1].value >= 0 {
		updated_field, to_add := get_basin(new_field, new_field[location.y][location.x+1])
		new_field = updated_field
		basin = append(basin, to_add...)
	}
	if location.y > 0 && new_field[location.y-1][location.x].value < 9 && new_field[location.y-1][location.x].value >= 0 {
		updated_field, to_add := get_basin(new_field, new_field[location.y-1][location.x])
		new_field = updated_field
		basin = append(basin, to_add...)
	}
	if location.y < len(new_field)-1 && new_field[location.y+1][location.x].value < 9 && new_field[location.y+1][location.x].value >= 0 {
		updated_field, to_add := get_basin(new_field, new_field[location.y+1][location.x])
		new_field = updated_field
		basin = append(basin, to_add...)
	}

	return new_field, basin
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

	basins := [][]Point{}
	for _, loc := range low_points {
		if loc.value >= 0 {
			updated_field, basin := get_basin(field, loc)
			field = updated_field
			basins = append(basins, basin)
		}
	}

	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	fmt.Println(len(basins[0]) * len(basins[1]) * len(basins[2]))
}
