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

func reverse(rules []ReactorRule) []ReactorRule {
	for i, j := 0, len(rules)-1; i < j; i, j = i+1, j-1 {
		rules[i], rules[j] = rules[j], rules[i]
	}
	return rules
}

func get_rule_volume(rule ReactorRule) int64 {
	return int64(math.Abs(float64((rule.x.max - rule.x.min + 1) * (rule.y.max - rule.y.min + 1) * (rule.z.max - rule.z.min + 1))))
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
	is_on := int64(0)
	if check.on {
		is_on = get_rule_volume(check)
	}

	if len(rules) > 0 {
		total_sub := total_on(rules[1:], rules[0])
		total_overlap_sub := int64(0)
		overlap := find_overlap(rules, check)
		if len(overlap) > 0 {
			trimmed := trim_rules(overlap, check)
			total_overlap_sub = total_on(trimmed[1:], trimmed[0])
		}

		is_on += total_sub
		is_on -= total_overlap_sub
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
		trimmed := reverse(trim_rules(overlap, rule))

		if rule.on {
			turned_on += get_rule_volume(rule)
		}

		if len(trimmed) == 1 && trimmed[0].on {
			turned_on -= get_rule_volume(trimmed[0])
		} else if len(trimmed) > 1 {
			turned_on -= total_on(trimmed[1:], trimmed[0])
		}

	}

	// fmt.Printf("%+v\n", rules)
	fmt.Println(turned_on)
}
