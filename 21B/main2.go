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

type ValueBranch struct {
	place     int64
	value     int64
	universes int64
}

type Player struct {
	id            int64
	branches      map[string]ValueBranch
	universes_won int64
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
	// for i := int64(3); i <= 6; i += 3 {
	// 	for j := int64(3); j <= 6; j += 3 {
	// 		total := i + j
	// 		if _, ok := distribution[total]; !ok {
	// 			distribution[total] = 0
	// 		}
	// 		distribution[total]++
	// 	}
	// }
	// fmt.Printf("%+v\n", distribution)
	return distribution
}

func new_player_place(value int64) int64 {
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

func get_winning(ongoing [][]int64, wins []int64) []int64 {
	new_ongoings := [][]int64{}
	new_wins := wins
	distribution := get_roll_distribution()

	for _, branch := range ongoing {
		player_index := len(branch) % 2
		prev_value := int64(0)
		if len(branch) > 1 {
			prev_value = branch[len(branch)-2]
		}
		for rolled, times := range distribution {
			new_value := rolled + prev_value
			for i := int64(0); i < times; i++ {
				if new_value >= 21 {
					new_wins[player_index]++
				} else {
					fmt.Println(prev_value, rolled, new_value)
					new_ongoings = append(new_ongoings, append(branch, new_value))
				}
			}
		}
	}

	if len(new_ongoings) != 0 {
		fmt.Println(new_ongoings)
		fmt.Println(new_wins, len(new_ongoings))
		return get_winning(new_ongoings, new_wins)
	}

	return new_wins
}

func main() {
	// file, err := os.Open("./input.txt")
	file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	players := []Player{}

	r, err := regexp.Compile("(\\d+) starting position: (\\d+)")
	if_err(err)

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())
		p, err := strconv.ParseInt(line[1], 10, 64)
		if_err(err)
		start, err := strconv.ParseInt(line[2], 10, 64)
		if_err(err)

		key := fmt.Sprintf("%d,%d", 0, start)
		players = append(players, Player{
			id:            p,
			branches:      map[string]ValueBranch{key: {place: start, value: 0, universes: 1}},
			universes_won: 0,
		})
	}

	// winnings := get_winning([][]int64{{}}, []int64{0, 0})
	// fmt.Println(winnings)

	rolls := 0
	universes_spawned := int64(1)
	player_universes := []int64{0, 0}

	for true {
		distribution := get_roll_distribution()
		player_index := rolls % 2

		new_branches := map[string]ValueBranch{}
		new_universes_spawned := int64(0)
		universes_spawned_previously := int64(0)

		for _, branch := range players[player_index].branches {
			universes_spawned_previously += branch.universes
		}

		for rolled, times := range distribution {
			for _, branch := range players[player_index].branches {
				new_place := new_player_place(branch.place + rolled)
				new_value := branch.value + new_place
				new_key := fmt.Sprintf("%d,%d", new_value, new_place)
				// new_value := branch.value + rolled
				// new_key := fmt.Sprintf("%d", new_value)
				new_universes := (times * universes_spawned * branch.universes) / universes_spawned_previously

				if new_value >= 21 {
					player_universes[player_index] += new_universes
					new_universes_spawned -= new_universes
				} else {
					if _, ok := new_branches[new_key]; !ok {
						new_branches[new_key] = ValueBranch{
							// place:     new_place,
							value:     new_value,
							universes: new_universes,
						}
					} else {
						new_branches[new_key] = ValueBranch{
							// place:     new_place,
							value:     new_value,
							universes: new_branches[new_key].universes + new_universes,
						}
					}
				}
			}
			new_universes_spawned += universes_spawned * times
		}

		fmt.Println(player_universes, new_universes_spawned, universes_spawned_previously)
		fmt.Printf("%+v\n\n", new_branches)

		if len(new_branches) == 0 {
			break
		}

		rolls++
		players[player_index].branches = new_branches
		universes_spawned = new_universes_spawned
	}

	// for true {
	// 	fmt.Println(universes_spawned, universes_removed)
	// 	distribution := get_roll_distribution()
	// 	player_index := rolls % 2

	// 	new_branches := map[string]ValueBranch{}
	// 	new_universes_spawned := int64(0)
	// 	new_universes_removed := int64(0)

	// 	for key_d, value_d := range distribution {
	// 		for _, value_p := range players[player_index].branches {
	// 			new_place := new_player_place(value_p.place + key_d)
	// 			new_value := new_place + value_p.value
	// 			new_key := fmt.Sprintf("%d,%d", new_value, new_place)
	// 			new_universes := (value_p.universes - universes_removed) * universes_spawned * value_d

	// 			if new_value >= 21 {
	// 				player_universes[player_index] += new_universes
	// 				new_universes_removed += new_universes
	// 			} else {
	// 				if _, ok := new_branches[new_key]; !ok {
	// 					new_branches[new_key] = ValueBranch{
	// 						place:     new_place,
	// 						value:     new_value,
	// 						universes: new_universes,
	// 					}
	// 				} else {
	// 					new_branches[new_key] = ValueBranch{
	// 						place:     new_place,
	// 						value:     new_value,
	// 						universes: new_branches[new_key].universes + new_universes,
	// 					}
	// 				}
	// 				new_universes_spawned += value_d
	// 			}
	// 		}
	// 	}

	// 	if len(new_branches) == 0 {
	// 		break
	// 	}

	// 	fmt.Println(player_universes)
	// 	fmt.Printf("%+v\n\n", new_branches)

	// 	rolls++
	// 	players[player_index].branches = new_branches
	// 	universes_spawned = new_universes_spawned
	// 	universes_removed = new_universes_removed
	// }

	fmt.Println(player_universes, players)
}
