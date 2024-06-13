package day9

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	y int
	x int
}

func (p point) calcFloat() float32 {
	result := float32(p.x)
	radix := float32(p.y)

	for radix > 1 {
		radix = radix / 10
	}

	result += radix
	return result
}

func (p point) toString() string {
	x := strconv.Itoa(p.x)
	y := strconv.Itoa(p.y)
	return "x:" + x + " y:" + y
}

func Solve(path string) {
	data := util.GetFileLines(path)

	head := point{0, 0}
	tail := point{0, 0}
	positions := make(map[string]bool)
	positions[tail.toString()] = true

	for _, line := range data {
		//fmt.Println(line)
		parts := strings.Split(line, " ")
		exe, _ := strconv.Atoi(parts[1])
		for ; exe > 0; exe-- {
			switch parts[0] {
			case "R":
				head.x++
			case "U":
				head.y++
			case "L":
				head.x--
			case "D":
				head.y--
			}
			t := isTouching(head, tail)
			if !t {
				follow(&head, &tail, parts[0])
				positions[tail.toString()] = true
			}
			// fmt.Println("H:", head)
			// fmt.Println("T:", tail)
			//printTestState(head, tail)
		}
	}

	// fmt.Println(positions)
	fmt.Println(len(positions))
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	head := point{0, 0}
	tails := [9]point{}
	positions := make(map[string]bool)
	positions[tails[8].toString()] = true

	for _, line := range data {
		// fmt.Println(line)
		parts := strings.Split(line, " ")
		exe, _ := strconv.Atoi(parts[1])
		for ; exe > 0; exe-- {
			switch parts[0] {
			case "R":
				head.x++
			case "U":
				head.y++
			case "L":
				head.x--
			case "D":
				head.y--
			}

			if !isTouching(head, tails[0]) {
				follow(&head, &tails[0], parts[0])
				for i := 1; i < len(tails); i++ {
					if !isTouching(tails[i-1], tails[i]) {
						follow2(&tails[i-1], &tails[i])
					} else {
						break
					}
				}
				positions[tails[8].toString()] = true
			}
		}
		// fmt.Println()
		// printBigTestState(head, tails[:])
	}

	fmt.Println(len(positions))
}

func follow(a, b *point, dir string) {
	switch dir {
	case "R":
		b.x++
		if a.y > b.y {
			b.y++
		} else if a.y < b.y {
			b.y--
		}
	case "U":
		b.y++
		if a.x > b.x {
			b.x++
		} else if a.x < b.x {
			b.x--
		}
	case "L":
		b.x--
		if a.y > b.y {
			b.y++
		} else if a.y < b.y {
			b.y--
		}
	case "D":
		b.y--
		if a.x > b.x {
			b.x++
		} else if a.x < b.x {
			b.x--
		}
	}
}

func follow2(a, b *point) {
	// move left or right
	if a.x == b.x {
		if a.y > b.y {
			b.y++
		} else {
			b.y--
		}
		return
	}
	// move up or down
	if a.y == b.y {
		if a.x > b.x {
			b.x++
		} else {
			b.x--
		}
		return
	}
	// diagonals
	if a.x > b.x {
		if a.y > b.y {
			b.x++
			b.y++
		} else {
			b.x++
			b.y--
		}
	} else {
		if a.y > b.y {
			b.x--
			b.y++
		} else {
			b.x--
			b.y--
		}
	}
}

func isTouching(a, b point) bool {
	// fmt.Println(util.Abs(a.x - b.x))
	// fmt.Println(util.Abs(a.y - b.y))
	if util.Abs(a.x-b.x) <= 1 && util.Abs(a.y-b.y) <= 1 {
		return true
	} else {
		return false
	}
}

func printTestState(a, b point) {
	field := make([][]string, 5)
	for i := range field {
		line := make([]string, 6)
		for j := range line {
			line[j] = "."
		}
		field[i] = line
	}
	field[0][0] = "s"
	field[b.y][b.x] = "T"
	field[a.y][a.x] = "H"

	fmt.Println("======")
	for i := len(field) - 1; i >= 0; i-- {
		for _, y := range field[i] {
			fmt.Print(y)
		}
		fmt.Println()
	}
	fmt.Println("======")
}

func printBigTestState(head point, tails []point) {
	field := make([][]string, 21)
	for i := range field {
		line := make([]string, 26)
		for j := range line {
			line[j] = "."
		}
		field[i] = line
	}
	field[0+5][0+11] = "s"
	for i := len(tails) - 1; i >= 0; i-- {
		field[tails[i].y+5][tails[i].x+11] = strconv.Itoa(i + 1)
	}
	field[head.y+5][head.x+11] = "H"

	fmt.Println("======")
	for i := len(field) - 1; i >= 0; i-- {
		for _, y := range field[i] {
			fmt.Print(y)
		}
		fmt.Println()
	}
	fmt.Println("======")
}
