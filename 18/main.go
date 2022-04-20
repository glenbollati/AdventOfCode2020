package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	NUMBER tokenType = iota
	ADD
	MULT
	OPEN
	CLOSE
)

type tokenType int

type Token struct {
	value string
	typ   tokenType
}

func TokenizeLine(line string) (tokSlice []Token) {
	numStart := -1
	for pos, val := range line {
		if unicode.IsDigit(val) {
			if numStart == -1 {
				numStart = pos
			}
			if pos == len(line)-1 {
				tokSlice = append(tokSlice, Token{line[numStart : pos+1], NUMBER})
			}
			continue
		}
		if numStart > -1 {
			tokSlice = append(tokSlice, Token{line[numStart:pos], NUMBER})
			numStart = -1
		}
		switch val {
		case '+':
			tokSlice = append(tokSlice, Token{"+", ADD})
		case '*':
			tokSlice = append(tokSlice, Token{"*", MULT})
		case '(':
			tokSlice = append(tokSlice, Token{"(", OPEN})
		case ')':
			tokSlice = append(tokSlice, Token{")", CLOSE})
		}
	}
	return
}

func input(fname string) (si []string) {
	f, _ := os.Open(fname)
	for s := bufio.NewScanner(f); s.Scan(); {
		si = append(si, s.Text())
	}
	return
}

func isCompound(expr []Token) bool {
	for _, t := range expr {
		if t.typ == OPEN || t.typ == CLOSE {
			return true
		}
	}
	return false
}

// TODO: return error if expr is not reduceable, instead
// or aswell as returning original expression
func Reduce(expr []Token) []Token {
	start := 0
	for pos, tok := range expr {
		if tok.typ == OPEN {
			start = pos
			continue
		}
		if tok.typ == CLOSE {
			res := Parse(expr[start+1 : pos])
			return append(append(expr[:start], res), expr[pos+1:]...)
		}
	}
	return expr
}

func printExpr(expr []Token) {
	for _, tok := range expr {
		fmt.Print(tok.value + " ")
	}
	fmt.Println()
}

func Parse(expr []Token) Token {
	for isCompound(expr) {
		expr = Reduce(expr)
	}
	for len(expr) > 1 {
		ops, nums := []Token{}, []Token{}
		for _, tok := range expr {
			switch tok.typ {
			case NUMBER:
				nums = append(nums, tok)
			case ADD, MULT:
				ops = append(ops, tok)
			}
		}
		if len(ops) == 0 && len(nums) > 0 {
			panic("Ops at 0 but have multiple nums")
		}
		currPrec := maxPrec(ops)
		for pos, op := range ops {
			if opPrec[op.typ] == currPrec {
				result := ApplyOp(op, nums[pos], nums[pos+1])
				expr = ReSlice(expr, result, pos)
				break
			}
		}
	}
	return expr[0]
}

func ReSlice(expr []Token, res Token, pos int) []Token {
	return append(append(expr[:pos*2], res), expr[pos*2+3:]...)
}

func ApplyOp(op Token, a, b Token) (ret Token) {
	ai, err := strconv.Atoi(a.value)
	if err != nil {
		panic(err)
	}
	bi, err := strconv.Atoi(b.value)
	if err != nil {
		panic(err)
	}
	ret.typ = NUMBER
	switch op.typ {
	case ADD:
		ret.value = strconv.Itoa(ai + bi)
	case MULT:
		ret.value = strconv.Itoa(ai * bi)
	}
	return
}

var opPrec map[tokenType]int

func maxPrec(ops []Token) (x int) {
	for _, o := range ops {
		if opPrec[o.typ] > x {
			x = opPrec[o.typ]
		}
	}
	return
}

func solve(part2 bool) (count int) {
	if part2 {
		opPrec = map[tokenType]int{MULT: 0, ADD: 1}
	} else {
		opPrec = map[tokenType]int{MULT: 0, ADD: 0}
	}
	for _, l := range input("input.txt") {
		tok := Parse(TokenizeLine(l))
		val, err := strconv.Atoi(tok.value)
		if err != nil {
			panic(err)
		}
		count += val
	}
	return
}

func main() {
	fmt.Println("Part one:", solve(false))
	fmt.Println("Part two:", solve(true))
}
