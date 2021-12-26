package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type CoordRange struct {
	min int64
	max int64
}

type CoordRange3D struct {
	x CoordRange
	y CoordRange
	z CoordRange
}

type ReactorRule struct {
	x  CoordRange
	y  CoordRange
	z  CoordRange
	on bool
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func get_rule_volume(rule ReactorRule) int64 {
	return (rule.x.max - rule.x.min) * (rule.y.max - rule.y.min) * (rule.z.max - rule.z.min)
}

func trim_rules(rules []ReactorRule, mask ReactorRule) []ReactorRule {
	trimmed := []ReactorRule{}
	for _, rule := range rules {
		trimmed = append(trimmed, ReactorRule{
			x:  CoordRange{min: int64(math.Max(float64(rule.x.min), float64(mask.x.min))), max: int64(math.Min(float64(rule.x.max), float64(mask.x.max)))},
			y:  CoordRange{min: int64(math.Max(float64(rule.y.min), float64(mask.y.min))), max: int64(math.Min(float64(rule.y.max), float64(mask.y.max)))},
			z:  CoordRange{min: int64(math.Max(float64(rule.z.min), float64(mask.z.min))), max: int64(math.Min(float64(rule.z.max), float64(mask.z.max)))},
			on: rule.on,
		})
	}
	return trimmed
}

func find_overlap(rules []ReactorRule, check ReactorRule) []ReactorRule {
	overlap := []ReactorRule{}
	for _, rule := range rules {
		if rule.x.min <= check.x.max && check.x.min <= rule.x.max && rule.y.min <= check.y.max && check.y.min <= rule.y.max && rule.z.min <= check.z.max && check.z.min <= rule.z.max {
			overlap = append(overlap, rule)
		}
	}
	return overlap
}

func total_on(rules []ReactorRule, check ReactorRule) int64 {
	is_on := get_rule_volume(check)

	if len(rules) < 1 {
		return is_on
	} else if len(rules) == 1 {
		rule := rules[0]
		rule_volume := get_rule_volume(rule)
		if rule.on {
			is_on += rule_volume
		}
		for _, o := range find_overlap(rules, check) {
			is_on -= get_rule_volume(o)
		}
	}

	return is_on
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
		on := false
		if line[1] == "on" {
			on = true
		}
		x1, _ := strconv.ParseFloat(line[2], 64)
		x2, _ := strconv.ParseFloat(line[3], 64)
		y1, _ := strconv.ParseFloat(line[4], 64)
		y2, _ := strconv.ParseFloat(line[5], 64)
		z1, _ := strconv.ParseFloat(line[6], 64)
		z2, _ := strconv.ParseFloat(line[7], 64)

		rule := ReactorRule{
			x:  CoordRange{min: int64(math.Min(x1, x2)), max: int64(math.Max(x1, x2))},
			y:  CoordRange{min: int64(math.Min(y1, y2)), max: int64(math.Max(y1, y2))},
			z:  CoordRange{min: int64(math.Min(z1, z2)), max: int64(math.Max(z1, z2))},
			on: on,
		}

		rules = append(rules, rule)
	}

	turned_on := int64(0)
	for i, rule := range rules {
		overlap := find_overlap(rules[:i], rule)

		// fmt.Printf("%+v\n", overlap)
		// fmt.Println(len(overlap))
		if len(overlap) > 0 {
			trimmed := trim_rules(overlap, rule)
			count := int64(0)
			for _, t := range trimmed {
				count += (t.x.max - t.x.min) * (t.y.max - t.y.min) * (t.z.max - t.z.min)
			}
			fmt.Println(count)
		}
	}

	fmt.Println(turned_on)
}
