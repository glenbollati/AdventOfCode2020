package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bags map[string]map[string]int

func (bs Bags) addBag(line string) {
	clean := strings.ReplaceAll(line, ",", "")
	s := bufio.NewScanner(strings.NewReader(clean))
	s.Split(bufio.ScanWords)

	// The first two words describe the subject
	s.Scan()
	name := s.Text()
	s.Scan()
	name += " " + s.Text()

	// Now we scan until we encounter "no" or an integer
	bs[name] = make(map[string]int)
	for s.Scan() {
		if s.Text() == "no" {
			bs[name] = nil
		} else {
			num, err := strconv.Atoi(s.Text())
			// The next two words describe the contained bag
			if err == nil {
				s.Scan()
				contained := s.Text()
				s.Scan()
				contained += " " + s.Text()
				bs[name][contained] = num
			}
		}
	}
}

func targetBagCounter(bs Bags, b map[string]int) (count int) {
	if len(b) == 0 {
		return
	}
	for c, _ := range b {
		if c == "shiny gold" {
			count++
			continue
		}
		rc := targetBagCounter(bs, bs[c])
		count += rc
	}
	return
}

func countAllShinyGoldBags(bs Bags) int {
	count := 0
	for _, b := range bs {
		if targetBagCounter(bs, b) > 0 {
			count++
		}
	}
	return count
}

func getLines(fname string) (ls []string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	return ls
}

func bagsContained(bs Bags, name string) (count int) {
	topBag := bs[name]
	//fmt.Println("Checking in bag:", name)
	if len(topBag) == 0 {
		return
	}
	for b, n := range topBag {
		count += n
		count += bagsContained(bs, b) * n
	}
	return
}

func main() {
	ls := getLines("input.txt")
	bs := make(Bags)
	for _, l := range ls {
		bs.addBag(l)
	}
	fmt.Println(countAllShinyGoldBags(bs))       // 222
	fmt.Println(bagsContained(bs, "shiny gold")) // 13264
}
