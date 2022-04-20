package main

import "fmt"

const (
	p1Turns = 2020
	p2Turns = 30000000
)

var input = []int{0, 12, 6, 13, 20, 1, 17}
var testInput = []int{0, 3, 6}

func solve(numbers []int, numTurns int) int {
	// Map of number to turn nums in which it was spoken
	numMem := make(map[int][]int)
	for i, n := range numbers {
		numMem[n] = []int{i + 1}
		// DEBUG fmt.Println(i+1, ":", n)
	}
	// DEBUG fmt.Println(numMem)
	currNum := numbers[len(numbers)-1]
	for turn := len(numbers) + 1; turn <= numTurns; turn++ {
		if len(numMem[currNum]) == 1 {
			currNum = 0
		} else {
			spoken := numMem[currNum]
			currNum = spoken[len(spoken)-1] - spoken[len(spoken)-2]
		}
		numMem[currNum] = append(numMem[currNum], turn)
		// DEBUG fmt.Println(turn, ":", currNum)
	}
	return currNum
}

func main() {
	fmt.Println("Part one:", solve(input, p1Turns)) // 620
	fmt.Println("Part two:", solve(input, p2Turns)) // 110871
}
