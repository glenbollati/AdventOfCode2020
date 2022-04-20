package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetLines(fname string) (lines []string) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		panic(err)
	}
	return
}

func GetSections(fname, sep string) (lines [][]string) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	i := 0
	lines = [][]string{{}}
	for s.Scan() {
		if s.Text() == sep {
			i++
			lines = append(lines, []string{})
			//lines[i] = []string{}
			continue
		}
		lines[i] = append(lines[i], s.Text())
	}
	if s.Err() != nil {
		panic(err)
	}
	return
}

func GetDigitGrid(fname string) (grid [][]int) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanBytes)
	row := []int{}
	for s.Scan() {
		if s.Text() == "\n" {
			grid = append(grid, row)
			row = []int{}
			continue
		}
		row = append(row, ToInt(s.Text()))
	}
	if s.Err() != nil {
		panic(err)
	}
	return
}

func SplitFile(fname, sep string) []string {
	tmp, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	if tmp[len(tmp)-1] == '\n' {
		tmp = tmp[:len(tmp)-1]
	}
	return strings.Split(string(tmp), sep)
}

func SIContains(s []int, e int) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}

func SSContains(s []string, e string) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}

func SSAppendUnique(ss []string, s string) []string {
	if SSContains(ss, s) {
		return ss
	}
	return append(ss, s)
}

func SIAppendUnique(si []int, i int) []int {
	if SIContains(si, i) {
		return si
	}
	return append(si, i)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func SStoSI(ss []string) (si []int) {
	for _, s := range ss {
		si = append(si, ToInt(s))
	}
	return
}

func BinStrToI64(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ReverseStr(s string) (sn string) {
	for i := len(s) - 1; i >= 0; i-- {
		sn += string(s[i])
	}
	return
}

func SSRemove(ss []string, s string) (so []string) {
	for _, e := range ss {
		if e != s {
			so = append(so, e)
		}
	}
	return
}

func SSPop(ss []string, idx int) (so []string) {
	if idx >= len(ss) {
		log.Panicf("Cannot pop element %d from slice of length %d\n", idx, len(ss))
	}
	if idx == len(ss)-1 {
		return ss[:idx]
	}
	return append(ss[:idx], ss[idx+1])
}

func SIReverse(si []int) (so []int) {
	for i := len(si) - 1; i >= 0; i-- {
		so = append(so, si[i])
	}
	return
}

func SIFastRemove(si []int, i int) []int {
	si[i] = si[len(si)-1]
	si[len(si)-1] = 0
	si = si[:len(si)-1]
	return si
}

func SIRemove(si []int, i int) []int {
	if i == len(si)-1 {
		return si[:len(si)-1]
	}
	copy(si[i:], si[i+1:])
	si[len(si)-1] = 0
	si = si[:len(si)-1]
	return si
}

// The standard library has these for floats
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(si ...int) (min int) {
	for i, v := range si {
		if i == 0 || v < min {
			min = v
		}
	}
	return
}

func Max(si ...int) (max int) {
	for i, v := range si {
		if i == 0 || v > max {
			max = v
		}
	}
	return
}

func MinMax(si ...int) (min, max int) {
	for i, v := range si {
		if i == 0 {
			min, max = v, v
			continue
		}
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func SIUnique(si []int) (so []int) {
	m := make(map[int]struct{})
	for _, x := range si {
		if _, in := m[x]; in {
			continue
		}
		m[x] = struct{}{}
		so = append(so, x)
	}
	return
}

func TimeFunc(fn func(), name string) {
	start := time.Now()
	fn()
	fmt.Println(name, "took", time.Since(start))
}

// TODO: find a better way, this feels expensive
func SortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}
