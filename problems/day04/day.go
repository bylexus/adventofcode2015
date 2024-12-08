package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"

	"alexi.ch/aoc/2015/lib"
)

type Day struct {
	s1  int
	s2  int
	key string
}

func New() Day {
	return Day{s1: 0, s2: 0}
}

func (d *Day) Title() string {
	return "Day 04 - The Ideal Stocking Stuffer"
}

func (d *Day) Setup() {
	// var lines = lib.ReadLines("data/04-test-data.txt")
	var lines = lib.ReadLines("data/04-data.txt")
	d.key = strings.TrimSpace(lines[0])
	// fmt.Printf("%v\n", d.numbers)
}

func (d *Day) SolveProblem1() {
	d.s1 = 0
	nr := 0
	for {
		input := fmt.Sprintf("%s%d", d.key, nr)
		h := md5.Sum([]byte(input))
		hs := hex.EncodeToString(h[:])
		if strings.HasPrefix(hs, "00000") {
			d.s1 = nr
			break
		}
		nr++
	}
}

func (d *Day) SolveProblem2() {
	d.s2 = 0
	nr := 0
	for {
		input := fmt.Sprintf("%s%d", d.key, nr)
		h := md5.Sum([]byte(input))
		hs := hex.EncodeToString(h[:])
		if strings.HasPrefix(hs, "000000") {
			d.s2 = nr
			break
		}
		nr++
	}
}

func (d *Day) Solution1() string {
	return fmt.Sprintf("%d", d.s1)
}

func (d *Day) Solution2() string {
	return fmt.Sprintf("%d", d.s2)
}
