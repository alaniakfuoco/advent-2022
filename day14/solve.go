package day14

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	// build rocks
	rocks := make(map[util.Point]bool)
	for _, line := range data {
		corners := strings.Split(line, " -> ")
		for i := range corners {
			if i < len(corners)-1 {
				p1 := parsePoint(corners[i])
				p2 := parsePoint(corners[i+1])
				path := parsePath(p1, p2)
				rocks[p1] = true
				rocks[p2] = true
				for _, p := range path {
					rocks[p] = true
				}
			}
		}
	}

	// find threshold
	threshold := 0
	for p := range rocks {
		threshold = util.Max(threshold, p.Y)
	}
	fmt.Println("threashold:", threshold)
	//fmt.Println("rocks:", rocks)

	// simulate sands
	sands := make(map[util.Point]bool)
	sandSource := 500
	die := false
	// generate sand loop
	for !die {
		nextSandX := sandSource
		nextSandY := 0
		// simulate sand loop
		for !die {
			if nextSandY > threshold {
				die = true
				break
			}
			below := util.Point{X: nextSandX, Y: nextSandY + 1}
			belowLeft := util.Point{X: nextSandX - 1, Y: nextSandY + 1}
			belowRight := util.Point{X: nextSandX + 1, Y: nextSandY + 1}
			if rocks[below] || sands[below] {
				if rocks[belowLeft] || sands[belowLeft] {
					if rocks[belowRight] || sands[belowRight] { // something everywhere
						sands[util.Point{X: nextSandX, Y: nextSandY}] = true
						break
					} else { // something below and to the left but not the right
						nextSandX++
						nextSandY++
						continue
					}
				} else { // something below but nothing below left
					nextSandX--
					nextSandY++
					continue
				}
			} else { // nothing below
				nextSandY++
				continue
			}
		}
	}
	//fmt.Println("sands", sands)
	fmt.Println("count:", len(sands))
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	// build rocks
	rocks := make(map[util.Point]bool)
	for _, line := range data {
		corners := strings.Split(line, " -> ")
		for i := range corners {
			if i < len(corners)-1 {
				p1 := parsePoint(corners[i])
				p2 := parsePoint(corners[i+1])
				path := parsePath(p1, p2)
				rocks[p1] = true
				rocks[p2] = true
				for _, p := range path {
					rocks[p] = true
				}
			}
		}
	}

	// find threshold
	threshold := 0
	maxX := 0
	for p := range rocks {
		threshold = util.Max(threshold, p.Y)
		maxX = util.Max(maxX, p.X)
	}
	threshold += 2

	// add floor
	floor := parsePath(util.Point{X: 0, Y: threshold}, util.Point{X: maxX + 1000, Y: threshold})
	for _, p := range floor {
		rocks[p] = true
	}

	//fmt.Println("threashold:", threshold)
	//fmt.Println("rocks:", rocks)

	// simulate sands
	sands := make(map[util.Point]bool)
	sandSource := 500
	die := false
	// generate sand loop
	for !die {
		nextSandX := sandSource
		nextSandY := 0
		// simulate sand loop
		for !die {
			if nextSandY > threshold {
				die = true
				fmt.Println("went past floor")
				break
			}
			below := util.Point{X: nextSandX, Y: nextSandY + 1}
			belowLeft := util.Point{X: nextSandX - 1, Y: nextSandY + 1}
			belowRight := util.Point{X: nextSandX + 1, Y: nextSandY + 1}
			if rocks[below] || sands[below] {
				if rocks[belowLeft] || sands[belowLeft] {
					if rocks[belowRight] || sands[belowRight] { // something everywhere
						sands[util.Point{X: nextSandX, Y: nextSandY}] = true
						if nextSandX == 500 && nextSandY == 0 {
							die = true
						}
						break
					} else { // something below and to the left but not the right
						nextSandX++
						nextSandY++
						continue
					}
				} else { // something below but nothing below left
					nextSandX--
					nextSandY++
					continue
				}
			} else { // nothing below
				nextSandY++
				continue
			}
		}
	}
	//fmt.Println("sands", sands)
	fmt.Println("count:", len(sands))
}

func parsePath(a, b util.Point) []util.Point {
	result := []util.Point{}
	if a.X == b.X {
		for nextY := util.Min(a.Y, b.Y) + 1; nextY < util.Max(a.Y, b.Y); nextY++ {
			result = append(result, util.Point{X: a.X, Y: nextY})
		}
	} else {
		for nextX := util.Min(a.X, b.X) + 1; nextX < util.Max(a.X, b.X); nextX++ {
			result = append(result, util.Point{X: nextX, Y: a.Y})
		}
	}
	return result
}

func parsePoint(s string) util.Point {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	return util.Point{X: x, Y: y}
}
