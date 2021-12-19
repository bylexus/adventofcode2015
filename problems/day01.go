package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 01 - Not Quite Lisp
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2015/lib"
)

type Day01 struct {
	input     []rune
	solution1 int
	solution2 int
}

func (p *Day01) GetName() string {
	return "Not Quite Lisp"
}

func (p *Day01) Init() {
	// Read input
	// p.input = lib.ParseIntLines(lib.ReadInputLines("input/day01-test.txt"))
	p.input = []rune(lib.ReadInputLines("input/day01-input.txt")[0])

}

func (p *Day01) Run1() {
	counter := 0
	for _, r := range p.input {
		if r == '(' {
			counter++
		}
		if r == ')' {
			counter--
		}
	}
	p.solution1 = counter
}

func (p *Day01) Run2() {
	counter := 0
	for i, r := range p.input {
		if r == '(' {
			counter++
		}
		if r == ')' {
			counter--
		}
		if counter == -1 {
			p.solution2 = i + 1
			break
		}
	}
}

func (p *Day01) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day01) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
