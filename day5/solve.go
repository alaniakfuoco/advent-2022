package day5

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

type instruc struct {
	target     int
	source     int
	executions int
}

type crateStack struct {
	stacks     []*util.Stack[string]
	stackCount int
}

func Solve(path string) {
	data := util.GetFileLines(path)

	// find the index of the input separation
	sepIndex := 0
	for i, line := range data {
		if line == "" {
			sepIndex = i
			break
		}
	}

	cs := buildCrateStack(data[sepIndex-1], sepIndex-1)

	// parse the stacks first
	// start from the bottom and work up, skipping the numbers
	for i := sepIndex - 2; i >= 0; i-- {
		stack := -1
		count := 0
		for _, v := range data[i] {
			count++
			if (count+2)%4 == 0 {
				stack++
			}
			if v != '[' && v != ']' && v != ' ' { // if the rune is a letter
				cs.stacks[stack].Push(string(v))
			}
		}
	}

	fmt.Println("===== Starting State =====")
	for _, q := range cs.stacks {
		q.Print()
	}

	for i := sepIndex + 1; i < len(data); i++ { // run instructions
		ins := parseIntstruc(data[i])
		for j := 0; j < ins.executions; j++ {
			ele := cs.stacks[ins.source-1].Pop()
			cs.stacks[ins.target-1].Push(ele)
		}
	}

	fmt.Println()
	fmt.Println("===== Finishing State =====")
	for _, q := range cs.stacks {
		q.Print()
	}
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	// find the index of the input separation
	sepIndex := 0
	for i, line := range data {
		if line == "" {
			sepIndex = i
			break
		}
	}

	cs := buildCrateStack(data[sepIndex-1], sepIndex-1)

	// parse the stacks first
	// start from the bottom and work up, skipping the numbers
	for i := sepIndex - 2; i >= 0; i-- {
		stack := -1
		count := 0
		for _, v := range data[i] {
			count++
			if (count+2)%4 == 0 {
				stack++
			}
			if v != '[' && v != ']' && v != ' ' { // if the rune is a letter
				cs.stacks[stack].Push(string(v))
			}
		}
	}

	fmt.Println("===== Starting State =====")
	for _, q := range cs.stacks {
		q.Print()
	}

	for i := sepIndex + 1; i < len(data); i++ { // run instructions
		ins := parseIntstruc(data[i])
		eles := util.NewStack[string](ins.executions)
		for i := 0; i < ins.executions; i++ {
			eles.Push(cs.stacks[ins.source-1].Pop())
		}
		for i := 0; i < ins.executions; i++ {
			cs.stacks[ins.target-1].Push(eles.Pop())
		}
	}

	fmt.Println()
	fmt.Println("===== Finishing State =====")
	for _, q := range cs.stacks {
		q.Print()
	}
}

// example: move 1 from 2 to 1
func parseIntstruc(in string) instruc {
	parts := strings.Split(in, " ")
	t, _ := strconv.Atoi(parts[5])
	s, _ := strconv.Atoi(parts[3])
	e, _ := strconv.Atoi(parts[1])
	result := instruc{
		target:     t,
		source:     s,
		executions: e,
	}
	return result
}

func buildCrateStack(s string, height int) crateStack {
	length := len(s) + 1
	size := length / 4
	cs := crateStack{
		stacks:     make([]*util.Stack[string], size),
		stackCount: size,
	}

	for i := 0; i < cs.stackCount; i++ {
		cs.stacks[i] = util.NewStack[string](height)
	}

	return cs
}
