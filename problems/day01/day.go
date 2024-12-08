package day01

import (
	"fmt"

	"alexi.ch/aoc/2015/lib"
)

type Day struct {
	input string
	s1    int
	s2    int
}

func New() Day {
	return Day{}
}

func (d *Day) Title() string {
	return "Day 01 - Not Quite Lisp"
}

func (d *Day) Setup() {
	var lines = lib.ReadLines("data/01-data.txt")
	// var lines = lib.ReadLines("data/01-test-data.txt")
	d.input = lines[0]
	// fmt.Printf("%#v\n", d.rightNumber)
}

func (d *Day) SolveProblem1() {
	d.s1 = 0
	for _, r := range d.input {
		if r == '(' {
			d.s1++
		}
		if r == ')' {
			d.s1--
		}
	}
}

func (d *Day) SolveProblem2() {
	d.s2 = 0
	for i, r := range d.input {
		if r == '(' {
			d.s2++
		}
		if r == ')' {
			d.s2--
		}
		if d.s2 == -1 {
			d.s2 = i + 1
			break
		}
	}
}

func (d *Day) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
