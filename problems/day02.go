package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 02 - I Was Told There Would Be No Math
// ----------------------------------------------------------------------------

import (
	"fmt"
	"sort"

	"alexi.ch/aoc2015/lib"
)

type Day02Dim struct {
	l int
	w int
	h int
}

type Day02 struct {
	input     []Day02Dim
	solution1 int
	solution2 int
}

func (p *Day02) GetName() string {
	return "I Was Told There Would Be No Math"
}

func (p *Day02) Init() {
	// Read input
	// input := lib.ParseGroupMatch(lib.ReadInputLines("input/day02-test.txt"), `(\d+)x(\d+)x(\d+)`)
	input := lib.ParseGroupMatch(lib.ReadInputLines("input/day02-input.txt"), `(\d+)x(\d+)x(\d+)`)
	for _, d := range input {
		p.input = append(p.input, Day02Dim{
			l: lib.ToInt(d[1]),
			w: lib.ToInt(d[2]),
			h: lib.ToInt(d[3]),
		})
	}
}

func (p *Day02) Run1() {
	tot := 0
	for _, d := range p.input {
		lw := d.l * d.w
		wh := d.w * d.h
		hl := d.h * d.l
		min := lw
		if wh < min {
			min = wh
		}
		if hl < min {
			min = hl
		}
		tot += 2*lw + 2*wh + 2*hl + min
	}
	p.solution1 = tot
}

func (p *Day02) Run2() {
	tot := 0
	for _, d := range p.input {
		dims := []int{d.h, d.l, d.w}
		sort.Ints(dims)
		tot += 2*dims[0] + 2*dims[1] + d.h*d.l*d.w
	}
	p.solution2 = tot
}

func (p *Day02) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day02) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
