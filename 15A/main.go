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

type Path struct {
	path  []Point
	value int
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func contains_point(current_path []Point, p Point) bool {
	contains := false
	for _, cp := range current_path {
		contains = contains || (cp.x == p.x && cp.y == p.y)
		if contains {
			break
		}
	}
	return contains
}

func get_value_path(path []Point) int {
	count := 0
	for _, p := range path {
		count += p.value
	}
	return count
}

func find_paths(grid [][]Point, current_path []Point, end Point) []Path {
	new_paths := []Path{}
	last_point := current_path[len(current_path)-1]

	if last_point.x == end.x && last_point.y == end.y {
		new_path := Path{path: current_path, value: get_value_path(current_path)}
		return []Path{new_path}
	}

	if last_point.y > 0 && !contains_point(current_path, grid[last_point.y-1][last_point.x]) {
		new_paths = append(new_paths, find_paths(grid, append(current_path, grid[last_point.y-1][last_point.x]), end)...)
	}
	if last_point.y < len(grid)-1 && !contains_point(current_path, grid[last_point.y+1][last_point.x]) {
		new_paths = append(new_paths, find_paths(grid, append(current_path, grid[last_point.y+1][last_point.x]), end)...)
	}
	if last_point.x > 0 && !contains_point(current_path, grid[last_point.y][last_point.x-1]) {
		new_paths = append(new_paths, find_paths(grid, append(current_path, grid[last_point.y][last_point.x-1]), end)...)
	}
	if last_point.x < len(grid[last_point.y])-1 && !contains_point(current_path, grid[last_point.y][last_point.x+1]) {
		new_paths = append(new_paths, find_paths(grid, append(current_path, grid[last_point.y][last_point.x+1]), end)...)
	}

	return new_paths
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

	// all_paths := find_paths(grid, []Point{grid[0][0]}, grid[len(grid)-1][len(grid[len(grid)-1])-1])

	// fmt.Println(all_paths[0])

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
