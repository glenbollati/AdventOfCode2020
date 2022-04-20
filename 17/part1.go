package main

import (
	"bufio"
	"fmt"
	"os"
)

type Grid map[[3]int]bool

func (g Grid) getMax() (max [3]int) {
	for pos := range g {
		for i := range pos {
			if pos[i] > max[i] {
				max[i] = pos[i]
			}
		}
	}
	return
}

func (g Grid) String() string {
	max := g.getMax()
	var output string
	for z := 0; z <= max[2]; z++ {
		output += fmt.Sprintf("z = %d\n", z)
		for y := 0; y <= max[1]; y++ {
			for x := 0; x <= max[0]; x++ {
				if g[[3]int{x, y, z}] {
					output += "#"
				} else {
					output += "."
				}
			}
			output += "\n"
		}
		output += "\n"
	}
	return output
}

// Expand by 2 along each dimension
func ExpandGrid(oldGrid Grid) (newGrid Grid) {
	max := oldGrid.getMax()
	toAdd := [][3]int{}

	xNew := [2]int{0, max[0] + 3}
	yNew := [2]int{0, max[1] + 3}
	zNew := [2]int{0, max[2] + 3}

	for z := 0; z < zNew[1]; z++ {
		for y := 0; y < yNew[1]; y++ {
			for x := 0; x < xNew[1]; x++ {
				toAdd = append(toAdd, [3]int{x, y, z})
			}
		}
	}

	newGrid = make(Grid)
	for k, v := range oldGrid {
		newPos := [3]int{k[0] + 1, k[1] + 1, k[2] + 1}
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
	y, z := 0, 0
	for s := bufio.NewScanner(f); s.Scan(); {
		for x := range s.Text() {
			active := false
			if s.Text()[x] == '#' {
				active = true
			}
			grid[[3]int{x, y, z}] = active
		}
		y++
	}
	return
}

func inRange(pos, from [3]int) bool {
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

func (g Grid) activeNearby(from [3]int) (count int) {
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
		// DEBUG fmt.Printf("Cycle %d\n", cycle)
		var toFlip [][3]int
		grid = ExpandGrid(grid)
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
	grid := loadInput("input.txt") // 211
	fmt.Println("Part one:", runCycles(grid, 6))
}
