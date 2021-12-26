package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type CoordRange struct {
	min int
	max int
}

type ReactorRule struct {
	x  CoordRange
	y  CoordRange
	z  CoordRange
	on int
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, err := regexp.Compile("^(on|off) x=(-?\\d+)..(-?\\d+),y=(-?\\d+)..(-?\\d+),z=(-?\\d+)..(-?\\d+)$")
	if_err(err)

	rules := []ReactorRule{}

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())
		on := 0
		if line[1] == "on" {
			on = 1
		}
		x_min, _ := strconv.Atoi(line[2])
		x_max, _ := strconv.Atoi(line[3])
		y_min, _ := strconv.Atoi(line[4])
		y_max, _ := strconv.Atoi(line[5])
		z_min, _ := strconv.Atoi(line[6])
		z_max, _ := strconv.Atoi(line[7])

		rule := ReactorRule{
			x:  CoordRange{min: x_min, max: x_max},
			y:  CoordRange{min: y_min, max: y_max},
			z:  CoordRange{min: z_min, max: z_max},
			on: on,
		}

		rules = append(rules, rule)
	}

	reactor := [101][101][101]int{}

	for _, rule := range rules {
		for i := rule.x.min; i <= rule.x.max; i++ {
			if -50 <= i && i <= 50 {
				for j := rule.y.min; j <= rule.y.max; j++ {
					if -50 <= j && j <= 50 {
						for k := rule.z.min; k <= rule.z.max; k++ {
							if -50 <= k && k <= 50 {
								reactor[i+50][j+50][k+50] = rule.on
							}
						}
					}
				}
			}
		}
	}

	count := 0
	for _, axis1 := range reactor {
		for _, axis2 := range axis1 {
			for _, axis3 := range axis2 {
				count += axis3
			}
		}
	}

	fmt.Println(count)
}
