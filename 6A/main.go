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

func age_fish(fish []int, days int) []int {
	if days < 1 {
		return fish
	}

	new_fish := fish
	baby_fish := 0
	for i, f := range new_fish {
		if f-1 < 0 {
			new_fish[i] = 6
			baby_fish++
		} else {
			new_fish[i] = f - 1
		}
	}
	for i := 0; i < baby_fish; i++ {
		new_fish = append(new_fish, 8)
	}

	return age_fish(new_fish, days-1)
}

func main() {
	file, err := os.Open("./input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fish := []int{}
	for scanner.Scan() {
		fish_strings := strings.Split(scanner.Text(), ",")
		for _, ch := range fish_strings {
			d, err := strconv.Atoi(ch)
			if_err(err)
			fish = append(fish, d)
		}
	}

	fmt.Println(len(age_fish(fish, 80)))
}
