package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntSet map[int]struct{}
type Ticket []int

var allFields map[string][]int
var myTicket Ticket
var nearbyTickets []Ticket

func (s IntSet) Contains(t int) bool {
	_, in := s[t]
	return in
}

func newTicket(s string) (t Ticket) {
	toConv := strings.Split(s, ",")
	for _, c := range toConv {
		val, _ := strconv.Atoi(c)
		t = append(t, val)
	}
	return
}

func addField(s string) {
	split := strings.Split(s, ":")
	ranges := strings.Split(split[1], "or")
	values := []int{}
	for _, r := range ranges {
		expanded := ExpandRange(strings.TrimSpace(r))
		for _, e := range expanded {
			values = append(values, e)
		}
	}
	allFields[split[0]] = values
}

func AllValid() (valid IntSet) {
	valid = make(IntSet)
	for _, vs := range allFields {
		for _, v := range vs {
			valid[v] = struct{}{}
		}
	}
	return
}

func ExpandRange(s string) (r []int) {
	split := strings.Split(s, "-")
	min, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])
	for i := min; i <= max; i++ {
		r = append(r, i)
	}
	return
}

func loadInput(fname string) {
	f, _ := os.Open(fname)
	defer f.Close()
	s := bufio.NewScanner(f)
	allFields = make(map[string][]int)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		if s.Text() == "your ticket:" {
			break
		}
		addField(s.Text())
	}
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		if s.Text() == "nearby tickets:" {
			break
		}
		myTicket = newTicket(s.Text())
	}
	for s.Scan() {
		if s.Text() == "nearby tickets:" {
			continue
		}
		nearbyTickets = append(nearbyTickets, newTicket(s.Text()))
	}
}

func solve1() (sum int, validTickets []Ticket) {
	validValues := AllValid()
	for _, ticket := range nearbyTickets {
		valid := true
		for _, v := range ticket {
			if !validValues.Contains(v) {
				sum += v
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return
}

func sliceContains(si []int, v int) bool {
	for _, s := range si {
		if s == v {
			return true
		}
	}
	return false
}

func matchingFields(value int) (fields []string) {
	for k, v := range allFields {
		if sliceContains(v, value) {
			fields = append(fields, k)
		}
	}
	return
}

func solve2(validTickets []Ticket) (depMult int64) {
	// Map each field number to all possible field names (for each ticket)
	// keeping a count of how many times each field number is mapped to
	// a field name
	fieldMap := make(map[int]map[string]int)
	for _, ticket := range validTickets {
		for field, value := range ticket {
			if _, ok := fieldMap[field]; !ok {
				fieldMap[field] = make(map[string]int)
			}
			matches := matchingFields(value)
			for _, m := range matches {
				fieldMap[field][m] += 1
			}
		}
	}
	// For each field, remove field names that do not match for ALL tickets
	// this is where the count above comes in handy
	for num, nameCount := range fieldMap {
		for name, count := range nameCount {
			//if count < len(fieldMap) {
			if count < len(validTickets) {
				delete(fieldMap[num], name)
			}
		}
	}

	// Assign the field numbers with only one field name option and remove them from the pool
	// of possible field names to assign, this should go to completion assuming every round
	// of filtering will yield some field numbers to which only one field name can be matched
	// (The count is now redundant info)
	assigned := make(map[string]int)
	for len(assigned) < len(fieldMap) {
		for num, names := range fieldMap {
			// Remove assigned names
			for n := range names {
				if _, in := assigned[n]; in {
					delete(names, n)
				}
			}
			if len(names) == 1 {
				for n := range names {
					assigned[n] = num
				}
				continue
			}
		}
	}
	depMult = 1
	for name, num := range assigned {
		if strings.HasPrefix(name, "departure") {
			depMult *= int64(myTicket[num])
		}
	}
	return
}

func main() {
	loadInput("input.txt")
	sum, validTickets := solve1()
	fmt.Println("Part one:", sum)                  // 19060
	fmt.Println("Part two:", solve2(validTickets)) // 953713095011
}
