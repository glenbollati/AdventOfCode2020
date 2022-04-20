package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Run struct {
	movement string
	dx, dy   int
	trees    int
	moves    int
}

type Map struct {
	ls  []string
	pos struct {
		x, y int
	}
}

func newMap(ls []string) *Map {
	m := Map{
		ls: ls,
	}
	return &m
}

func newRun(movement string, x, y int) *Run {
	s := Run{
		movement: movement,
		dx:       x,
		dy:       y,
		trees:    0,
		moves:    0,
	}
	return &s
}

func (m *Map) advance(r *Run) (canMove bool) {
	if m.pos.y+r.dy >= len(m.ls) {
		m.reset()
		return false
	}
	m.pos.y += r.dy
	m.pos.x += r.dx
	if m.pos.x >= len(m.ls[m.pos.y]) {
		m.pos.x = m.pos.x - len(m.ls[m.pos.y])
	}
	r.moves++
	if m.ls[m.pos.y][m.pos.x] == '#' {
		r.trees++
	}
	return true
}

func (m *Map) reset() {
	m.pos.x = 0
	m.pos.y = 0
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var ls []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ls = append(ls, sc.Text())
	}
	m := newMap(ls)

	// Part one
	fmt.Println("--- Part One ---")
	r3d1 := newRun("r3d1", 3, 1)
	for {
		if !m.advance(r3d1) {
			fmt.Println("Total moves:", r3d1.moves)
			fmt.Println("Trees hit:", r3d1.trees)
			break
		}
	}

	// Part two
	r3d1 = newRun("r3d1", 3, 1)
	r1d1 := newRun("r1d1", 1, 1)
	r5d1 := newRun("r5d1", 5, 1)
	r7d1 := newRun("r7d1", 7, 1)
	r1d2 := newRun("r1d2", 1, 2)

	allRun := [5]*Run{
		r3d1, r1d1,
		r5d1, r7d1,
		r1d2,
	}

	fmt.Println("\n--- Part two ---")
	for i := 0; i < len(allRun); i++ {
		for {
			if !m.advance(allRun[i]) {
				fmt.Printf("\n--- %s ---\n", allRun[i].movement)
				fmt.Println("Total moves:", allRun[i].moves)
				fmt.Println("Trees hit:", allRun[i].trees)
				break
			}
		}
	}
	fmt.Println("\n--- Part two answer ---")
	mulTrees := 1
	for i := 0; i < len(allRun); i++ {
		fmt.Print(allRun[i].trees, "*")
		mulTrees *= allRun[i].trees
	}
	fmt.Print(" = ", mulTrees)
}
