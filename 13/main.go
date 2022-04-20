package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input(fname string) (ts int, sids []string) {
	f, _ := os.Open(fname)
	defer f.Close()
	var si []string
	for s := bufio.NewScanner(f); s.Scan(); {
		si = append(si, s.Text())
	}
	ts, _ = strconv.Atoi(si[0])
	sids = strings.Split(si[1], ",")
	return
}

func parse1(sids []string) (ids []int) {
	for _, s := range sids {
		if id, err := strconv.Atoi(s); err == nil {
			ids = append(ids, id)
		}
	}
	return
}

func parse2(sids []string) (busses map[int]int) {
	busses = make(map[int]int)
	for i, s := range sids {
		if id, err := strconv.Atoi(s); err == nil {
			busses[i] = id
		}
	}
	return
}

func solve1(ts int, ids []int) int {
	var times []int
	for _, id := range ids {
		if ts%id == 0 {
			times = append(times, id)
		} else {
			offset := id - (ts % id)
			times = append(times, ts+offset)
		}
	}
	min := times[0]
	index := 0
	for i := 0; i < len(times); i++ {
		if times[i] < min {
			min = times[i]
			index = i
		}
	}
	return ids[index] * (min - ts)
}

func solve2(busses map[int]int) int {
	ts := 0
	step := 1
	for offset, id := range busses {
		for {
			ts += step
			if (ts+offset)%id == 0 {
				step *= id
				break
			}
		}
	}
	return ts
}

func main() {
	ts, sids := input("input.txt")
	ids := parse1(sids)
	busses := parse2(sids)
	fmt.Println("Part one:", solve1(ts, ids))
	fmt.Println("Part two:", solve2(busses))
}
