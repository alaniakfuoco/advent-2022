package day6

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	data := string(util.ReadFile(path))

	for i := 3; i < len(data); i++ {
		seen := make([]bool, 26)
		seen[calcIndex(data[i-3])] = true

		if seen[calcIndex(data[i-2])] {
			continue
		} else {
			seen[calcIndex(data[i-2])] = true
		}

		if seen[calcIndex(data[i-1])] {
			continue
		} else {
			seen[calcIndex(data[i-1])] = true
		}

		if seen[calcIndex(data[i])] {
			continue
		} else {
			fmt.Println(i + 1)
			break
		}
	}
}

func Solve2(path string) {
	data := string(util.ReadFile(path))

	for i := 13; i < len(data); i++ {
		seen := make([]bool, 26)
		seen[calcIndex(data[i-13])] = true

		die := false
		for j := 12; j > 0; j-- {
			if seen[calcIndex(data[i-j])] {
				die = true
				break
			} else {
				seen[calcIndex(data[i-j])] = true
			}
		}
		if die {
			continue
		}

		if seen[calcIndex(data[i])] {
			continue
		} else {
			fmt.Println(i + 1)
			break
		}
	}
}

func calcIndex(b byte) int {
	return int(b) - 97
}
