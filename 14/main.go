package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mask struct {
	toSet    uint64 // Part 1 only
	values   uint64 // Shared
	floating uint64 // Part 2 only
}

func (m *Mask) ApplyV1(b uint64) uint64 {
	return (b &^ m.toSet) | (m.values & m.toSet)
}

func (mask *Mask) ApplyV2(addr uint64) (combos []uint64) {
	addr = addr | mask.values
	m := mask.floating
	for combos = append(combos, m|(addr&^mask.floating)); m > 0; {
		m = (m - 1) & mask.floating
		combos = append(combos, m|(addr&^mask.floating))
	}
	return
}

func input(fname string) (si []string) {
	f, _ := os.Open(fname)
	for s := bufio.NewScanner(f); s.Scan(); {
		si = append(si, s.Text())
	}
	return
}

func splitInstr(s string) (uint64, uint64) {
	var addr, val uint64
	if _, err := fmt.Sscanf(s, "mem[%d] = %d", &addr, &val); err != nil {
		panic(err)
	}
	return addr, val
}

func toUint(s string) uint64 {
	num, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func getMask(s string) (mask Mask) {
	// V1
	var strMask string
	if _, err := fmt.Sscanf(s, "mask = %s", &strMask); err != nil {
		panic(err)
	}
	// These are the bits to set, so every 0 and 1 becomes 1 and every X becomes 0
	toSetRep := strings.NewReplacer("0", "1", "X", "0")
	toSet := toSetRep.Replace(strMask)

	// In part 1, these are the values of the bits to set,
	// so the values of the Xs don't matter.
	// We set them to 0 to reuse in part2
	values := strings.Replace(strMask, "X", "0", -1)

	floatingRep := strings.NewReplacer("1", "0", "X", "1")
	floating := floatingRep.Replace(strMask)

	mask.toSet = toUint(toSet)
	mask.values = toUint(values)
	mask.floating = toUint(floating)
	return mask
}

func solve(si []string, part2 bool) (sum uint64) {
	mem := make(map[uint64]uint64)
	var currMask Mask
	for _, s := range si {
		if strings.HasPrefix(s, "mask = ") {
			currMask = getMask(s)
			continue
		}
		addr, val := splitInstr(s)
		if part2 {
			addrs := currMask.ApplyV2(addr)
			for _, a := range addrs {
				mem[a] = val
			}
			continue
		} else {
			mem[addr] = currMask.ApplyV1(val)
		}
	}
	for _, v := range mem {
		sum += v
	}
	return
}

func main() {
	si := input("input.txt")
	fmt.Println("Part one:", solve(si, false)) // 18630548206046
	fmt.Println("Part two:", solve(si, true))  // 4254673508445
}
