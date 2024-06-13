package day10

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	data := util.GetFileLines(path)
	cpu := util.NewCompuder()

	signal := 0
	for cpu.ProgramCounter < len(data) {
		cpu.Cycle(data)
		if cpu.CycleCount%40 == 20 {
			fmt.Println(cpu.CycleCount, cpu.RegX, cpu.CycleCount*cpu.RegX)
			signal += cpu.CycleCount * cpu.RegX
		}
	}
	fmt.Println(signal)
	//fmt.Println(cpu)
}

func Solve2(path string) {
	data := util.GetFileLines(path)
	cpu := util.NewCompuder()

	crt := make([][]bool, 6)
	for i := range crt {
		line := make([]bool, 40)
		crt[i] = line
	}

	for cpu.ProgramCounter < len(data) {
		if cpu.CycleCount < 240 {
			x := (cpu.CycleCount - 1) % 40
			y := (cpu.CycleCount - 1) / 40

			if util.Abs(cpu.RegX-x) <= 1 {
				crt[y][x] = true
			}
		}
		cpu.Cycle(data)
	}

	printCrt(crt)
}

func printCrt(crt [][]bool) {
	for _, a := range crt {
		for _, s := range a {
			if s {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
