package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Amphipod struct {
	energy  int
	t       int
	sec_num int
	id      int
}

type Board struct {
	hall  []int
	rooms [][]int
	pods  []Amphipod
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func print_situation(situation Board) {
	t_to_s := []string{"A", "B", "C", "D"}
	fmt.Printf("\n#############\n")
	fmt.Printf("#")
	for _, pos := range situation.hall {
		to_print := "."
		if pos > 0 {
			to_print = fmt.Sprintf("%s", t_to_s[situation.pods[pos-1].t])
		}
		fmt.Printf(to_print)
	}
	fmt.Printf("#\n")

	for i := 0; i < len(situation.rooms[0]); i++ {
		prefix := "  #"
		if i == 0 {
			prefix = "###"
		}
		fmt.Printf(prefix)
		for _, room := range situation.rooms {
			to_print := "."
			if room[i] > 0 {
				to_print = fmt.Sprintf("%s", t_to_s[situation.pods[room[i]-1].t])
			}
			to_print += "#"
			fmt.Printf(to_print)
		}
		suffix := "\n"
		if i == 0 {
			suffix = "##\n"
		}
		fmt.Printf(suffix)
	}

	fmt.Printf("  #########\n")
}

func deep_copy_board(situation Board) Board {
	new_hall := []int{}
	new_rooms := [][]int{}
	new_pods := situation.pods[:]

	for _, id := range situation.hall {
		new_hall = append(new_hall, id)
	}

	for _, room := range situation.rooms {
		new_room := []int{}
		for _, id := range room {
			new_room = append(new_room, id)
		}
		new_rooms = append(new_rooms, new_room)
	}

	return Board{
		hall:  new_hall,
		rooms: new_rooms,
		pods:  new_pods,
	}
}

func is_done(situation Board) bool {
	done := true
	for i, room := range situation.rooms {
		for _, id := range room {
			done = done && id > 0 && situation.pods[id-1].t == i
		}
	}
	return done
}

func get_hallway_distance(situation Board, from, to, id int) int {
	is_clear := true
	min := int(math.Min(float64(from), float64(to)))
	max := int(math.Max(float64(from), float64(to)))

	for i := min; i <= max && is_clear; i++ {
		is_clear = is_clear && (situation.hall[i] == 0 || situation.hall[i] == id)
	}

	if is_clear {
		return max - min
	} else {
		return -1
	}
}

func move_into_room(situation Board, room_id int) int {
	room := situation.rooms[room_id]
	for i := len(room) - 1; i >= 0; i-- {
		if room[i] == 0 {
			return i
		} else if situation.pods[room[i]-1].t != room_id {
			return -1
		}
	}
	return -1
}

func move_out_room(situation Board, room_id, room_location int) int {
	room := situation.rooms[room_id]
	for i := 0; i < room_location; i++ {
		if room[i] != 0 {
			return -1
		}
	}
	return room_location
}

func is_in_correct_place(situation Board, room_id, room_location int) bool {
	room := situation.rooms[room_id]
	pod := situation.pods[room[room_location]-1]
	correct_place := pod.t == room_id
	for i := len(room) - 1; i > room_location; i-- {
		correct_place = correct_place && situation.pods[room[i]-1].t == room_id
	}
	return correct_place
}

func min(energy []int) int {
	min_energy := math.MaxInt
	for _, e := range energy {
		if e < min_energy {
			min_energy = e
		}
	}
	return min_energy
}

func find_done_states(situation Board, energy_spent int, previous_minimum int) []int {
	if is_done(situation) {
		return []int{energy_spent}
	} else if previous_minimum < energy_spent {
		return []int{}
	}
	energy_levels := []int{}

	// print_situation(situation)

	for i, pod_id := range situation.hall {
		if pod_id <= 0 {
			continue
		}

		pod := situation.pods[pod_id-1]
		room_hallway_id := 2 + pod.t*2

		distance := get_hallway_distance(situation, i, room_hallway_id, pod.id)
		room_place := move_into_room(situation, pod.t)

		if distance >= 0 && room_place >= 0 {
			new_situation := deep_copy_board(situation)
			pod_energy := distance*pod.energy + (1+room_place)*pod.energy

			new_situation.hall[i] = 0
			new_situation.rooms[pod.t][room_place] = pod.id

			// fmt.Println("Send 1", energy_spent+pod_energy, previous_minimum)
			energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy, previous_minimum)...)

			min_energy := min(energy_levels)
			if min_energy < previous_minimum {
				previous_minimum = min_energy
			}
		}
	}

	for i, room := range situation.rooms {
		for j, pod_id := range room {
			if pod_id <= 0 || is_in_correct_place(situation, i, j) {
				continue
			}

			pod := situation.pods[pod_id-1]

			out_distance := move_out_room(situation, i, j)
			if out_distance < 0 {
				continue
			}

			pod_hallway_id := 2 + i*2
			room_hallway_id := 2 + pod.t*2
			distance_to_room := get_hallway_distance(situation, pod_hallway_id, room_hallway_id, pod.id)
			room_place := move_into_room(situation, pod.t)

			if distance_to_room >= 0 && room_place >= 0 {
				new_situation := deep_copy_board(situation)
				pod_energy := (1+out_distance)*pod.energy + distance_to_room*pod.energy + (1+room_place)*pod.energy

				new_situation.rooms[i][j] = 0
				new_situation.rooms[pod.t][room_place] = pod.id

				// fmt.Println("Send 2")
				energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy, previous_minimum)...)

				min_energy := min(energy_levels)
				if min_energy < previous_minimum {
					previous_minimum = min_energy
				}
			} else {
				for h, pid := range situation.hall {
					if h == 2 || h == 4 || h == 6 || h == 8 {
						continue
					}
					if pid != 0 {
						continue
					}

					distance := get_hallway_distance(situation, pod_hallway_id, h, pod.id)
					if distance <= 0 {
						continue
					}

					new_situation := deep_copy_board(situation)
					pod_energy := (1+out_distance)*pod.energy + distance*pod.energy

					new_situation.rooms[i][j] = 0
					new_situation.hall[h] = pod.id

					// fmt.Println("Send 3")
					energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy, previous_minimum)...)

					min_energy := min(energy_levels)
					if min_energy < previous_minimum {
						previous_minimum = min_energy
					}
				}
			}
		}
	}

	return energy_levels
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, err := regexp.Compile("#([A|B|C|D])#([A|B|C|D])#([A|B|C|D])#([A|B|C|D])#")
	if_err(err)
	secs := [4]int{0, 0, 0, 0}
	type_map := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}
	energy_map := [4]int{1, 10, 100, 1000}

	// situation := Board{hall: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, rooms: [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}}, pods: []Amphipod{}}
	situation := Board{hall: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, rooms: [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, pods: []Amphipod{}}

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())

		if len(line) > 1 {
			for i, amphi := range line {
				if i == 0 {
					continue
				}
				t := type_map[amphi]
				total := 1
				for _, s := range secs {
					total += s
				}
				j := (total - 1) / 4

				situation.rooms[i-1][j] = total
				situation.pods = append(situation.pods, Amphipod{
					energy:  energy_map[t],
					t:       t,
					sec_num: secs[t],
					id:      total,
				})

				secs[t]++
			}
		}
	}

	done_states := find_done_states(situation, 0, math.MaxInt)

	// print_situation(situation)
	fmt.Println(min(done_states))
}
