package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input() (ls []int) {
	f, _ := os.Open("input.txt")
	for s := bufio.NewScanner(f); s.Scan(); {
		num, _ := strconv.Atoi(s.Text())
		ls = append(ls, num)
	}
	return
}

func valid(set []int, sum int) bool {
	for _, n := range set {
		for _, nn := range set {
			if n == nn {
				continue
			} else if n+nn == sum {
				return true
			}
		}
	}
	return false
}

func findInvalid(data []int) int {
	for i := 25; i < len(data); i++ {
		toCheck := data[i-25 : i]
		if !valid(toCheck, data[i]) {
			return data[i]
		}
	}
	return -1
}

func sumOfSlice(ss []int) (sum int) {
	for _, i := range ss {
		sum += i
	}
	return
}

func maxOfSlice(ss []int) (max int) {
	for _, i := range ss {
		if i > max {
			max = i
		}
	}
	return
}

func minOfSlice(ss []int) (min int) {
	min = ss[0]
	for _, i := range ss {
		if i > min {
			continue
		}
		min = i
	}
	return
}

func findSum(data []int, target int) (sum int) {
	i := 0
	vals := []int{}
	for sumOfSlice(vals) < target {
		i++
		vals = append(vals, data[i])
		for sumOfSlice(vals) > target {
			ii := 1
			vals = vals[ii:]
			ii++
		}
	}
	return minOfSlice(vals) + maxOfSlice(vals)
}

func main() {
	data := input()
	invalid := findInvalid(data)
	fmt.Println("Part one:", invalid)                // 776203571
	fmt.Println("Part two:", findSum(data, invalid)) // 104800569
}
