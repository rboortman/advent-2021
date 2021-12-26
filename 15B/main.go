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

func expand_grid(grid [][]Point) [][]Point {
	new_grid := [][]Point{}

	for i := 0; i < 5; i++ {
		for y, row := range grid {
			new_full_row := []Point{}
			for j := 0; j < 5; j++ {
				for x, p := range row {
					new_point := Point{x: j*len(grid[y]) + x, y: i*len(grid) + y, value: (p.value + i + j) % 9}
					if new_point.value == 0 {
						new_point.value = 9
					}
					new_full_row = append(new_full_row, new_point)
				}
			}
			new_grid = append(new_grid, new_full_row)
		}
	}

	return new_grid
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]Point{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		d_line := []Point{}
		i := len(grid)
		for j, val := range line {
			d, err := strconv.Atoi(val)
			if_err(err)
			d_line = append(d_line, Point{x: j, y: i, value: d})
		}
		grid = append(grid, d_line)
	}

	grid = expand_grid(grid)

	graph := newGraph()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			source := fmt.Sprintf("%d,%d", i, j)
			if i > 0 {
				dest := fmt.Sprintf("%d,%d", i-1, j)
				graph.addEdge(source, dest, grid[i-1][j].value)
			}
			if i < len(grid)-1 {
				dest := fmt.Sprintf("%d,%d", i+1, j)
				graph.addEdge(source, dest, grid[i+1][j].value)
			}
			if j > 0 {
				dest := fmt.Sprintf("%d,%d", i, j-1)
				graph.addEdge(source, dest, grid[i][j-1].value)
			}
			if j < len(grid[i])-1 {
				dest := fmt.Sprintf("%d,%d", i, j+1)
				graph.addEdge(source, dest, grid[i][j+1].value)
			}
		}
	}

	source := fmt.Sprintf("%d,%d", 0, 0)
	dest := fmt.Sprintf("%d,%d", len(grid)-1, len(grid[len(grid)-1])-1)
	size, _ := graph.getPath(source, dest)
	fmt.Println(size)
	// for _, p := range path {
	// 	fmt.Println(p)
	// }

	// Example
	// graph := newGraph()
	// graph.addEdge("S", "B", 4)
	// graph.addEdge("S", "C", 2)
	// graph.addEdge("B", "C", 1)
	// graph.addEdge("B", "D", 5)
	// graph.addEdge("C", "D", 8)
	// graph.addEdge("C", "E", 10)
	// graph.addEdge("D", "E", 2)
	// graph.addEdge("D", "T", 6)
	// graph.addEdge("E", "T", 2)
	// fmt.Println(graph.getPath("S", "T"))
}
