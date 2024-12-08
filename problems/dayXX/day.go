package dayxx

import (
	"fmt"

	"alexi.ch/aoc/2015/lib"
)

type Day struct {
	s1 uint64
	s2 uint64
}

func New() Day {
	return Day{s1: 0, s2: 0}
}

func (d *Day) Title() string {
	return "Day XX - xxxxxx"
}

func (d *Day) Setup() {
	// var lines = lib.ReadLines("data/xx-test-data.txt")
	var lines = lib.ReadLines("data/xx-data.txt")
	for _, line := range lines {
		line = line
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day) SolveProblem1() {
	d.s1 = 0
}

func (d *Day) SolveProblem2() {
	d.s2 = 0
}

func (d *Day) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
