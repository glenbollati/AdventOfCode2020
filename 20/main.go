package main

import (
	aoc "aoc/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	HPLANE = iota
	VPLANE
	CW
	CCW

	BT = 0
	BR = 1
	BB = 2
	BL = 3
)

var (
	globalCornerTiles []int
	cTile             Tile
	tiles             map[int]*Tile
	bs                map[string][]int
)

type Tile struct {
	full    []string
	borders [4]string
}

func (t *Tile) SetBorders() {
	t.borders[BT] = t.full[0]
	t.borders[BB] = t.full[len(t.full)-1]
	var left, right string
	for _, row := range t.full {
		left += string(row[0])
		right += string(row[len(row)-1])
	}
	t.borders[BL] = left
	t.borders[BR] = right
}

func (t *Tile) TrimBorders() {
	var top, right, bot, left string

	EOT := len(t.full) - 1

	top = t.full[1][1:EOT]
	bot = t.full[len(t.full)-2][1:EOT]

	for _, row := range t.full[1:EOT] {
		left += string(row[1])
		right += string(row[len(row)-2])
	}
	t.borders[BT] = top
	t.borders[BB] = bot
	t.borders[BL] = left
	t.borders[BR] = right
}

func (t *Tile) PrintBorders() {
	fmt.Println(t.borders[BT])
	for i := 1; i < len(t.borders[BR])-1; i++ {
		fmt.Print(string(t.borders[BL][i]))
		for j := 1; j < len(t.borders[BR])-1; j++ {
			fmt.Print(" ")
		}
		fmt.Println(string(t.borders[BR][i]))
	}
	fmt.Println(t.borders[BB])
}

func (t Tile) String() string {
	s := t.borders[BT] + "\n"
	for i := 1; i < len(t.borders[BR])-1; i++ {
		s += string(t.borders[BR][i])
		s += "\n"
	}
	s += t.borders[BB] + "\n"
	return s
}

func loadTiles(fname string) {
	f, _ := os.Open(fname)
	defer f.Close()

	s := bufio.NewScanner(f)
	tiles = make(map[int]*Tile)
	var buff []string
	var tilenum int

	for s.Scan() {
		// empty line
		if s.Text() == "" {
			tiles[tilenum] = &Tile{full: buff}
			tiles[tilenum].SetBorders()
			continue
		}
		// a line in a tile
		if !strings.Contains(s.Text(), "Tile") {
			buff = append(buff, s.Text())
			continue
		}
		// the tile number
		tilenum = aoc.ToInt(strings.ReplaceAll(strings.Split(s.Text(), " ")[1], ":", ""))
		buff = []string{}
	}
	tiles[tilenum] = &Tile{full: buff}
	tiles[tilenum].SetBorders()

	if err := s.Err(); err != nil {
		panic(err)
	}
}

// Flips the tile by 180 deg
func (t *Tile) Flip(plane int) (flipped Tile) {
	switch plane {
	case HPLANE:
		flipped.borders[BT] = t.borders[BB]
		flipped.borders[BR] = aoc.ReverseStr(t.borders[BR])
		flipped.borders[BB] = t.borders[BT]
		flipped.borders[BL] = aoc.ReverseStr(t.borders[BL])
	case VPLANE:
	}
	return
}

// Rotates the tile by 90 deg CW
func (t *Tile) Rotate() (rotated Tile) {
	rotated.borders[BT] = t.borders[BL]
	rotated.borders[BR] = t.borders[BT]
	rotated.borders[BB] = t.borders[BR]
	rotated.borders[BL] = t.borders[BB]
	return
}

func debug(num int) {
	fmt.Println("Tile: ", num)
	fmt.Println("Flipped on HPLANE")
	fmt.Println(tiles[num].Flip(HPLANE))
}

func p1() {
	// Collect the borders into a map, storing the tile numbers that contain them
	bs = make(map[string][]int)
	for i, t := range tiles {
		for _, b := range t.borders {
			bs[b] = aoc.SIAppendUnique(bs[b], i)
		}
	}
	// Combine borders that are the same but reversed
	for b, _ := range bs {
		if _, hasRev := bs[aoc.ReverseStr(b)]; hasRev {
			for _, t := range bs[aoc.ReverseStr(b)] {
				bs[b] = aoc.SIAppendUnique(bs[b], t)
			}
			delete(bs, aoc.ReverseStr(b))
		}
	}

	// Count the tiles that have 2 isolated borders (a corner) and multiple their IDs together
	mul := 1
	for id, _ := range tiles {
		count := 0
		for _, b := range bs {
			if len(b) == 1 && aoc.SIContains(b, id) {
				count++
			}
		}
		if count == 2 { // corner
			mul *= id
			globalCornerTiles = append(globalCornerTiles, id)
		}
	}
	fmt.Println(mul)
	fmt.Println(globalCornerTiles)
}

func main() {
	loadTiles("input.txt")
	p1()
}
