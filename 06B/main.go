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

func age_fish(fish [9]int, days int) [9]int {
	if days < 1 {
		return fish
	}

	new_fish := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i, c := range fish {
		if i < 1 {
			new_fish[6] += c
			new_fish[8] += c
		} else {
			new_fish[i-1] += c
		}
	}

	return age_fish(new_fish, days-1)
}

func main() {
	file, err := os.Open("./input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fish := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for scanner.Scan() {
		fish_strings := strings.Split(scanner.Text(), ",")
		for _, ch := range fish_strings {
			d, err := strconv.Atoi(ch)
			if_err(err)
			fish[d]++
		}
	}

	days := 256
	new_fish := age_fish(fish, days)

	sum := 0
	for _, c := range new_fish {
		sum += c
	}
	fmt.Println(sum)
}
