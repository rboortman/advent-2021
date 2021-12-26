package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func fold_dots(dots [][]string, fold_line Fold) [][]string {
	fmt.Println(len(dots), fold_line)
	new_dots := [][]string{}
	if fold_line.direction == "x" {
		for _, line := range dots {
			new_line := []string{}
			if len(dots)/2 < fold_line.location {
				for i, c := range line {
					if i < fold_line.location {
						new_line = append(new_line, c)
					} else if i > fold_line.location && c == "#" {
						new_line[len(new_line)-(i-fold_line.location)] = c
					}
				}
			} else {
				for i := len(line) - 1; i >= 0; i-- {
					if i > fold_line.location {
						new_line = append(new_line, line[i])
					} else if i < fold_line.location && line[i] == "#" {
						new_line[len(new_line)-(fold_line.location-i)+1] = "#"
					}
				}
			}
			new_dots = append(new_dots, new_line)
		}
	} else if fold_line.direction == "y" {
		if len(dots)/2 <= fold_line.location {
			for i, line := range dots {
				if i < fold_line.location {
					new_dots = append(new_dots, line)
				} else if i > fold_line.location {
					for j, c := range line {
						if c == "#" {
							new_dots[len(new_dots)-(i-fold_line.location)][j] = "#"
						}
					}
				}
			}
		} else {
			for i := len(dots) - 1; i >= 0; i-- {
				if i > fold_line.location {
					new_dots = append(new_dots, dots[i])
				} else if i < fold_line.location {
					for j, c := range dots[i] {
						if c == "#" {
							new_dots[len(new_dots)-(fold_line.location-i)][j] = "#"
						}
					}
				}
			}
		}
	}
	return new_dots
}

type Fold struct {
	direction string
	location  int
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coords := [][]int{}
	folds := []Fold{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if strings.Contains(line, ",") {
			coord_string := strings.Split(line, ",")
			x, err := strconv.Atoi(coord_string[0])
			if_err(err)
			y, err := strconv.Atoi(coord_string[1])
			if_err(err)
			coords = append(coords, []int{x, y})
		} else {
			s := strings.Fields(line)
			fold_string := strings.Split(s[2], "=")
			loc, err := strconv.Atoi(fold_string[1])
			if_err(err)
			folds = append(folds, Fold{direction: fold_string[0], location: loc})
		}
	}

	largest_x, largest_y := 0, 0
	for _, coord := range coords {
		if coord[0] > largest_x {
			largest_x = coord[0]
		}
		if coord[1] > largest_y {
			largest_y = coord[1]
		}
	}
	largest_x++
	largest_y++

	dots := [][]string{}
	for i := 0; i < largest_y; i++ {
		row := []string{}
		for j := 0; j < largest_x; j++ {
			row = append(row, ".")
		}
		dots = append(dots, row)
	}

	for _, coord := range coords {
		dots[coord[1]][coord[0]] = "#"
	}

	for _, f := range folds {
		dots = fold_dots(dots, f)
	}

	fmt.Println(dots)
}
