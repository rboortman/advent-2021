package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

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
	polymer := []string{}
	insertions := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if strings.Contains(line, "->") {
			insertion_rule := strings.Split(line, " -> ")
			insertions[insertion_rule[0]] = insertion_rule[1]
		} else {
			polymer = strings.Split(line, "")
		}
	}

	steps := 10
	for i := 0; i < steps; i++ {
		for j := len(polymer) - 2; j >= 0; j-- {
			pair := polymer[j] + polymer[j+1]
			if to_insert, ok := insertions[pair]; ok {
				polymer = append(polymer[:j+2], polymer[j+1:]...)
				polymer[j+1] = to_insert
			}
		}
	}

	counts := map[string]int{}
	for _, v := range polymer {
		if _, ok := counts[v]; !ok {
			counts[v] = 0
		}
		counts[v]++
	}

	most := 0
	least := math.MaxInt
	for _, value := range counts {
		if value > most {
			most = value
		}
		if value < least {
			least = value
		}
	}

	fmt.Println(most - least)
}
