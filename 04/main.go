package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func parseLine(l string, p *Passport) {
	kvs := strings.Split(l, " ")
	for _, kv := range kvs {
		k := strings.Split(kv, ":")[0]
		v := strings.Split(kv, ":")[1]
		switch k {
		case "byr":
			p.byr = v
		case "iyr":
			p.iyr = v
		case "eyr":
			p.eyr = v
		case "hgt":
			p.hgt = v
		case "hcl":
			p.hcl = v
		case "ecl":
			p.ecl = v
		case "pid":
			p.pid = v
		}
	}
}

func passports(s *bufio.Scanner) (ps []*Passport) {
	lastPass := new(Passport)
	for s.Scan() {
		if s.Text() == "" {
			ps = append(ps, lastPass)
			lastPass = new(Passport)
			continue
		}
		parseLine(s.Text(), lastPass)
	}
	ps = append(ps, lastPass)
	return
}

func isValid(p *Passport) (bool, error) {
	if p.byr == "" || p.iyr == "" || p.eyr == "" ||
		p.hgt == "" || p.hcl == "" || p.ecl == "" || p.pid == "" {
		return false, nil
	}

	byr, err := strconv.Atoi(p.byr)
	if err != nil {
		return false, err
	}
	if byr < 1920 || byr > 2002 {
		return false, nil
	}
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false, err
	}
	if iyr < 2010 || iyr > 2020 {
		return false, nil
	}
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false, err
	}
	if eyr < 2020 || eyr > 2030 {
		return false, nil
	}
	if !strings.HasSuffix(p.hgt, "in") && !strings.HasSuffix(p.hgt, "cm") {
		return false, nil
	}
	if strings.HasSuffix(p.hgt, "in") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "in"))
		if err != nil {
			return false, err
		}
		if hgt < 59 || hgt > 75 {
			return false, nil
		}
	}
	if strings.HasSuffix(p.hgt, "cm") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(p.hgt, "cm"))
		if err != nil {
			return false, err
		}
		if hgt < 150 || hgt > 193 {
			return false, nil
		}
	}
	if !strings.HasPrefix(p.hcl, "#") {
		return false, nil
	}
	if len(p.hcl) != 7 {
		return false, nil
	}
	acceptedHcls := "0123456789abcdef"
	for _, l := range strings.TrimPrefix(p.hcl, "#") {
		if !strings.Contains(acceptedHcls, string(l)) {
			return false, nil
		}
	}
	acceptedEcls := "amb blu brn gry grn hzl oth"
	if !strings.Contains(acceptedEcls, string(p.ecl)) {
		return false, nil
	}
	if len(p.pid) != 9 {
		return false, nil
	}
	for _, l := range p.pid {
		if _, err := strconv.Atoi(string(l)); err != nil {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	ps := passports(s)
	fmt.Println("Count (all):", len(ps))
	valids := 0
	for _, p := range ps {
		v, err := isValid(p)
		if err != nil {
			log.Fatal(err)
		}
		if v {
			valids++
		}
	}
	fmt.Println("Count (valid):", valids)
}
