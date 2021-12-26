package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_for_bingo(board [][]int) bool {
	has_bingo := false
	colums := []int{}
	for i, line := range board {
		line_sum := 0
		for j, cell := range line {
			line_sum += cell
			if i == 0 {
				colums = append(colums, cell)
			} else {
				colums[j] += cell
			}
		}
		has_bingo = line_sum == len(line)*-1
		if has_bingo {
			break
		}
	}
	if !has_bingo {
		for _, sum := range colums {
			has_bingo = sum == len(colums)*-1
			if has_bingo {
				break
			}
		}
	}
	return has_bingo
}

func cross_number(boards [][][]int, number int) [][][]int {
	new_boards := boards
	for i, board := range new_boards {
		for j, line := range board {
			for k, cell := range line {
				if cell == number {
					new_boards[i][j][k] = -1
				}
			}
		}
	}
	return new_boards
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	digits := []int{}
	boards := [][][]int{}
	currentBoard := [][]int{}
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 50 {
			split_digits := strings.Split(line, ",")
			for _, ch := range split_digits {
				d, err := strconv.Atoi(ch)
				if err != nil {
					panic(err)
				}
				digits = append(digits, d)
			}
		} else if len(line) < 10 {
			if len(currentBoard) > 1 {
				boards = append(boards, currentBoard)
				currentBoard = [][]int{}
			}
		} else {
			bingo_line := strings.Fields(line)
			bingo_digits := []int{}
			for _, ch := range bingo_line {
				d, err := strconv.Atoi(ch)
				if err != nil {
					panic(err)
				}
				bingo_digits = append(bingo_digits, d)
			}
			currentBoard = append(currentBoard, bingo_digits)
		}
	}
	boards = append(boards, currentBoard)

	bingo_board := -1
	winning_digit := -1
	for _, d := range digits {
		boards = cross_number(boards, d)
		for i, board := range boards {
			if check_for_bingo(board) {
				bingo_board = i
			}
		}
		if bingo_board > 0 {
			winning_digit = d
			break
		}
	}

	sum := 0
	for _, line := range boards[bingo_board] {
		for _, cell := range line {
			if cell > 0 {
				sum += cell
			}
		}
	}

	// fmt.Println(boards)
	fmt.Println(sum * winning_digit)
	// fmt.Printf("horizontal: %d\nvertical: %d\ntogether: %d\n", horizontal, vertical, horizontal*vertical)
}
