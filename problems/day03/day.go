package day03

import (
	"fmt"

	"alexi.ch/aoc/2015/lib"
)

type Day struct {
	s1    int
	s2    int
	input string
}

func New() Day {
	return Day{s1: 0, s2: 0}
}

func (d *Day) Title() string {
	return "Day 03 - Perfectly Spherical Houses in a Vacuum"
}

func (d *Day) Setup() {
	// var lines = lib.ReadLines("data/03-test-data.txt")
	var lines = lib.ReadLines("data/03-data.txt")
	d.input = lines[0]
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day) SolveProblem1() {
	d.s1 = 0
	visitMap := make(map[lib.Coord]int)
	actCoord := lib.NewCoord0()
	visitMap[actCoord] = 1

	for _, r := range d.input {
		switch r {
		case '^':
			actCoord.Y--
		case '>':
			actCoord.X++
		case 'v':
			actCoord.Y++
		case '<':
			actCoord.X--
		}
		visitMap[actCoord]++
	}

	d.s1 = len(visitMap)
}

func (d *Day) SolveProblem2() {
	d.s2 = 0
	visitMap := make(map[lib.Coord]int)
	santaCoord := lib.NewCoord0()
	visitMap[santaCoord] = 1
	roboCoord := lib.NewCoord0()

	for i := 0; i < len(d.input); i++ {
		var actCoord *lib.Coord
		if i%2 == 0 {
			actCoord = &santaCoord
		} else {
			actCoord = &roboCoord
		}

		switch d.input[i] {
		case '^':
			actCoord.Y--
		case '>':
			actCoord.X++
		case 'v':
			actCoord.Y++
		case '<':
			actCoord.X--
		}
		visitMap[*actCoord]++
	}

	d.s2 = len(visitMap)
}

func (d *Day) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
