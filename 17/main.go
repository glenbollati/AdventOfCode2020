package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid map[[4]int]bool

func (g Grid) getMax() (max [4]int) {
	for pos := range g {
		for i := range pos {
			if pos[i] > max[i] {
				max[i] = pos[i]
			}
		}
	}
	return
}

// Expand by 2 along each dimension
func ExpandGrid(oldGrid Grid) (newGrid Grid) {
	max := oldGrid.getMax()
	toAdd := [][4]int{}

	xMax := max[0] + 3
	yMax := max[1] + 3
	zMax := max[2] + 3
	wMax := max[3] + 3

	for w := 0; w < wMax; w++ {
		for z := 0; z < zMax; z++ {
			for y := 0; y < yMax; y++ {
				for x := 0; x < xMax; x++ {
					toAdd = append(toAdd, [4]int{x, y, z, w})
				}
			}
		}
	}

	newGrid = make(Grid)
	for k, v := range oldGrid {
		newPos := [4]int{k[0] + 1, k[1] + 1, k[2] + 1, k[3] + 1}
		newGrid[newPos] = v
	}
	for _, pos := range toAdd {
		if _, in := newGrid[pos]; in {
			continue
		}
		newGrid[pos] = false
	}
	return
}

func loadInput(fname string) (grid Grid) {
	grid = make(Grid)
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	y, z, w := 0, 0, 0
	for s := bufio.NewScanner(f); s.Scan(); {
		for x := range s.Text() {
			active := false
			if s.Text()[x] == '#' {
				active = true
			}
			grid[[4]int{x, y, z, w}] = active
		}
		y++
	}
	return
}

func inRange(pos, from [4]int) bool {
	if pos == from {
		return false
	}
	for i := range pos {
		if pos[i]-from[i] > 1 || pos[i]-from[i] < -1 {
			return false
		}
	}
	return true
}

func (g Grid) activeNearby(from [4]int) (count int) {
	for pos, active := range g {
		if inRange(pos, from) && active {
			count++
		}
	}
	return
}

func runCycles(grid Grid, num int) (activeCount int) {
	// DEBUG fmt.Println(grid)
	for cycle := 0; cycle < num; cycle++ {
		// DEBUG
		fmt.Printf("Cycle %d\n", cycle)
		var toFlip [][4]int
		grid = ExpandGrid(grid)
		fmt.Printf("Grid expanded\n")
		for pos, active := range grid {
			ac := grid.activeNearby(pos)
			if (active && !(ac == 2 || ac == 3)) || (!active && ac == 3) {
				toFlip = append(toFlip, pos)
			}
		}
		for _, pos := range toFlip {
			grid[pos] = !grid[pos]
		}
		// DEBUG fmt.Println(grid)
	}
	for _, active := range grid {
		if active {
			activeCount++
		}
	}
	return
}

func main() {
	grid := loadInput("input.txt") // Part one: 211, Part two: 1952
	fmt.Println("Part two:", runCycles(grid, 6))
}
