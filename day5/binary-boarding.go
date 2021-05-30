package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	max := 0
	var seats []int

	s := bufio.NewScanner(f)
	for s.Scan() {
		temp := CalcSeatID(s.Text())
		if temp > max {
			max = temp
		}
		seats = append(seats, temp)
	}

	// Part 1
	fmt.Println(max)

	// Part 2
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	fmt.Println(MySeat(seats))
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func MySeat(seats []int) int {
	for _, seatA := range seats[1 : len(seats)-1] {
		for _, seatB := range seats[1 : len(seats)-1] {
			diff := seatB - seatA
			if diff < 0 {
				diff = diff * -1
			}
			inBetween := (seatB + seatA) / 2
			if diff == 2 && !contains(seats, inBetween) {
				return inBetween
			}
		}
	}
	return 0
}

func ColSel(seat byte, start *int, end *int) {
	if seat == byte('F') {
		*end = (*start + *end) / 2
	} else {
		*start = ((*start + *end) / 2) + 1
	}
}

func RowSel(seat byte, start *int, end *int) {
	if seat == byte('L') {
		*end = (*start + *end) / 2
	} else {
		*start = ((*start + *end) / 2) + 1
	}
}

func CalcSeatID(seat string) int {
	var start, end, row, col int
	start, end, row, col = 0, 127, 0, 0

	ColSel(seat[0], &start, &end)
	ColSel(seat[1], &start, &end)
	ColSel(seat[2], &start, &end)
	ColSel(seat[3], &start, &end)
	ColSel(seat[4], &start, &end)
	ColSel(seat[5], &start, &end)
	if seat[6] == 'F' {
		row = start
	} else {
		row = end
	}

	start, end = 0, 7

	RowSel(seat[7], &start, &end)
	RowSel(seat[8], &start, &end)
	if seat[9] == 'L' {
		col = start
	} else {
		col = end
	}

	// fmt.Println(row, col)

	return row*8 + col
}
