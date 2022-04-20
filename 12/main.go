package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ship struct {
	facing     int
	posN, posE int
	wp         Waypoint
}

type Waypoint struct {
	posN, posE int
}

var ship *Ship

func newShip() *Ship {
	s := Ship{
		facing: 90, // N: 0, E: 90, S: 180, W: 270
		posN:   0,
		posE:   0,
		wp: Waypoint{
			posN: 1,
			posE: 10,
		},
	}
	return &s
}

func (s *Ship) Debug() {
	fmt.Printf("Facing: %d, posN: %d, posE: %d WP: posN: %d, posE: %d\n",
		s.facing, s.posN, s.posE, s.wp.posN, s.wp.posE)
}

func (s *Ship) Forward(n int) {
	switch s.facing {
	case 0:
		s.posN += n
	case 90:
		s.posE += n
	case 180:
		s.posN -= n
	case 270:
		s.posE -= n
	}
}

func (s *Ship) Move(dir string, n int) {
	switch dir {
	case "N":
		s.posN += n
	case "E":
		s.posE += n
	case "S":
		s.posN -= n
	case "W":
		s.posE -= n
	}
}

func (s *Ship) MHDist() int {
	if s.posN < 0 {
		s.posN *= -1
	}
	if s.posE < 0 {
		s.posE *= -1
	}
	return s.posN + s.posE
}

func (s *Ship) Rotate(dir string, n int) {
	switch dir {
	case "R":
		s.facing += n
	case "L":
		s.facing -= n
	}
	for s.facing >= 360 {
		s.facing -= 360
	}
	for s.facing < 0 {
		s.facing += 360
	}
}

func (s *Ship) ToWp(n int) {
	s.posN = s.posN + (s.wp.posN * n)
	s.posE = s.posE + (s.wp.posE * n)
}

func (s *Ship) MoveWp(dir string, n int) {
	switch dir {
	case "N":
		s.wp.posN += n
	case "E":
		s.wp.posE += n
	case "S":
		s.wp.posN -= n
	case "W":
		s.wp.posE -= n
	}
}

/*************************************
	Rotating a point (x, y) by 90deg:
	CW yields (y, -x)
	CCW yields (-y, x)
**************************************/
func (s *Ship) RotateWp(dir string, n int) {
	switch dir {
	case "R":
		for i := 0; i < (n / 90); i++ {
			newPos := [2]int{-s.wp.posE, s.wp.posN}
			s.wp.posE = newPos[1]
			s.wp.posN = newPos[0]
		}
	case "L":
		for i := 0; i < (n / 90); i++ {
			newPos := [2]int{s.wp.posE, -s.wp.posN}
			s.wp.posE = newPos[1]
			s.wp.posN = newPos[0]
		}
	}
}

func exec(instrs []string, p1 bool) {
	for _, instr := range instrs {
		av := strings.SplitN(instr, "", 2)
		a := av[0]
		v, err := strconv.Atoi(av[1])
		if err != nil {
			log.Fatal(err)
		}
		if p1 {
			switch a {
			case "N", "E", "S", "W":
				ship.Move(a, v)
			case "R", "L":
				ship.Rotate(a, v)
			case "F":
				ship.Forward(v)
			}
		} else {
			switch a {
			case "N", "E", "S", "W":
				ship.MoveWp(a, v)
			case "R", "L":
				ship.RotateWp(a, v)
			case "F":
				ship.ToWp(v)
			}
		}
	}
}

func main() {
	ship = newShip()
	instrs := input()
	exec(instrs, true)
	fmt.Println("Part one:", ship.MHDist()) // 1457
	ship = newShip()
	exec(instrs, false)
	fmt.Println("Part two:", ship.MHDist()) //106860
}

func input() (ni []string) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	for s := bufio.NewScanner(f); s.Scan(); {
		ni = append(ni, s.Text())
	}
	return ni
}
