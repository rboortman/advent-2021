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

func sum(arr []int) int {
	s := 0
	for _, d := range arr {
		s += d
	}
	return s
}

func print_situation(situation Board) {
	t_to_s := []string{"A", "B", "C", "D"}
	fmt.Printf("\n########################\n")
	fmt.Printf("#")
	for _, pos := range situation.hall {
		to_print := ". "
		if pos > 0 {
			to_print = fmt.Sprintf("%s%d", t_to_s[situation.pods[pos-1].t], situation.pods[pos-1].id)
		}
		fmt.Printf(to_print)
	}
	fmt.Printf("#\n")

	fmt.Printf("#####")
	for _, room := range situation.rooms {
		to_print := ". "
		if room[0] > 0 {
			to_print = fmt.Sprintf("%s%d", t_to_s[situation.pods[room[0]-1].t], situation.pods[room[0]-1].id)
		}
		to_print += "##"
		fmt.Printf(to_print)
	}
	fmt.Printf("###\n")

	fmt.Printf("    #")
	for _, room := range situation.rooms {
		to_print := ". "
		if room[1] > 0 {
			to_print = fmt.Sprintf("%s%d", t_to_s[situation.pods[room[1]-1].t], situation.pods[room[1]-1].id)
		}
		to_print += "##"
		fmt.Printf(to_print)
	}
	fmt.Printf("\n")

	fmt.Printf("    #################\n")
}

func deep_copy_board(situation Board) Board {
	new_hall := []int{}
	new_rooms := [][]int{}
	new_pods := situation.pods[:]

	for _, id := range situation.hall {
		new_hall = append(new_hall, id)
	}

	for _, room := range situation.rooms {
		new_rooms = append(new_rooms, []int{room[0], room[1]})
	}

	return Board{
		hall:  new_hall,
		rooms: new_rooms,
		pods:  new_pods,
	}
}

func determain_type(s string) int {
	t := -1
	switch s {
	case "A":
		t = 0
	case "B":
		t = 1
	case "C":
		t = 2
	case "D":
		t = 3
	}
	return t
}

func can_move(situation Board, to_move int, last_moved int) bool {
	allowed := false
	if to_move == last_moved {
		return allowed
	}

	pod := situation.pods[to_move-1]

	for i, id := range situation.hall {
		if id == to_move {
			if situation.rooms[pod.t][0] != 0 && (situation.rooms[pod.t][1] != 0 || situation.rooms[pod.t][1] != pod.t) {
				return allowed
			}
			room_hallway_id := 2 + pod.t*2

			max := int(math.Max(float64(room_hallway_id), float64(i)))
			min := int(math.Min(float64(room_hallway_id), float64(i)))

			for i := min; i < max; i++ {
				if situation.hall[i] != 0 {
					return allowed
				}
			}

			allowed = true
			break
		}
	}

	if !allowed {
		for i, room := range situation.rooms {
			if room[0] == to_move || room[1] == to_move {
				if room[0] == 0 || room[0] == to_move {
					hallway_i := 2 + i*2
					if situation.hall[hallway_i-1] == 0 || situation.hall[hallway_i+1] == 0 {
						allowed = true
						break
					}
				}
			}
		}
	}

	return allowed
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

func find_pod(situation Board, id int) Amphipod {
	to_return := Amphipod{
		t:       -1,
		energy:  -1,
		sec_num: -1,
		id:      -1,
	}
	for _, pod := range situation.pods {
		if pod.id == id {
			to_return = pod
			break
		}
	}
	return to_return
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

func find_done_states(situation Board, energy_spent int) []int {
	if is_done(situation) {
		// fmt.Println("got here")
		return []int{energy_spent}
	}
	// print_situation(situation)

	if energy_spent > 16000 {
		return []int{}
	}

	energy_levels := []int{}

amphipod_loop:
	for _, pod := range situation.pods {

	hall_loop:
		for i, id := range situation.hall {
			if id == pod.id {
				other_pod := find_pod(situation, situation.rooms[pod.t][1])
				// fmt.Printf("pod: %+v, other_pod: %+v, if: %t, %t, %t, \n", pod, other_pod, situation.rooms[pod.t][0] != 0, other_pod.t != -1, other_pod.t != pod.t)
				if situation.rooms[pod.t][0] != 0 || (other_pod.t != -1 && other_pod.t != pod.t) {
					continue amphipod_loop
				}
				room_hallway_id := 2 + pod.t*2
				distance := get_hallway_distance(situation, i, room_hallway_id, pod.id)

				if distance < 0 {
					continue amphipod_loop
				}

				pod_energy := distance*pod.energy + pod.energy
				new_situation := deep_copy_board(situation)
				new_situation.hall[i] = 0
				if situation.rooms[pod.t][1] == 0 {
					pod_energy += pod.energy
					new_situation.rooms[pod.t][1] = pod.id
				} else {
					new_situation.rooms[pod.t][0] = pod.id
				}
				// fmt.Println("Find through 1")
				energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy)...)
				break hall_loop
			}
		}

	room_loop:
		for i, room := range situation.rooms {
			if room[0] == pod.id || room[1] == pod.id {
				if i == pod.t {
					if room[1] == pod.id {
						continue amphipod_loop
					}
					other_pod := find_pod(situation, room[1])
					if other_pod.t == pod.t {
						continue amphipod_loop
					}
				}
				if room[0] != pod.id && room[0] != 0 {
					continue amphipod_loop
				}

				other_pod := find_pod(situation, situation.rooms[pod.t][1])
				pod_hallway_id := 2 + i*2
				base_energy := pod.energy
				if room[1] == pod.id {
					base_energy += pod.energy
				}

				if other_pod.t == -1 || (situation.rooms[pod.t][0] == 0 && other_pod.t == pod.t) {
					room_hallway_id := 2 + pod.t*2
					distance := get_hallway_distance(situation, pod_hallway_id, room_hallway_id, pod.id)
					if distance >= 0 {
						pod_energy := distance*pod.energy + base_energy
						new_situation := deep_copy_board(situation)

						if room[1] == pod.id {
							new_situation.rooms[i][1] = 0
						} else {
							new_situation.rooms[i][0] = 0
						}

						if situation.rooms[pod.t][1] == 0 {
							pod_energy += pod.energy * 2
							new_situation.rooms[pod.t][1] = pod.id
						} else {
							pod_energy += pod.energy
							new_situation.rooms[pod.t][0] = pod.id
						}
						// fmt.Println("Find through 2")
						energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy)...)
						break room_loop
					}
				}

				for j := 0; j < len(situation.hall); j++ {
					if j != 0 && j != 10 && j%2 == 0 {
						continue
					}

					distance := get_hallway_distance(situation, pod_hallway_id, j, pod.id)
					if distance >= 0 {
						pod_energy := distance*pod.energy + base_energy
						new_situation := deep_copy_board(situation)

						new_situation.hall[j] = pod.id
						// fmt.Printf("%+v, %+v, i: %d, energy spent: %d\n", pod, new_situation.rooms[i], i, energy_spent)
						if room[1] == pod.id {
							new_situation.rooms[i][1] = 0
						} else {
							new_situation.rooms[i][0] = 0
						}

						// fmt.Println("Find through 3")
						energy_levels = append(energy_levels, find_done_states(new_situation, energy_spent+pod_energy)...)
					}
				}
				break room_loop
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
	energy_map := [4]int{1, 10, 100, 1000}

	situation := Board{hall: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, rooms: [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}}, pods: []Amphipod{}}

	for scanner.Scan() {
		line := r.FindStringSubmatch(scanner.Text())

		if len(line) > 1 {
			for i, amphi := range line {
				if i == 0 {
					continue
				}
				t := determain_type(amphi)
				total := sum(secs[:]) + 1
				j := 0
				if total > 4 {
					j = 1
				}

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

	done_states := find_done_states(situation, 0)

	// print_situation(situation)
	// fmt.Println(done_states)
	least := math.MaxInt
	for _, ds := range done_states {
		if ds < least {
			least = ds
		}
	}
	fmt.Println(least)
}
