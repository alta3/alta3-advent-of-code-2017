package main

import (
	"math"
	"fmt"
)

func main() {
	input := 265149
	side, nsq := FindTheUpperSquareRoot(input)
	center := (side - 1)/2

	diff := nsq - input
	i := 1
	for diff > side  {
		diff = diff - side
		i += 1
	}

	premath := diff - (center)
	fhwalks := math.Abs(float64(premath))
	hwalks := int(fhwalks)


	p1 := center + hwalks
	fmt.Println(p1)
	p2 := NextLargestSpiralGeneration(input)
	fmt.Println(p2)
}

func FindTheUpperSquareRoot(num int) (side, nsq int) {
	sq := math.Sqrt(float64(num))
	fside := math.Ceil(sq)
	side = int(fside)
	if side % 2 == 0 {
		side += 1
	}
	nsq = side * side
	return
}

func NextLargestSpiralGeneration(num int) (int) {
	array := [70][70]int{}
	i := 1
	x := 55
	y := 55

	array[x][y] = 1

	for i < 100 {
		x, y = move_right(i, x, y, &array)
		x, y = move_up(i, x, y, &array)
		i += 1
		x, y = move_left(i, x, y, &array)
		x, y = move_down(i, x, y, &array)
		i += 1
	}

	return 1000000
}

func get_sum_from_around(x int, y int, array *[70][70]int) (int) {
	v1 := array[x-1][y]
	v2 := array[x+1][y]
	v3 := array[x][y-1]
	v4 := array[x][y+1]
	v5 := array[x-1][y-1]
	v6 := array[x-1][y+1]
	v7 := array[x+1][y+1]
	v8 := array[x+1][y-1]
	sum := v1 + v2 + v3 + v4 + v5 + v6 + v7 + v8
	return sum
}


func move_right(steps int, x int, y int, array *[70][70]int) (int, int){
	for steps >= 1 {
		nv := get_sum_from_around(x+1, y, array)
		x += 1
		(*array)[x][y] = nv
		steps = steps - 1
	}
	return x, y
}

func move_up(steps int, x int, y int, array *[70][70]int) (int, int) {
	for steps >= 1 {
		nv := get_sum_from_around(x, y+1, array)
		y += 1
		(*array)[x][y] = nv
		steps = steps - 1
	}
	return x, y
}
func move_left(steps int, x int, y int, array *[70][70]int) (int, int) {
	for steps >= 1 {
		nv := get_sum_from_around(x-1, y, array)
		x = x - 1
		(*array)[x][y] = nv
		steps = steps - 1
	}
	return x, y
}
func move_down(steps int, x int, y int, array *[70][70]int) (int, int) {
	for steps >= 1 {
		nv := get_sum_from_around(x, y-1, array)
		y = y - 1
		(*array)[x][y] = nv
		steps = steps - 1
	}
	return x, y
}
