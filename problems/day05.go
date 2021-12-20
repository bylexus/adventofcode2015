package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 05 - Doesn't He Have Intern-Elves For This?
// ----------------------------------------------------------------------------

import (
	"fmt"
	"regexp"
	"strings"

	"alexi.ch/aoc2015/lib"
)

type Day05 struct {
	input     []string
	solution1 uint64
	solution2 uint64
}

func (p *Day05) GetName() string {
	return "Doesn't He Have Intern-Elves For This?"
}

func (p *Day05) Init() {
	// Read input
	// p.input = lib.ReadInputLines("input/day05-test.txt")
	p.input = lib.ReadInputLines("input/day05-input.txt")
}

func (p *Day05) checkDouble(in string) bool {
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			return true
		}
	}
	return false
}

func (p *Day05) checkSameWithGap(in string) bool {
	for i := 0; i < len(in)-2; i++ {
		if in[i] == in[i+2] {
			return true
		}
	}
	return false
}

func (p *Day05) checkPair(in string) bool {
	for i := 0; i < len(in)-1; i++ {
		if strings.Index(in[i+2:], in[i:i+2]) > -1 {
			return true
		}
	}
	return false
}

func (p *Day05) Run1() {
	vovelRe, err := regexp.Compile(`[aeiou].*[aeiou].*[aeiou]`)
	if err != nil {
		panic(err)
	}
	forbiddenRe, err := regexp.Compile(`ab|cd|pq|xy`)
	if err != nil {
		panic(err)
	}

	countNice := 0
	for _, input := range p.input {
		i := []byte(input)
		if vovelRe.Match(i) == true && p.checkDouble(input) == true && forbiddenRe.Match(i) != true {
			countNice++
		}
	}

	p.solution1 = uint64(countNice)
}

func (p *Day05) Run2() {
	countNice := 0
	for _, input := range p.input {
		if p.checkSameWithGap(input) && p.checkPair(input) {
			countNice++
		}
	}
	p.solution2 = uint64(countNice)
}

func (p *Day05) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day05) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
