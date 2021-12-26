package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	is_operator bool
	packets     []Packet
	value       uint64
	version     uint64
	type_id     uint64
}

func if_err(err error) {
	if err != nil {
		panic(err)
	}
}

func read_bits(bit_string []uint64, n int) uint64 {
	result := uint64(0)
	for i := 0; i < n; i++ {
		result = result << 1
		result += bit_string[i]
	}
	return result
}

func parse_bit_string(bit_string []uint64) (Packet, int) {
	v := read_bits(bit_string, 3)
	t := read_bits(bit_string[3:], 3)

	position := 6
	new_packet := Packet{
		version: v,
		type_id: t,
	}

	if t == 4 {
		result := []uint64{}
		continue_reading := true
		for ok := true; ok; ok = continue_reading {
			result = append(result, bit_string[position+1:position+5]...)
			continue_reading = read_bits(bit_string[position:], 1) == 1
			position += 5
		}
		new_packet.value = read_bits(result, len(result))
	} else {
		new_packet.is_operator = true
		is_length_operator := read_bits(bit_string[position:], 1) == 0
		position++

		if is_length_operator {
			l := int(read_bits(bit_string[position:], 15))
			position += 15
			total_read := 0

			for ok := true; ok; ok = read_bits(bit_string[position+total_read:], l-total_read) > 0 {
				sub_packet, read := parse_bit_string(bit_string[position+total_read:])
				new_packet.packets = append(new_packet.packets, sub_packet)
				total_read += read
			}
			position += l
		} else {
			l := int(read_bits(bit_string[position:], 11))
			position += 11
			total_read := 0

			for i := 0; i < l; i++ {
				sub_packet, read := parse_bit_string(bit_string[position+total_read:])
				new_packet.packets = append(new_packet.packets, sub_packet)
				total_read += read
			}
			position += total_read
		}
	}

	return new_packet, position
}

func sum_versions(pack Packet) uint64 {
	sum := pack.version
	for _, p := range pack.packets {
		sum += sum_versions(p)
	}
	return sum
}

func asBits(h string) []uint64 {
	val, err := strconv.ParseUint(h, 16, 32)
	if_err(err)

	bits := []uint64{}
	for i := 0; i < 4; i++ {
		bits = append([]uint64{val & 0x1}, bits...)
		// or
		// bits = append(bits, val & 0x1)
		// depending on the order you want
		val = val >> 1
	}
	return bits
}

func main() {
	file, err := os.Open("./input.txt")
	// file, err := os.Open("./sample-input.txt")
	if_err(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	message := []uint64{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		for _, val := range line {
			message = append(message, asBits(val)...)
		}
	}

	// fmt.Println(message)

	// fmt.Println(read_bits(message[4:], 3))
	packets, _ := parse_bit_string(message)
	// fmt.Printf("%+v\nread: %d\n", packets, read)
	fmt.Println(sum_versions(packets))
}
