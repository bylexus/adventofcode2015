package day02

import (
	"fmt"
	"regexp"
	"slices"

	"alexi.ch/aoc/2015/lib"
)

type Dim struct {
	w int
	l int
	h int
}

func (d Dim) Area() int {
	return 2*d.l*d.w + 2*d.w*d.h + 2*d.h*d.l
}

func (d Dim) ExtraPaper() int {
	return slices.Min([]int{d.w * d.l, d.l * d.h, d.h * d.w})
}

type Day struct {
	s1   int
	s2   int
	dims []Dim
}

func New() Day {
	return Day{s1: 0, s2: 0}
}

func (d *Day) Title() string {
	return "Day 02 - I Was Told There Would Be No Math"
}

func (d *Day) Setup() {
	// var lines = lib.ReadLines("data/02-test-data.txt")
	var lines = lib.ReadLines("data/02-data.txt")
	var matcher = regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)
	for _, line := range lines {
		parts := matcher.FindStringSubmatch(line)
		if len(parts) == 4 {
			dim := Dim{l: lib.StrToInt(parts[1]), w: lib.StrToInt(parts[2]), h: lib.StrToInt(parts[3])}
			d.dims = append(d.dims, dim)
		}
	}
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day) SolveProblem1() {
	d.s1 = 0
	for _, dim := range d.dims {
		d.s1 += dim.Area() + dim.ExtraPaper()
	}
}

func (d *Day) SolveProblem2() {
	d.s2 = 0
	for _, dim := range d.dims {
		sides := []int{dim.w, dim.l, dim.h}
		slices.Sort(sides)
		d.s2 += sides[0]*2 + sides[1]*2 + dim.l*dim.h*dim.w
	}
}

func (d *Day) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
