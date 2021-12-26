package main

import (
	"fmt"
	"math"
)

func get_x_rotation(angle float64) [3][3]int {
	return [3][3]int{
		[3]int{1, 0, 0},
		[3]int{0, int(math.Cos(float64(angle))), -int(math.Sin(float64(angle)))},
		[3]int{0, int(math.Sin(float64(angle))), int(math.Cos(float64(angle)))},
	}
}

func get_y_rotation(angle float64) [3][3]int {
	return [3][3]int{
		[3]int{int(math.Cos(float64(angle))), 0, int(math.Sin(float64(angle)))},
		[3]int{0, 1, 0},
		[3]int{-int(math.Sin(float64(angle))), 0, int(math.Cos(float64(angle)))},
	}
}

func get_z_rotation(angle float64) [3][3]int {
	return [3][3]int{
		[3]int{int(math.Cos(float64(angle))), -int(math.Sin(float64(angle))), 0},
		[3]int{int(math.Sin(float64(angle))), int(math.Cos(float64(angle))), 0},
		[3]int{0, 0, 1},
	}
}

func multiply(m1, m2 [3][3]int) [3][3]int {
	new_m := [3][3]int{}

	for i := 0; i < len(m1); i++ {
		new_row := [3]int{0, 0, 0}
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m1[0]); k++ {
				new_row[j] += m1[i][k] * m2[j][k]
			}
		}
		new_m[i] = new_row
	}

	return new_m
}

func main() {
	for x := 0.0; x < 4.0; x++ {
		for y := 0.0; y < 4; y++ {
			for z := 0.0; z < 4; z++ {
				xy := multiply(get_x_rotation(math.Pi*(x/2.0)), get_y_rotation(math.Pi*(y/2.0)))
				xyz := multiply(xy, get_z_rotation(math.Pi*(z/2.0)))
				fmt.Println(xyz)
			}
		}
	}
}
