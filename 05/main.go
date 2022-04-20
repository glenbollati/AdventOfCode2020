package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Portion int
type grid map[int][]int

const (
	Lower Portion = iota
	Upper
)

type Seat struct {
	code string
	row  int
	col  int
	id   int
}

func sliceContains(s []int, i int) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}

func split(a [2]int, p Portion) (b [2]int) {
	mp := (a[1] + 1 + a[0]) / 2
	switch p {
	case Lower:
		b = [2]int{a[0], mp - 1}
	case Upper:
		b = [2]int{mp, a[1]}
	}
	return
}

func getID(r, c int) int {
	return (r * 8) + c
}

func maxID(ss []Seat) int {
	max := 0
	for _, s := range ss {
		if s.id > max {
			max = s.id
		}
	}
	return max
}

func converge(s string, limits [2]int) int {
	for _, c := range s {
		switch c {
		case 'F', 'L':
			limits = split(limits, Lower)
		case 'B', 'R':
			limits = split(limits, Upper)
		}
		if limits[0] == limits[1] {
			break
		}
	}
	return limits[0]
}

func newSeat(code string) Seat {
	r := converge(code[:7], [2]int{0, 127})
	c := converge(code[7:], [2]int{0, 7})
	id := getID(r, c)
	return Seat{
		code: code,
		row:  r,
		col:  c,
		id:   id,
	}
}

func mySeat(seats []Seat) (m int) {
	// Populate a slice of all possible seat IDs
	var possible []int
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			possible = append(possible, getID(i, j))
		}
	}

	// Populate a slice of all occupied seat IDs
	var occupied []int
	for _, s := range seats {
		occupied = append(occupied, s.id)
	}

	// Find the possible seats that are not occupied
	var empty []int
	for _, p := range possible {
		if !sliceContains(occupied, p) {
			empty = append(empty, p)
		}
	}

	// Find the unoccupied seat for which seats with ID+1 and ID-1
	// are occupied
	for _, s := range empty {
		if sliceContains(occupied, s+1) && sliceContains(occupied, s-1) {
			m = s
		}
	}
	return
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	var seats []Seat
	for s.Scan() {
		seats = append(seats, newSeat(s.Text()))

	}

	fmt.Println("Part one:", maxID(seats))
	fmt.Println("Part two:", mySeat(seats))
}
