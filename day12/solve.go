package day12

import (
	"advent/util"
	"fmt"
	"sort"
)

type node struct {
	loc       util.Point
	heuristic int
	steps     int
}

func Solve(path string) {
	field, s, goal := buildGrid(path)
	height := len(field)
	lenght := len(field[0])

	start := node{
		loc:       s,
		heuristic: 0,
		steps:     0,
	}

	visited := make(map[util.Point]int)
	toSearch := []node{start}
	// found := false
	for len(toSearch) > 0 {
		sort.Slice(toSearch, func(i, j int) bool {
			// if toSearch[i].heuristic == toSearch[j].heuristic {
			// 	return toSearch[i].steps < toSearch[j].steps
			// } else {
			// 	return toSearch[i].heuristic < toSearch[j].heuristic
			// }
			return toSearch[i].heuristic+toSearch[i].steps < toSearch[j].heuristic+toSearch[j].steps
		})
		current := toSearch[0]
		visited[current.loc] = current.steps
		toSearch = toSearch[1:]
		//currentValue := string(field[current.loc.Y][current.loc.X])
		for _, n := range current.loc.GetNeighbours(lenght, height) {
			var inSearch bool
			for _, sv := range toSearch {
				if n.Equals(sv.loc) {
					inSearch = true
					sv.steps = util.Min(current.steps+1, sv.steps)
					break
				}
			}

			canClimb := field[n.Y][n.X] - field[current.loc.Y][current.loc.X]
			if (canClimb <= 1) && !inSearch {
				minSteps := current.steps + 1
				if vstep, vok := visited[n]; vok {
					minSteps = util.Min(vstep, current.steps+1)
					visited[n] = minSteps
				} else {
					newNode := node{n, heuristic(n, goal), minSteps}
					// if newNode.heuristic == 0 {
					// 	fmt.Println(newNode.steps)
					// 	found = true
					// 	break
					// }
					toSearch = append(toSearch, newNode)
				}
			}
		}
	}

	fmt.Println(visited[goal])
	//printPath(visited, lenght, height)
}

func Solve2(path string) {
	field, _, goal := buildGrid(path)
	height := len(field)
	lenght := len(field[0])

	start := node{
		loc:       goal,
		heuristic: 0,
		steps:     0,
	}

	visited := make(map[util.Point]int)
	toSearch := []node{start}
	found := false
	for len(toSearch) > 0 && !found {
		current := toSearch[0]
		visited[current.loc] = current.steps
		toSearch = toSearch[1:]
		for _, n := range current.loc.GetNeighbours(lenght, height) {
			var inSearch bool
			for _, sv := range toSearch {
				if n.Equals(sv.loc) {
					inSearch = true
					sv.steps = util.Min(current.steps+1, sv.steps)
					break
				}
			}

			canClimb := field[current.loc.Y][current.loc.X] - field[n.Y][n.X]
			if (canClimb <= 1) && !inSearch {
				minSteps := current.steps + 1
				if vstep, vok := visited[n]; vok {
					minSteps = util.Min(vstep, current.steps+1)
					visited[n] = minSteps
				} else {
					newNode := node{n, 0, minSteps}
					if field[n.Y][n.X] == 'a' {
						fmt.Println(newNode.steps)
						found = true
						break
					}
					toSearch = append(toSearch, newNode)
				}
			}
		}
	}
}

func buildGrid(path string) ([][]rune, util.Point, util.Point) {
	data := util.GetFileLines(path)

	start := util.Point{}
	goal := util.Point{}

	// build grid
	field := make([][]rune, len(data))
	for i, line := range data {
		fieldLine := make([]rune, len(line))
		for j, r := range line {
			if r == 'S' {
				start.X = j
				start.Y = i
				fieldLine[j] = 'a'
			} else if r == 'E' {
				goal.X = j
				goal.Y = i
				fieldLine[j] = 'z'
			} else {
				fieldLine[j] = r
			}
		}
		field[i] = fieldLine
	}

	return field, start, goal
}

func heuristic(a, b util.Point) int {
	return util.Abs(a.X-b.X) + util.Abs(a.Y-b.Y)
}

//	func nodeCompare(a, b node) bool {
//		return a.heuristic < b.heuristic
//	}

func printPath(points map[util.Point]int, length, height int) {
	field := make([][]int, height)
	for i := range field {
		field[i] = make([]int, length)
	}

	for p, v := range points {
		field[p.Y][p.X] = v
	}

	for _, a := range field {
		fmt.Println(a)
	}
}
