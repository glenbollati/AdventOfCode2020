package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func groups(fname string) []string {
	f, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	gs := strings.Split(string(f), "\n\n")
	gs[len(gs)-1] = strings.TrimSuffix(gs[len(gs)-1], "\n")
	return gs
}

func countUnique(g string) (n int) {
	g = strings.ReplaceAll(g, "\n", "")
	var met string
	for _, c := range g {
		if !strings.Contains(met, string(c)) {
			met += string(c)
			n++
		}
	}
	return n
}

func countCommon(g string) (n int) {
	ps := strings.Split(g, "\n")
	gg := strings.ReplaceAll(g, "\n", "")

	var summed string
	for _, rc := range gg {
		c := string(rc)
		if strings.Count(gg, c) == len(ps) && !strings.Contains(summed, c) {
			summed += c
			n++
		}
	}
	return n
}

func main() {
	gs := groups("input.txt")

	unique := 0
	for _, g := range gs {
		unique += countUnique(g)
	}

	common := 0
	for _, g := range gs {
		common += countCommon(g)
	}
	fmt.Println("Part one:", unique)
	fmt.Println("Part two:", common)
}
