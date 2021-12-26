package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x float64
	y float64
	z float64
}

func ROTATIONS() [24][3][3]float64 {
	return [24][3][3]float64{{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
		{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},
		{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
		{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
		{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
		{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
		{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
		{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},
		{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
		{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
		{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
		{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
		{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
		{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
		{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
		{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
		{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
		{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
		{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
		{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
		{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
		{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
		{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}}}
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func coords_contains(coords []Coord, p Coord) bool {
	for _, c := range coords {
		if c.x == p.x && c.y == p.y && c.z == p.z {
			return true
		}
	}
	return false
}

func apply_rotate(p Coord, rotation [3][3]float64) Coord {
	return Coord{
		x: (p.x*rotation[0][0] + p.y*rotation[0][1] + p.z*rotation[0][2]),
		y: (p.x*rotation[1][0] + p.y*rotation[1][1] + p.z*rotation[1][2]),
		z: (p.x*rotation[2][0] + p.y*rotation[2][1] + p.z*rotation[2][2]),
	}
}

func apply_transform(p Coord, rotation [3][3]float64, translate [3]float64) Coord {
	return Coord{
		x: (p.x*rotation[0][0] + p.y*rotation[0][1] + p.z*rotation[0][2]) - translate[0],
		y: (p.x*rotation[1][0] + p.y*rotation[1][1] + p.z*rotation[1][2]) - translate[1],
		z: (p.x*rotation[2][0] + p.y*rotation[2][1] + p.z*rotation[2][2]) - translate[2],
	}
}

func fit_sensors(s1, s2 []Coord) []Coord {
	merged := s1

overall:
	for _, rot := range ROTATIONS() {
		for i, base2 := range s2 {
			rotated := apply_rotate(base2, rot)
			for _, base1 := range s1 {
				translate := [3]float64{
					rotated.x - base1.x,
					rotated.y - base1.y,
					rotated.z - base1.z,
				}
				temp_coords := []Coord{apply_transform(base2, rot, translate)}
				matches := []int{i}

				for j, test := range s2 {
					if i == j {
						continue
					}

					transformed_test := apply_transform(test, rot, translate)
					if coords_contains(s1, transformed_test) {
						matches = append(matches, j)
					}
					temp_coords = append(temp_coords, transformed_test)
				}

				if len(matches) >= 12 {
					for _, to_insert := range temp_coords {
						if !coords_contains(merged, to_insert) {
							merged = append(merged, to_insert)
						}
					}
					break overall
				}
			}
		}
	}

	return merged
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sensors := [][]Coord{}
	sensor := []Coord{}
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")

		if len(coords) == 3 {
			x, err := strconv.ParseFloat(coords[0], 64)
			if_err(err)
			y, err := strconv.ParseFloat(coords[1], 64)
			if_err(err)
			z, err := strconv.ParseFloat(coords[2], 64)
			if_err(err)
			sensor = append(sensor, Coord{x: x, y: y, z: z})
		} else {
			if len(sensor) > 0 {
				sensors = append(sensors, sensor)
			}
			sensor = []Coord{}
		}
	}
	sensors = append(sensors, sensor)

	fitted := sensors[0]
	remain := sensors[1:]
	i := 0

	for len(remain) > 0 {
		new_fitted := fit_sensors(fitted, remain[i])
		if len(new_fitted) == len(fitted) {
			i++
		} else {
			fitted = new_fitted
			remain = append(remain[:i], remain[i+1:]...)
			i = 0
		}
	}

	fmt.Println(len(fitted))
}
