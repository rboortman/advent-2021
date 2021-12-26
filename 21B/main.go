package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type PlayerState struct {
	score    int64
	position int64
}

type GameState struct {
	playerA   PlayerState
	playerB   PlayerState
	universes int64
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func get_roll_distribution() map[int64]int64 {
	distribution := map[int64]int64{}
	for i := int64(1); i <= 3; i++ {
		for j := int64(1); j <= 3; j++ {
			for k := int64(1); k <= 3; k++ {
				total := i + j + k
				if _, ok := distribution[total]; !ok {
					distribution[total] = 0
				}
				distribution[total]++
			}
		}
	}
	return distribution
}

func new_player_place(value int64) int64 {
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

	r, err := regexp.Compile("(\\d+) starting position: (\\d+)")
	if_err(err)

	firstState := GameState{
		universes: 1,
	}

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())
		p, err := strconv.ParseInt(line[1], 10, 64)
		if_err(err)
		start, err := strconv.ParseInt(line[2], 10, 64)
		if_err(err)

		player := PlayerState{
			score:    0,
			position: start,
		}

		if p == 1 {
			firstState.playerA = player
		} else {
			firstState.playerB = player
		}
	}

	states := []GameState{firstState}
	wins := [2]int64{0, 0}
	rolls := 0

	distribution := [][]int64{{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1}}

	for true {
		new_states := []GameState{}

		for _, state := range states {
			for _, dist := range distribution {
				current_player := state.playerA
				if rolls%2 == 1 {
					current_player = state.playerB
				}
				new_position := new_player_place(current_player.position + dist[0])
				new_score := current_player.score + new_position
				new_universes := state.universes * dist[1]

				if new_score >= 21 {
					wins[rolls%2] += new_universes
					continue
				}

				new_game_state := GameState{}

				if rolls%2 == 0 {
					new_game_state = GameState{
						universes: new_universes,
						playerA: PlayerState{
							position: new_position,
							score:    new_score,
						},
						playerB: state.playerB,
					}
				} else {
					new_game_state = GameState{
						universes: new_universes,
						playerA:   state.playerA,
						playerB: PlayerState{
							position: new_position,
							score:    new_score,
						},
					}
				}

				new_states = append(new_states, new_game_state)
			}
		}

		states = new_states
		rolls++

		if len(states) == 0 {
			break
		}
	}

	fmt.Printf("%+v\n", wins)
}
