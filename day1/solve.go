package day1

import (
	"advent/util"
	"fmt"
	"strconv"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	max := 0
	sum := 0
	for _, line := range data {
		if line == "" { // empty line
			if sum > max {
				max = sum
			}
			sum = 0
		} else { // number
			val, _ := strconv.Atoi(line)
			sum += val
		}
		// fmt.Print(line + " | ")
		// fmt.Print(strconv.Itoa(sum) + " | ")
		// fmt.Println(max)
	}

	if sum > max {
		max = sum
	}

	fmt.Println(max)

}

func Solve2(path string) {
	data := util.GetFileLines(path)

	max1 := 0
	max2 := 0
	max3 := 0
	sum := 0
	for _, line := range data {
		if line == "" { // empty line
			if sum > max3 {
				max3 = sum
			}
			if max3 > max2 {
				temp := max2
				max2 = max3
				max3 = temp
			}
			if max2 > max1 {
				temp := max1
				max1 = max2
				max2 = temp
			}
			sum = 0
		} else { // number
			val, _ := strconv.Atoi(line)
			sum += val
		}
	}

	if sum > max3 {
		max3 = sum
	}
	if max3 > max2 {
		temp := max2
		max2 = max3
		max3 = temp
	}
	if max2 > max1 {
		temp := max1
		max1 = max2
		max2 = temp
	}

	fmt.Println(max1, max2, max3)
	fmt.Println(max1 + max2 + max3)

}
