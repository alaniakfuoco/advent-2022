package util

import (
	"strconv"
	"strings"
)

type Compuder struct {
	CycleCount     int
	ProgramCounter int
	RegX           int
	NextInstruc    instruction
	Executing      int
}

type instruction struct {
	code  string
	input int
}

func NewCompuder() *Compuder {
	return &Compuder{
		1,
		0,
		1,
		instruction{},
		0,
	}
}

func (c *Compuder) Execute(program []string) {
	for c.ProgramCounter < len(program) {
		c.Cycle(program)
	}
}

func (c *Compuder) Cycle(program []string) {
	if c.Executing > 0 { // execute current command
		c.Executing--
		if c.Executing == 0 {
			c.RunNextInstruction()
		}
	} else { // read next command
		ins := program[c.ProgramCounter]
		c.ProgramCounter++
		parts := strings.Split(ins, " ")
		switch parts[0] {
		case "noop":
			c.Executing = 0
		case "addx":
			c.Executing = 1
			val, _ := strconv.Atoi(parts[1])
			c.NextInstruc = instruction{
				code:  parts[0],
				input: val,
			}
		}
	}
	c.CycleCount++
}

func (c *Compuder) RunNextInstruction() {
	switch c.NextInstruc.code {
	case "addx":
		//fmt.Println(c.NextInstruc)
		c.RegX += c.NextInstruc.input
	}
}
