package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Vector struct {
	x int
	y int
}

type Coord struct {
	x int
	y int
}

type Area struct {
	topLeft     Coord
	bottomRight Coord
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func lands_in_trench(shot Vector, probe Coord, target Area, highest int) (int, int) {
	// fmt.Printf("%+v\n", probe)
	if shot.x == 0 {
		if probe.x < target.topLeft.x {
			return -1, highest
		} else if probe.x > target.bottomRight.x {
			return 1, highest
		} else if probe.y < target.bottomRight.y {
			return 2, highest
		}
	}

	if target.topLeft.x <= probe.x && probe.x <= target.bottomRight.x && target.bottomRight.y <= probe.y && probe.y <= target.topLeft.y {
		return 0, highest
	}

	new_vector := Vector{x: int(math.Max(float64(shot.x-1), 0)), y: shot.y - 1}
	new_coord := Coord{x: probe.x + shot.x, y: probe.y + shot.y}
	new_highest := int(math.Max(float64(highest), float64(new_coord.y)))

	return lands_in_trench(new_vector, new_coord, target, new_highest)
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	target := Area{}
	for scanner.Scan() {
		r, err := regexp.Compile("x=(-?\\d+)..(-?\\d+), y=(-?\\d+)..(-?\\d+)")
		if_err(err)
		line := r.FindStringSubmatch(scanner.Text())

		d1, err := strconv.ParseFloat(line[1], 64)
		if_err(err)
		d2, err := strconv.ParseFloat(line[2], 64)
		if_err(err)

		d3, err := strconv.ParseFloat(line[3], 64)
		if_err(err)
		d4, err := strconv.ParseFloat(line[4], 64)
		if_err(err)

		target = Area{
			topLeft:     Coord{x: int(math.Min(d1, d2)), y: int(math.Max(d3, d4))},
			bottomRight: Coord{x: int(math.Max(d1, d2)), y: int(math.Min(d3, d4))},
		}
	}

	// lands_in_trench(Vector{7, 2}, Coord{0, 0}, target)

	top := 0
	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			result, highest := lands_in_trench(Vector{x: i, y: j}, Coord{x: 0, y: 0}, target, 0)
			if result == 0 {
				if highest > top {
					top = highest
				}
				// fmt.Printf("Top: %d, speed: %d, angle: %d\n", highest, i, j)
			}
		}
	}

	fmt.Println(top)
}
