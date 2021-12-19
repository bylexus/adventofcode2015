package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 03 - Perfectly Spherical Houses in a Vacuum
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2015/lib"
	"alexi.ch/aoc2015/lib/types"
)

type Day03 struct {
	input     []rune
	solution1 int
	solution2 int
	locations map[types.Point]int
}

func (p *Day03) GetName() string {
	return "Perfectly Spherical Houses in a Vacuum"
}

func (p *Day03) Init() {
	// Read input
	// input := lib.ParseGroupMatch(lib.ReadInputLines("input/day02-test.txt"), `(\d+)x(\d+)x(\d+)`)
	p.input = []rune(lib.ReadInputLines("input/day03-input.txt")[0])
}

func (p *Day03) Run1() {
	p.locations = make(map[types.Point]int)
	pos := types.Point{X: 0, Y: 0}
	posMap := map[rune]types.Point{
		'^': {X: 0, Y: -1},
		'v': {X: 0, Y: 1},
		'<': {X: -1, Y: 0},
		'>': {X: 1, Y: 0},
	}
	p.locations[pos] = 1
	count := 1
	for _, d := range p.input {
		pos = pos.Add(posMap[d])
		p.locations[pos]++
		if p.locations[pos] == 1 {
			count++
		}
	}
	p.solution1 = count
}

func (p *Day03) Run2() {
	p.locations = make(map[types.Point]int)
	pos := []types.Point{
		{X: 0, Y: 0},
		{X: 0, Y: 0},
	}
	posMap := map[rune]types.Point{
		'^': {X: 0, Y: -1},
		'v': {X: 0, Y: 1},
		'<': {X: -1, Y: 0},
		'>': {X: 1, Y: 0},
	}
	p.locations[pos[0]] = 1
	count := 1
	pIndex := 0
	for _, d := range p.input {
		pos[pIndex] = pos[pIndex].Add(posMap[d])
		p.locations[pos[pIndex]]++
		if p.locations[pos[pIndex]] == 1 {
			count++
		}
		pIndex = (pIndex + 1) % len(pos)
	}
	p.solution2 = count
}

func (p *Day03) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day03) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
