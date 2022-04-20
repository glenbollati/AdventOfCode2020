package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	solve(true)
	solve(false)
}

func countOccupied(ss [][]string) (occ int) {
	for r := 0; r < len(ss); r++ {
		for c := 0; c < len(ss[0]); c++ {
			if ss[r][c] == "#" {
				occ++
			}
		}
	}
	return
}

func solve(p1 bool) {
	s := getSeats()
	var change bool
	for {
		s, change = step(s, p1)
		if !change {
			break
		}
	}
	if p1 {
		fmt.Println("Part one:", countOccupied(s))
	} else {
		fmt.Println("Part two:", countOccupied(s))
	}
}

func getSeats() (ss [][]string) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	for s := bufio.NewScanner(f); s.Scan(); {
		ss = append(ss, strings.Split(s.Text(), ""))
	}
	return
}

func printSeats(ss [][]string) {
	for _, r := range ss {
		for _, state := range r {
			print(state)
		}
		println()
	}
	println()
}

func step(prev [][]string, p1 bool) (ret [][]string, change bool) {
	for r := 0; r < len(prev); r++ {
		var row []string
		for c := 0; c < len(prev[0]); c++ {
			count := 0
			if p1 {
				// PART 1
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						rr := r + dr
						cc := c + dc
						if rr < 0 || rr >= len(prev) ||
							cc < 0 || cc >= len(prev[0]) ||
							(dr == 0 && dc == 0) {
							continue
						}
						if prev[rr][cc] == "#" {
							count++
						}
					}
				}
			} else {
				// PART 2
				dirs := [8][2]int{
					{0, -1},
					{1, -1},
					{1, 0},
					{1, 1},
					{0, 1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
				}
				for _, d := range dirs {
					rr := r
					cc := c
					for {
						rr += d[0]
						cc += d[1]
						if rr < 0 || rr >= len(prev) ||
							cc < 0 || cc >= len(prev[0]) {
							break
						}
						if prev[rr][cc] == "." {
							continue
						}
						if prev[rr][cc] == "L" {
							break
						}
						if prev[rr][cc] == "#" {
							count++
							break
						}
					}
				}
			}
			newCol := prev[r][c]
			if prev[r][c] == "L" && count == 0 {
				newCol = "#"
				change = true
			} else if prev[r][c] == "#" {
				if (p1 && count >= 4) || (!p1 && count >= 5) {
					newCol = "L"
					change = true
				}
			}
			row = append(row, newCol)
		}
		ret = append(ret, row)
	}
	return
}
