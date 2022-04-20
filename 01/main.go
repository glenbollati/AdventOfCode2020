package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func partOne(nums []int) error {
	for i, n := range nums {
		for ii, nn := range nums {
			if ii == i {
				continue
			}
			if n+nn == 2020 {
				fmt.Println(n, "+", nn, "=", n+nn)
				fmt.Println("ANSWER:", n, "*", nn, "=", n*nn)
				return nil
			}
		}
	}
	return errors.New("Answer not found!")
}

func partTwo(nums []int) error {
	for i, n := range nums {
		for ii, nn := range nums {
			for iii, nnn := range nums {
				if i == ii || ii == iii || i == iii {
					continue
				}
				if n+nn+nnn == 2020 {
					fmt.Printf("%d + %d + %d = %d\n", n, nn, nnn, n+nn+nnn)
					fmt.Printf("ANSWER: %d * %d * %d = %d\n", n, nn, nnn, n*nn*nnn)
					return nil
				}
			}
		}
	}
	return errors.New("Answer not found!")
}

func main() {
	lines := linesFromFile("input.txt")
	var nums []int
	for i, l := range lines {
		n, err := strconv.Atoi(l)
		if err != nil {
			fmt.Println(i, err)
		}
		nums = append(nums, n)
	}
	if err := partOne(nums); err != nil {
		log.Fatal(err)
	}
	if err := partTwo(nums); err != nil {
		log.Fatal(err)
	}
}
