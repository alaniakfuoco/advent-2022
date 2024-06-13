package day23

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	data := util.GetFileLines(path)
	elves := make(map[util.Point]bool)
	for i, line := range data {
		for j, char := range line {
			if char == '#' {
				elves[util.Point{X: j, Y: i}] = true
			}
		}
	}
	elfCount := len(elves)

	rounds := 10
	considerOrder := "NSWE"
	// fmt.Println(elves)
	for i := 0; i < rounds; i++ {
		// make proposals
		elfMovement := make(map[util.Point]util.Point)
		proposals := make(map[util.Point]int)
		for elf, exists := range elves {
			if exists && hasAdjectent(elf, elves) {
				n := util.Point{X: elf.X, Y: elf.Y - 1}
				s := util.Point{X: elf.X, Y: elf.Y + 1}
				w := util.Point{X: elf.X - 1, Y: elf.Y}
				e := util.Point{X: elf.X + 1, Y: elf.Y}
				nw := util.Point{X: elf.X - 1, Y: elf.Y - 1}
				ne := util.Point{X: elf.X + 1, Y: elf.Y - 1}
				sw := util.Point{X: elf.X - 1, Y: elf.Y + 1}
				se := util.Point{X: elf.X + 1, Y: elf.Y + 1}

				for _, char := range considerOrder {
					var done bool
					switch char {
					case 'N':
						if !elves[n] && !elves[nw] && !elves[ne] {
							elfMovement[elf] = n
							proposals[n] = proposals[n] + 1
							done = true
						}
					case 'S':
						if !elves[s] && !elves[sw] && !elves[se] {
							elfMovement[elf] = s
							proposals[s] = proposals[s] + 1
							done = true
						}
					case 'W':
						if !elves[w] && !elves[sw] && !elves[nw] {
							elfMovement[elf] = w
							proposals[w] = proposals[w] + 1
							done = true
						}
					case 'E':
						if !elves[e] && !elves[ne] && !elves[se] {
							elfMovement[elf] = e
							proposals[e] = proposals[e] + 1
							done = true
						}
					}
					if done {
						break
					}
				}

			}
		}
		considerOrder = considerOrder[1:] + considerOrder[:1] // push front char to the back
		// fmt.Println(proposals)

		// make movements
		for elf, movement := range elfMovement {
			if proposals[movement] == 1 {
				elves[elf] = false
				elves[movement] = true
			}
		}
	}

	fmt.Println(calcEmptySpace(elves, elfCount))
	// for elf, exists := range elves {
	// 	if exists {
	// 		fmt.Println(elf)
	// 	}
	// }
}

func Solve2(path string) {
	data := util.GetFileLines(path)
	elves := make(map[util.Point]bool)
	for i, line := range data {
		for j, char := range line {
			if char == '#' {
				elves[util.Point{X: j, Y: i}] = true
			}
		}
	}

	roundCount := 0
	considerOrder := "NSWE"
	// fmt.Println(elves)
	for {
		// make proposals
		elfMovement := make(map[util.Point]util.Point)
		proposals := make(map[util.Point]int)
		for elf, exists := range elves {
			if exists && hasAdjectent(elf, elves) {
				n := util.Point{X: elf.X, Y: elf.Y - 1}
				s := util.Point{X: elf.X, Y: elf.Y + 1}
				w := util.Point{X: elf.X - 1, Y: elf.Y}
				e := util.Point{X: elf.X + 1, Y: elf.Y}
				nw := util.Point{X: elf.X - 1, Y: elf.Y - 1}
				ne := util.Point{X: elf.X + 1, Y: elf.Y - 1}
				sw := util.Point{X: elf.X - 1, Y: elf.Y + 1}
				se := util.Point{X: elf.X + 1, Y: elf.Y + 1}

				for _, char := range considerOrder {
					var done bool
					switch char {
					case 'N':
						if !elves[n] && !elves[nw] && !elves[ne] {
							elfMovement[elf] = n
							proposals[n] = proposals[n] + 1
							done = true
						}
					case 'S':
						if !elves[s] && !elves[sw] && !elves[se] {
							elfMovement[elf] = s
							proposals[s] = proposals[s] + 1
							done = true
						}
					case 'W':
						if !elves[w] && !elves[sw] && !elves[nw] {
							elfMovement[elf] = w
							proposals[w] = proposals[w] + 1
							done = true
						}
					case 'E':
						if !elves[e] && !elves[ne] && !elves[se] {
							elfMovement[elf] = e
							proposals[e] = proposals[e] + 1
							done = true
						}
					}
					if done {
						break
					}
				}

			}
		}
		considerOrder = considerOrder[1:] + considerOrder[:1] // push front char to the back
		// fmt.Println(proposals)

		// make movements
		for elf, movement := range elfMovement {
			if proposals[movement] == 1 {
				elves[elf] = false
				elves[movement] = true
			}
		}
		roundCount++
		if len(proposals) == 0 {
			break
		}
	}

	fmt.Println(roundCount)
}

func hasAdjectent(elf util.Point, elves map[util.Point]bool) bool {
	for _, p := range adjacentPointsPlus(elf) {
		if elves[p] {
			return true
		}
	}
	return false
}

func adjacentPoints(p util.Point) []util.Point {
	neighbours := []util.Point{}
	neighbours = append(neighbours, util.Point{X: p.X, Y: p.Y - 1})
	neighbours = append(neighbours, util.Point{X: p.X, Y: p.Y + 1})
	neighbours = append(neighbours, util.Point{X: p.X - 1, Y: p.Y})
	neighbours = append(neighbours, util.Point{X: p.X + 1, Y: p.Y})
	return neighbours
}

func adjacentPointsPlus(p util.Point) []util.Point {
	neighbours := []util.Point{}
	neighbours = append(neighbours, util.Point{X: p.X, Y: p.Y - 1})
	neighbours = append(neighbours, util.Point{X: p.X, Y: p.Y + 1})
	neighbours = append(neighbours, util.Point{X: p.X - 1, Y: p.Y})
	neighbours = append(neighbours, util.Point{X: p.X + 1, Y: p.Y})
	neighbours = append(neighbours, util.Point{X: p.X - 1, Y: p.Y - 1})
	neighbours = append(neighbours, util.Point{X: p.X + 1, Y: p.Y - 1})
	neighbours = append(neighbours, util.Point{X: p.X - 1, Y: p.Y + 1})
	neighbours = append(neighbours, util.Point{X: p.X + 1, Y: p.Y + 1})
	return neighbours
}

func calcEmptySpace(elves map[util.Point]bool, elfCount int) int {
	var minX, minY, maxX, maxY int
	minX = 99999
	minY = 99999
	maxX = -99999
	maxY = -99999

	count := 0
	for k, v := range elves {
		if v {
			count++
			if k.X < minX {
				minX = k.X
			}
			if k.X > maxX {
				maxX = k.X
			}
			if k.Y < minY {
				minY = k.Y
			}
			if k.Y > maxY {
				maxY = k.Y
			}
		}
	}

	xLength := util.Abs(maxX-minX) + 1
	yLength := util.Abs(maxY-minY) + 1
	rectangleArea := xLength * yLength
	// fmt.Println("minX: " + strconv.Itoa(minX))
	// fmt.Println("minY: " + strconv.Itoa(minY))
	// fmt.Println("maxX: " + strconv.Itoa(maxX))
	// fmt.Println("maxY: " + strconv.Itoa(maxY))
	// fmt.Println("rectangleArea: " + strconv.Itoa(rectangleArea))
	// fmt.Println("elfCount: " + strconv.Itoa(elfCount))
	return rectangleArea - elfCount
}
