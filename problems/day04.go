package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 04 - The Ideal Stocking Stuffer
// ----------------------------------------------------------------------------

import (
	"crypto/md5"
	"fmt"
)

type Day04 struct {
	input     string
	solution1 uint64
	solution2 uint64
}

func (p *Day04) GetName() string {
	return "The Ideal Stocking Stuffer"
}

func (p *Day04) Init() {
	// Read input
	// input := lib.ParseGroupMatch(lib.ReadInputLines("input/day04-test.txt"), `(\d+)x(\d+)x(\d+)`)
	// p.input = "abcdef"
	p.input = "yzbqklnj"

}

func (p *Day04) findNr(nrOfZeroes int) uint64 {
	var nr uint64 = 0
	var zeroes string = ""
	for i := 0; i < nrOfZeroes; i++ {
		zeroes += "0"
	}
	for {
		data := []byte(p.input + fmt.Sprint(nr))
		hash := fmt.Sprintf("%x", md5.Sum(data))
		if hash[:nrOfZeroes] == zeroes {
			break
		}
		nr++
	}
	return nr
}

func (p *Day04) Run1() {
	p.solution1 = p.findNr(5)
}

func (p *Day04) Run2() {
	p.solution2 = p.findNr(6)
}

func (p *Day04) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day04) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
