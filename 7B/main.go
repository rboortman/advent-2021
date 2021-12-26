package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x float64
	y float64
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func distance_to_point(points []Point, p_calc Point) int {
	distance := 0

	for _, p := range points {
		d := int(math.Abs(p.x - p_calc.x))
		total := 0
		for i := 0; i <= d; i++ {
			total += i
		}
		distance += total
	}

	return distance
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := []Point{}
	for scanner.Scan() {
		points_strings := strings.Split(scanner.Text(), ",")
		for _, ch := range points_strings {
			d, err := strconv.Atoi(ch)
			if_err(err)
			points = append(points, Point{x: float64(d), y: 1.0})
		}
	}

	sum := 0.0
	for _, p := range points {
		sum += p.x
	}
	mean := sum / float64(len(points))
	rounded_mean := math.Round(mean) - 1

	fmt.Println(mean)
	fmt.Println(rounded_mean)
	fmt.Println(distance_to_point(points, Point{x: rounded_mean, y: 1.0}))
}
