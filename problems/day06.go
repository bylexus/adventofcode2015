package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 06 - Probably a Fire Hazard
// ----------------------------------------------------------------------------

import (
	"fmt"

	"alexi.ch/aoc2015/lib"
	"alexi.ch/aoc2015/lib/types"
)

type Day06Input struct {
	verb string
	from types.Point
	to   types.Point
}

type Day06 struct {
	solution1 int
	solution2 int
	grid      [][]int
	input     []Day06Input
}

func (p *Day06) GetName() string {
	return "Probably a Fire Hazard"
}

func (p *Day06) Init() {
	// Read input
	// p.input = lib.ReadInputLines("input/day05-test.txt")
	input := lib.ParseGroupMatch(lib.ReadInputLines("input/day06-input.txt"), `(turn off|turn on|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	p.grid = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		p.grid[i] = make([]int, 1000)
	}

	for _, line := range input {
		in := Day06Input{
			verb: line[1],
			from: types.Point{X: lib.ToInt(line[2]), Y: lib.ToInt(line[3])},
			to:   types.Point{X: lib.ToInt(line[4]), Y: lib.ToInt(line[5])},
		}
		p.input = append(p.input, in)
	}
}

func (p *Day06) Run1() {
	p.grid = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		p.grid[i] = make([]int, 1000)
	}
	for _, in := range p.input {
		for x := in.from.X; x <= in.to.X; x++ {
			for y := in.from.Y; y <= in.to.Y; y++ {
				switch in.verb {
				case "turn on":
					p.grid[y][x] = 1
				case "turn off":
					p.grid[y][x] = 0
				case "toggle":
					if p.grid[y][x] == 1 {
						p.grid[y][x] = 0
					} else {
						p.grid[y][x] = 1
					}
				}
			}
		}
	}
	count := 0
	for y := 0; y < len(p.grid); y++ {
		for x := 0; x < len(p.grid[y]); x++ {
			if p.grid[y][x] == 1 {
				count++
			}
		}
	}
	p.solution1 = count
}

func (p *Day06) Run2() {
	p.grid = make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		p.grid[i] = make([]int, 1000)
	}
	for _, in := range p.input {
		for x := in.from.X; x <= in.to.X; x++ {
			for y := in.from.Y; y <= in.to.Y; y++ {
				switch in.verb {
				case "turn on":
					p.grid[y][x] += 1
				case "turn off":
					p.grid[y][x] = lib.MaxInt(0, p.grid[y][x]-1)
				case "toggle":
					p.grid[y][x] += 2
				}
			}
		}
	}
	count := 0
	for y := 0; y < len(p.grid); y++ {
		for x := 0; x < len(p.grid[y]); x++ {
			count += p.grid[y][x]
		}
	}
	p.solution2 = count
}

func (p *Day06) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day06) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
