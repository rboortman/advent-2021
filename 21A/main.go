package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

type Player struct {
	id     int
	place  int
	value  int
	rolled int
}

func get_total_rolls(roll_id int) (int, int, int) {
	base := (roll_id * 3) % 100
	roll1 := (base+1)%101 + (base+1)/101
	roll2 := (base+2)%101 + (base+2)/101
	roll3 := (base+3)%101 + (base+3)/101
	return roll1, roll2, roll3
}

func new_player_place(value int) int {
	// new_value := value%10 + value/10
	// if new_value > 10 {
	// 	new_value = new_player_place(new_value)
	// }
	new_value := value % 10
	if new_value == 0 {
		new_value = 10
	}
	return new_value
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	players := []Player{}

	r, err := regexp.Compile("(\\d+) starting position: (\\d+)")
	if_err(err)

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())
		p, err := strconv.Atoi(line[1])
		if_err(err)
		start, err := strconv.Atoi(line[2])
		if_err(err)

		players = append(players, Player{
			id:     p,
			place:  start,
			value:  0,
			rolled: 0,
		})
	}

	rolls := 0
	for true {
		to_add1, to_add2, to_add3 := get_total_rolls(rolls)
		player_index := rolls % 2

		end_value := players[player_index].place + to_add1 + to_add2 + to_add3
		players[player_index].place = new_player_place(end_value)
		players[player_index].value += players[player_index].place
		players[player_index].rolled += 3

		// fmt.Printf("Player %d rolls (%d+%d+%d) and moves to place %d for a total score of %d\n", player_index+1, to_add1, to_add2, to_add3, players[player_index].place, players[player_index].value)

		if players[player_index].value >= 1000 {
			break
		}
		rolls++
	}

	loser := Player{}
	for _, p := range players {
		if p.value < 1000 {
			loser = p
		}
	}

	fmt.Println(((rolls + 1) * 3) * loser.value)
}
