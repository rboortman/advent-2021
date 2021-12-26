package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cave struct {
	out    []string
	name   string
	is_big bool
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func find_paths(chart map[string]Cave, path []string, from_cave string, has_visited_small_cave bool) [][]string {
	new_paths := [][]string{}

	for _, c := range chart[from_cave].out {
		if c == "end" {
			new_paths = append(new_paths, append(path, c))
		} else if c == "start" {
			continue
		} else if !contains(path, c) {
			new_paths = append(new_paths, find_paths(chart, append(path, c), c, has_visited_small_cave)...)
		} else {
			if chart[c].is_big {
				new_paths = append(new_paths, find_paths(chart, append(path, c), c, has_visited_small_cave)...)
			} else if !has_visited_small_cave {
				new_paths = append(new_paths, find_paths(chart, append(path, c), c, true)...)
			}
		}
	}

	return new_paths
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chart := map[string]Cave{}
	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		A, B := path[0], path[1]
		if _, ok := chart[A]; !ok {
			chart[A] = Cave{out: []string{}, name: A, is_big: strings.ToUpper(A) == A}
		}
		if _, ok := chart[B]; !ok {
			chart[B] = Cave{out: []string{}, name: B, is_big: strings.ToUpper(B) == B}
		}

		chart[A] = Cave{out: append(chart[A].out, B), name: A, is_big: chart[A].is_big}
		chart[B] = Cave{out: append(chart[B].out, A), name: B, is_big: chart[B].is_big}
	}

	fmt.Println(len(find_paths(chart, []string{}, "start", false)))
}
