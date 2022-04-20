package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func linesFromFile(fname string) (ls []string) {
	f, _ := os.Open(fname)
	defer f.Close()
	for s := bufio.NewScanner(f); s.Scan(); {
		ls = append(ls, s.Text())
	}
	return
}

func flippable(instrs []string) (toFlip []int) {
	for i, r := range instrs {
		if strings.HasPrefix(r, "jmp") || strings.HasPrefix(r, "nop") {
			toFlip = append(toFlip, i)
		}
	}
	return
}

func splitInstruction(s string) (string, int) {
	split := strings.Split(s, " ")
	cmd := split[0]
	val, _ := strconv.Atoi(split[1])
	return cmd, val
}

func runProgram(instrs []string, flip int) (acc int, success bool) {
	visited := map[int]bool{}
	i := 0
	for i < len(instrs) {
		if visited[i] {
			return acc, false
		}
		visited[i] = true

		cmd, val := splitInstruction(instrs[i])
		if i == flip {
			cmd = switchCommands(cmd)
		}
		switch cmd {
		case "jmp":
			i += val
			continue
		case "acc":
			acc += val
		}
		i++
	}
	return acc, true
}

func switchCommands(cmd string) string {
	if cmd == "nop" {
		return "jmp"
	} else if cmd == "jmp" {
		return "nop"
	}
	return cmd
}

func main() {
	instrs := linesFromFile("input.txt")
	part1, _ := runProgram(instrs, -1)
	fmt.Println("Part one:", part1)

	for _, i := range flippable(instrs) {
		part2, success := runProgram(instrs, i)
		if success {
			fmt.Println("Part two:", part2)
		}
	}
}
