package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passValidPartOne(l string) bool {
	ss := strings.Split(l, " ")
	sh := strings.Split(ss[0], "-")
	min, err := strconv.Atoi(sh[0])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(sh[1])
	if err != nil {
		log.Fatal(err)
	}
	char := strings.Trim(ss[1], ":")
	pass := ss[2]

	count := strings.Count(pass, char)
	if count < min || count > max {
		return false
	}
	return true
}

func passValidPartTwo(l string) bool {
	ss := strings.Split(l, " ")
	sh := strings.Split(ss[0], "-")
	pos1, err := strconv.Atoi(sh[0])
	if err != nil {
		log.Fatal(err)
	}
	pos2, err := strconv.Atoi(sh[1])
	if err != nil {
		log.Fatal(err)
	}
	char := strings.Trim(ss[1], ":")
	pass := ss[2]

	p1 := string(pass[pos1-1])
	p2 := string(pass[pos2-1])

	if p1 == char && p2 == char {
		return false
	}
	if p1 != char && p2 != char {
		return false
	}
	return true
}

func linesFromFile(fname string) (lines []string) {
	f, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(f), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return
}

func main() {
	ls := linesFromFile("input.txt")
	var validPassesOne []string
	for _, l := range ls {
		if passValidPartOne(l) {
			validPassesOne = append(validPassesOne, l)
		}
	}
	fmt.Println("Valid passwords per policy one:", len(validPassesOne))

	var validPassesTwo []string
	for _, l := range ls {
		if passValidPartTwo(l) {
			validPassesTwo = append(validPassesTwo, l)
		}
	}
	fmt.Println("Valid passwords per policy two:", len(validPassesTwo))
}
