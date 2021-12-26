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
	polymer_pairs := map[string]int{}
	counts := map[string]int{}
	insertions := map[string]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if strings.Contains(line, "->") {
			insertion_rule := strings.Split(line, " -> ")
			insertions[insertion_rule[0]] = insertion_rule[1]
		} else {
			polymer := strings.Split(line, "")
			for i := 0; i < len(polymer)-1; i++ {
				v := polymer[i] + polymer[i+1]
				if _, ok := polymer_pairs[v]; !ok {
					polymer_pairs[v] = 0
				}
				polymer_pairs[v]++
				if _, ok := counts[polymer[i]]; !ok {
					counts[polymer[i]] = 0
				}
				counts[polymer[i]]++
			}
			if _, ok := counts[polymer[len(polymer)-1]]; !ok {
				counts[polymer[len(polymer)-1]] = 0
			}
			counts[polymer[len(polymer)-1]]++
		}
	}

	steps := 40
	for i := 0; i < steps; i++ {
		new_polymer_pairs := map[string]int{}

		for key, value := range polymer_pairs {
			if to_add, ok := insertions[key]; ok {
				add1, add2 := key[:1]+to_add, to_add+key[1:]

				if _, ok := new_polymer_pairs[add1]; !ok {
					new_polymer_pairs[add1] = 0
				}
				new_polymer_pairs[add1] += value
				if _, ok := new_polymer_pairs[add2]; !ok {
					new_polymer_pairs[add2] = 0
				}
				new_polymer_pairs[add2] += value

				if _, ok := counts[to_add]; !ok {
					counts[to_add] = 0
				}
				counts[to_add] += value
			}
		}

		polymer_pairs = new_polymer_pairs
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
