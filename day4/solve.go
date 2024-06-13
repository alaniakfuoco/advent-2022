package day4

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

type area struct {
	min int
	max int
}

func Solve(path string) {
	data := util.GetFileLines(path)

	count := 0
	for _, line := range data {
		a1, a2 := parseAreas(line)
		// fmt.Print(a1, a2, " : ")
		if a1.min == a2.min && a1.max == a2.max { // equal
			//fmt.Println("equal")
			count++
		} else if a1.min == a2.min && a1.max > a2.max { // left
			// fmt.Println("left")
			count++
		} else if a1.min < a2.min && a1.max == a2.max { // right
			// fmt.Println("right")
			count++
		} else if a1.min < a2.min && a1.max > a2.max { // middle
			// fmt.Println("middle")
			count++
		}
	}
	fmt.Println(count)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	count := 0
	for _, line := range data {
		a1, a2 := parseAreas(line)
		// fmt.Print(a1, a2, " : ")
		if a2.min >= a1.min && a2.min <= a1.max {
			count++
		} else if a2.max >= a1.min && a2.max <= a1.max {
			count++
		}
	}
	fmt.Println(count)
}

func parseAreas(s string) (area, area) {
	pair := strings.Split(s, ",")
	a1 := parseArea(pair[0])
	a2 := parseArea(pair[1])

	if a1.size() >= a2.size() {
		return a1, a2
	} else {
		return a2, a1
	}
}

func parseArea(s string) area {
	pair := strings.Split(s, "-")
	mn, _ := strconv.Atoi(pair[0])
	mx, _ := strconv.Atoi(pair[1])
	return area{
		min: mn,
		max: mx,
	}
}

func (a area) size() int {
	return a.max - a.min + 1
}
