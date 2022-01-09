package problems

// ----------------------------------------------------------------------------
// AOC 2015
// ----------
//
// Day 07 - Some Assembly Required
// ----------------------------------------------------------------------------

import (
	"fmt"
	"regexp"
	"strconv"

	"alexi.ch/aoc2015/lib"
)

type Day07Gate struct {
	fn     string
	v1     uint16
	v2     uint16
	w1     string
	w2     string
	outVal uint16
	hasOut bool
}

func (g *Day07Gate) String() string {
	return fmt.Sprintf("fn: %v, v1: %v, v2: %v, w1: %v, w2: %v, outVal: %v, hasOut: %v\n",
		g.fn,
		g.v1,
		g.v2,
		g.w1,
		g.w2,
		g.outVal,
		g.hasOut,
	)
}

type Day07 struct {
	solution1 uint16
	solution2 int
	gates     map[string]*Day07Gate
}

func (p *Day07) GetName() string {
	return "Some Assembly Required"
}

func (p *Day07) Init() {
	// Read input
	// input := lib.ReadInputLines("input/day07-test.txt")
	input := lib.ReadInputLines("input/day07-input.txt")
	p.gates = make(map[string]*Day07Gate)

	for _, i := range input {

		bingateR := regexp.MustCompile("(.*) (AND|OR|LSHIFT|RSHIFT) (.*) -> (.*)")
		notGateR := regexp.MustCompile("NOT (.*) -> (.*)")
		inOutR := regexp.MustCompile("(.*) -> (.*)")
		binGate := bingateR.FindStringSubmatch(i)
		notGate := notGateR.FindStringSubmatch(i)
		inOut := inOutR.FindStringSubmatch(i)

		// binary gate:
		if len(binGate) > 0 {
			gate := Day07Gate{}
			gate.fn = binGate[2]
			val, err := strconv.Atoi(binGate[1])
			if err == nil {
				gate.v1 = uint16(val)
				gate.w1 = ""
			} else {
				gate.w1 = binGate[1]
				gate.v1 = 0
			}
			val, err = strconv.Atoi(binGate[3])
			if err == nil {
				gate.v2 = uint16(val)
				gate.w2 = ""
			} else {
				gate.w2 = binGate[3]
				gate.v2 = 0
			}
			p.gates[binGate[4]] = &gate
		} else if len(notGate) > 0 {
			// not gate:
			gate := Day07Gate{}
			gate.fn = "NOT"
			val, err := strconv.Atoi(notGate[1])
			if err == nil {
				gate.v1 = uint16(val)
				gate.w1 = ""
			} else {
				gate.w1 = notGate[1]
				gate.v1 = 0
			}
			p.gates[notGate[2]] = &gate
		} else if len(inOut) > 0 {
			// "wire" gate connecting a value/wire to another wire
			gate := Day07Gate{}
			gate.fn = "WIRE"
			val, err := strconv.Atoi(inOut[1])
			if err == nil {
				gate.v1 = uint16(val)
				gate.w1 = ""
			} else {
				gate.w1 = inOut[1]
				gate.v1 = 0
			}
			p.gates[inOut[2]] = &gate
		} else {
			panic("Unknown input")
		}
	}

	fmt.Printf("Gates:\n%v\n\n", p.gates)
}

func (p *Day07) processWireGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 == "" {
			gate.outVal = gate.v1
			gate.hasOut = true
		} else {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.outVal = otherGate.outVal
				gate.hasOut = true
			}
		}
	}
}

func (p *Day07) processNotGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 == "" {
			gate.outVal = ^gate.v1
			gate.hasOut = true
		} else {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.outVal = ^otherGate.outVal
				gate.hasOut = true
			}
		}
	}
}

func (p *Day07) processAndGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 != "" {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v1 = otherGate.outVal
				gate.w1 = ""
			}
		}
		if gate.w2 != "" {
			otherGate := p.gates[gate.w2]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v2 = otherGate.outVal
				gate.w2 = ""
			}
		}
		if gate.w1 == "" && gate.w2 == "" {
			gate.outVal = gate.v1 & gate.v2
			gate.hasOut = true
		}
	}
}

func (p *Day07) processOrGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 != "" {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v1 = otherGate.outVal
				gate.w1 = ""
			}
		}
		if gate.w2 != "" {
			otherGate := p.gates[gate.w2]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v2 = otherGate.outVal
				gate.w2 = ""
			}
		}
		if gate.w1 == "" && gate.w2 == "" {
			gate.outVal = gate.v1 | gate.v2
			gate.hasOut = true
		}
	}
}

func (p *Day07) processLshiftGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 != "" {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v1 = otherGate.outVal
				gate.w1 = ""
			}
		}
		if gate.w2 != "" {
			otherGate := p.gates[gate.w2]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v2 = otherGate.outVal
				gate.w2 = ""
			}
		}
		if gate.w1 == "" && gate.w2 == "" {
			gate.outVal = gate.v1 << gate.v2
			gate.hasOut = true
		}
	}
}

func (p *Day07) processRshiftGate(gate *Day07Gate) {
	if gate.hasOut != true {
		if gate.w1 != "" {
			otherGate := p.gates[gate.w1]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v1 = otherGate.outVal
				gate.w1 = ""
			}
		}
		if gate.w2 != "" {
			otherGate := p.gates[gate.w2]
			p.processGate(otherGate)
			if otherGate.hasOut == true {
				gate.v2 = otherGate.outVal
				gate.w2 = ""
			}
		}
		if gate.w1 == "" && gate.w2 == "" {
			gate.outVal = gate.v1 >> gate.v2
			gate.hasOut = true
		}
	}
}

func (p *Day07) processGate(gate *Day07Gate) {
	switch gate.fn {
	case "WIRE":
		p.processWireGate(gate)
	case "NOT":
		p.processNotGate(gate)
	case "AND":
		p.processAndGate(gate)
	case "OR":
		p.processOrGate(gate)
	case "LSHIFT":
		p.processLshiftGate(gate)
	case "RSHIFT":
		p.processRshiftGate(gate)
	}
}

func (p *Day07) processGates() bool {
	repeat := false
	for _, gate := range p.gates {
		if gate.hasOut != true {
			p.processGate(gate)
			repeat = true
		}
	}
	return repeat
}

func (p *Day07) Run1() {
	p.processGates()
	fmt.Printf("Gates: %v\n", p.gates)
	gateA, defined := p.gates["a"]
	if defined != true {
		panic("Oops! Gate a not found!")
	}
	p.solution1 = gateA.outVal
}

func (p *Day07) Run2() {
	p.solution2 = 0
}

func (p *Day07) GetSolution1() string {
	return fmt.Sprintf("%v\n", p.solution1)
}

func (p *Day07) GetSolution2() string {
	return fmt.Sprintf("%v\n", p.solution2)
}
