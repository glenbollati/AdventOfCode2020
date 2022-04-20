package main

import (
	aoc "aoc/utils"
	"fmt"
	"sort"
	"strings"
)

var players [][]int
var totalCards int

func load(fname string) {
	sections := aoc.GetSections(fname, "")
	for i := range sections {
		players = append(players, []int{})
		for j, line := range sections[i] {
			if j == 0 || strings.Contains(line, "Player") {
				continue
			}
			players[i] = append(players[i], aoc.ToInt(line))
		}
		totalCards += len(players[i])
	}
}

type winCard struct {
	idx   int
	value int
}

func newWinCard() winCard {
	return winCard{idx: -1, value: 0}
}

func p1() {
	var winner int
	var playedCards []int
	var round int
WIN:
	for {
		//fmt.Printf("-- Round %d\n", round)
		round++
		// Check for end of game
		for i := range players {
			//fmt.Println(len(players[i]))
			if len(players[i]) == totalCards {
				winner = i
				break WIN
			}
		}
		// Run a round
		// Play the cards
		playedCards = []int{}
		for _, p := range players {
			//fmt.Printf("Player %d plays: %d\n", i, p[0])
			playedCards = append(playedCards, p[0])
		}
		//fmt.Println("Played cards: ", playedCards)
		// Find the winner
		wc := newWinCard()
		for i, c := range playedCards {
			if c > wc.value {
				wc.value = c
				wc.idx = i
			}
		}
		//fmt.Println("Winning card:", wc)
		// Redistribute the cards (winner takes all)
		sort.Sort(sort.Reverse(sort.IntSlice(playedCards)))
		for i := range players {
			//fmt.Printf("Player %d before: %d\n", i, players[i])
			if i == wc.idx { // winner
				players[i] = append(players[i][1:], playedCards...)
				//fmt.Printf("Player %d after: %d\n", i, players[i])
				continue
			}
			players[i] = players[i][1:]
			//fmt.Printf("Player %d after: %d\n", i, players[i])
		}
	}
	// Calculate final score
	acc := 0
	mult := len(players[winner])
	for i := range players[winner] {
		//fmt.Printf("%d * %d\n", players[winner][i], mult)
		acc += (players[winner][i] * mult)
		mult -= 1
	}
	fmt.Println(acc)
}

var rounds [][][]int

func main() {
	load("input.txt")
	p1() // 32272
}
