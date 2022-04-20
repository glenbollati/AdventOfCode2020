package main

import (
	aoc "aoc/utils"
	"fmt"
	"sort"
	"strings"
)

var ss []string
var aMap map[string][]string // allergen : ingredients
var allIngredients []string
var mappedIngredients []string
var safeIngredients []string

func addFood(s string) {
	spl := strings.Split(strings.ReplaceAll(s, ")", ""), "(contains")
	allergens := strings.Split(strings.Trim(spl[1], " "), ", ")
	ingredients := strings.Split(strings.Trim(spl[0], " "), " ")
	for _, i := range ingredients {
		allIngredients = aoc.SSAppendUnique(allIngredients, i)
	}
	for _, a := range allergens {
		if len(aMap[a]) == 0 {
			aMap[a] = append(aMap[a], ingredients...)
			continue
		}
		aMap[a] = SSIntersect(aMap[a], ingredients)
	}
}

func main() {
	ss = aoc.GetLines("input.txt")
	aMap = make(map[string][]string)
	for _, s := range ss {
		addFood(s)
	}
	for _, as := range aMap {
		for _, a := range as {
			mappedIngredients = aoc.SSAppendUnique(mappedIngredients, a)
		}
	}
	// find the ingredients that are not a risk
	for _, igr := range allIngredients {
		if !aoc.SSContains(mappedIngredients, igr) {
			safeIngredients = append(safeIngredients, igr)
		}
	}
	//p1()
	p2()
}

func p1() {
	acc := 0
	for _, s := range ss {
		spl := strings.Split(strings.ReplaceAll(s, ")", ""), "(contains")
		ingredients := strings.Split(strings.Trim(spl[0], " "), " ")
		acc += len(SSIntersect(ingredients, safeIngredients))
	}
	fmt.Println(acc)
}

func p2() {
	// Filter out known combinations
	// keep doing this till we match everything
	// e.g.
	// dairy:[mxmxvkd] fish:[mxmxvkd sqjhc] soy:[sqjhc fvjkl]
	// ==> dairy:[mxmxvkd] fish:[sqjhc] soy:[fvjkl]
	for done := false; !done; {
		done = true
		for a, igr := range aMap {
			if len(igr) == 1 {
				for aa, igrr := range aMap {
					if a == aa || len(igrr) == 1 {
						continue
					}
					aMap[aa] = aoc.SSRemove(igrr, igr[0])
				}
				continue
			}
			done = false
		}
	}
	// Sort the results alphabetically by their allergen and separate them by commas
	allergens := []string{}
	for a := range aMap {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)
	for i, a := range allergens {
		if i == 0 {
			// Print without leading comma
			fmt.Printf("%s", aMap[a][0])
			continue
		}
		fmt.Printf(",%s", aMap[a][0])
	}
}

func SSIntersect(sa, sb []string) (si []string) {
	for _, s := range sa {
		if aoc.SSContains(sb, s) {
			si = append(si, s)
		}
	}
	return
}
