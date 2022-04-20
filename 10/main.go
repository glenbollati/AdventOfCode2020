package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	adapts := input()
	adapts = append(adapts, 0)
	sort.IntSlice(adapts).Sort()
	//sort.Slice(adapts, func(i, j int) bool { return adapts[i] < adapts[j] })
	adapts = append(adapts, adapts[len(adapts)-1]+3)

	one, three := joltDiffs(adapts)
	fmt.Println("Part one:", one*three)

	cache := make(map[int]int)
	fmt.Println("Part two:", joltCombos(adapts, 0, cache))
}

func input() (si []int) {
	f, _ := os.Open("input.txt")
	for s := bufio.NewScanner(f); s.Scan(); {
		i, _ := strconv.Atoi(s.Text())
		si = append(si, i)
	}
	return
}

func joltDiffs(adapts []int) (one, three int) {
	for i := 0; i < len(adapts)-1; i++ {
		if adapts[i+1]-adapts[i] == 1 {
			one++
		} else if adapts[i+1]-adapts[i] == 3 {
			three++
		} else {
			log.Fatalf("Found a bug: %d - %d is not 1 or 3!\n", adapts[i+1], adapts[i])
		}
	}
	return
}

func joltCombos(adapts []int, i int, cache map[int]int) int {
	if val, in := cache[i]; in {
		return val
	}
	if i == len(adapts)-1 {
		return 1
	}
	ans := 0
	for j := i + 1; j < len(adapts); j++ {
		if adapts[j]-adapts[i] <= 3 {
			ans += joltCombos(adapts, j, cache)
		} else {
			break
		}
	}
	cache[i] = ans
	return ans
}
