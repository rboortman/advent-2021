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

func age_octopus(grid [][]int) [][]int {
	new_grid := [][]int{}

	for _, row := range grid {
		new_row := []int{}
		for _, octopus := range row {
			new_row = append(new_row, (octopus+1)%10)
		}
		new_grid = append(new_grid, new_row)
	}

	return new_grid
}

func clean_and_count(grid [][]int) ([][]int, int) {
	new_grid := [][]int{}
	count := 0

	for _, row := range grid {
		new_row := []int{}
		for _, octopus := range row {
			if octopus < 0 {
				new_row = append(new_row, 0)
			} else {
				new_row = append(new_row, octopus)
			}
			if octopus < 1 {
				count++
			}
		}
		new_grid = append(new_grid, new_row)
	}

	return new_grid, count
}

func new_value(val int) int {
	to_return := (val + 1) % 10
	// if to_return == 0 {
	// 	to_return = -1
	// }
	return to_return
}

func cascade_flash(grid [][]int) ([][]int, int) {
	needs_cascading := false
	for _, row := range grid {
		for _, octopus := range row {
			needs_cascading = needs_cascading || octopus == 0
		}
	}
	if !needs_cascading {
		return clean_and_count(grid)
	}

	new_grid := grid

	for i, row := range new_grid {
		for j, octopus := range row {
			if octopus == 0 {
				v_length := len(new_grid) - 1
				h_length := len(row) - 1

				if i > 0 && j > 0 && new_grid[i-1][j-1] > 0 {
					new_grid[i-1][j-1] = new_value(new_grid[i-1][j-1])
				}
				if i > 0 && new_grid[i-1][j] > 0 {
					new_grid[i-1][j] = new_value(new_grid[i-1][j])
				}
				if i > 0 && j < h_length && new_grid[i-1][j+1] > 0 {
					new_grid[i-1][j+1] = new_value(new_grid[i-1][j+1])
				}
				if j < h_length && new_grid[i][j+1] > 0 {
					new_grid[i][j+1] = new_value(new_grid[i][j+1])
				}
				if i < v_length && j < h_length && new_grid[i+1][j+1] > 0 {
					new_grid[i+1][j+1] = new_value(new_grid[i+1][j+1])
				}
				if i < v_length && new_grid[i+1][j] > 0 {
					new_grid[i+1][j] = new_value(new_grid[i+1][j])
				}
				if i < v_length && j > 0 && new_grid[i+1][j-1] > 0 {
					new_grid[i+1][j-1] = new_value(new_grid[i+1][j-1])
				}
				if j > 0 && new_grid[i][j-1] > 0 {
					new_grid[i][j-1] = new_value(new_grid[i][j-1])
				}

				new_grid[i][j] = -1
			}
		}
	}

	return cascade_flash(new_grid)
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		grid_row := []int{}
		for _, s := range line {
			d, err := strconv.Atoi(s)
			if_err(err)
			grid_row = append(grid_row, d)
		}
		grid = append(grid, grid_row)
	}

	steps := 100
	counter := 0

	for i := 0; i < steps; i++ {
		new_grid, step_counter := cascade_flash(age_octopus(grid))
		grid = new_grid
		counter += step_counter
	}

	fmt.Println(grid, counter)
}
