package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func pretty_print_image(image [][]uint64) {
	for _, line := range image {
		pretty_line := ""
		for _, p := range line {
			to_print := "."
			if p == 1 {
				to_print = "#"
			}
			pretty_line += to_print
		}
		fmt.Println(pretty_line)
	}
}

func count_lit(image [][]uint64) int {
	counter := 0
	for _, line := range image {
		for _, p := range line {
			counter += int(p)
		}
	}
	return counter
}

func enlarge_image(image [][]uint64) [][]uint64 {
	first, last := []uint64{}, []uint64{}
	for i := 0; i < len(image[0])+2; i++ {
		first = append(first, uint64(0))
		last = append(last, uint64(0))
	}
	new_image := [][]uint64{first}

	for _, line := range image {
		new_line := []uint64{uint64(0)}
		new_line = append(new_line, line...)
		new_line = append(new_line, uint64(0))
		new_image = append(new_image, new_line)
	}

	new_image = append(new_image, last)

	return new_image
}

func get_value(image [][]uint64, i, j, iteration int) uint64 {
	base := uint64(iteration % 2)
	// if iteration % 2 == 0 {
	// 	base = uint64(1)
	// }
	if i < 0 || len(image) <= i {
		return base
	} else if j < 0 || len(image[i]) <= j {
		return base
	} else {
		return uint64(image[i][j])
	}
}

func enhance_image(image [][]uint64, enhancer []uint64, iteration int) [][]uint64 {
	// base_image := enlarge_image(enlarge_image(image))
	base_image := image
	new_image := [][]uint64{}

	for i := -1; i < len(base_image)+1; i++ {
		new_line := []uint64{}
		for j := -1; j < len(base_image[0])+1; j++ {
			b := uint64(0)
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					b = b << 1
					b += get_value(base_image, k, l, iteration)
				}
			}
			new_line = append(new_line, enhancer[b])
		}
		new_image = append(new_image, new_line)
	}

	return new_image
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	enhancer := []uint64{}
	image := [][]uint64{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		if len(line) < 1 {
			continue
		} else if len(line) > 200 {
			for _, p := range line {
				to_add := uint64(0)
				if p == "#" {
					to_add = uint64(1)
				}
				enhancer = append(enhancer, to_add)
			}
		} else {
			pixel_line := []uint64{}
			for _, p := range line {
				to_add := uint64(0)
				if p == "#" {
					to_add = uint64(1)
				}
				pixel_line = append(pixel_line, to_add)
			}
			image = append(image, pixel_line)
		}
	}

	for i := 0; i < 2; i++ {
		image = enhance_image(image, enhancer, i)
	}

	// pretty_print_image(image)
	fmt.Println(count_lit(image))
}
