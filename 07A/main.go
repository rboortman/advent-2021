package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
		distance += d
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

	// n := float64(len(points))
	// sumX := 0.0
	// sumY := 0.0
	// sumXY := 0.0
	// sumXX := 0.0

	// for _, p := range points {
	// 	sumX += p.x
	// 	sumY += p.y
	// 	sumXY += p.x * p.y
	// 	sumXX += p.x * p.x
	// }

	// fmt.Printf("sumX: %f\nsumY: %f\nsumXY: %f\nsumXX: %f\n", sumX, sumY, sumXY, sumXX)

	// base := (n*sumXX - sumX*sumX)
	// a := (n*sumXY - sumX*sumY) / base
	// // b := (sumXX*sumY - sumXY*sumX) / base
	// b := (sumY - a*sumX) / n

	// fmt.Printf("a: %f,b: %f, base: %f\n", a, b, base)

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})
	median := points[len(points)/2-2]

	// fmt.Println(points)
	fmt.Println(distance_to_point(points, median))
}
