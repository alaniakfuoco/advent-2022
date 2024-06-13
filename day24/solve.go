package day24

import (
	"advent/util"
	"fmt"
)

type blizzard struct {
	positon   util.Point
	direction rune
}

func Solve(path string) {
	data := util.GetFileLines(path)
	var start = util.Point{X: 1, Y: 0}
	var goal = util.Point{X: len(data[0]) - 2, Y: len(data) - 1}
	// fmt.Println("goal: " + goal.ToString())

	blizzers := []blizzard{}
	var blizLocations map[util.Point]bool
	for i, line := range data {
		for j, char := range line {
			if char == '^' || char == 'v' || char == '<' || char == '>' {
				newBliz := blizzard{
					positon:   util.Point{X: j, Y: i},
					direction: char,
				}
				blizzers = append(blizzers, newBliz)
				// blizLocations[newBliz.positon] = true
			}
		}
	}

	minute := 0
	choices := []util.Point{start}
	for {
		minute++
		blizLocations = moveBlizzers(blizzers, len(data), len(data[0]))
		// add new traversal options
		for _, choice := range choices {
			choices = choices[1:]
			choices = append(choices, getChoices(choice, blizLocations, goal, len(data[0])-1, len(data)-1)...)
		}
		// remove duplicates
		choices = removeDuplicates(choices)
		// check for end state
		fmt.Println(choices)
		if len(choices) == 0 {
			break
		}
		for _, choice := range choices {
			if choice.Equals(goal) {
				fmt.Println(minute)
				// fmt.Println(choices)
				return
			}
		}

		// fmt.Println(minute)
		// for i := 0; i < len(data); i++ {
		// 	for j := 0; j < len(data[0]); j++ {
		// 		if start.Equals(util.Point{X: j, Y: i}) || goal.Equals(util.Point{X: j, Y: i}) {
		// 			fmt.Print(".")
		// 		} else if i == 0 || j == 0 || i == len(data)-1 || j == len(data[0])-1 {
		// 			fmt.Print("#")
		// 		} else if blizLocations[util.Point{X: j, Y: i}] {
		// 			fmt.Print("x")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }

		// fmt.Println(len(choices))
		// time.Sleep(time.Second)
		// if minute > 18 {
		// 	// fmt.Println(choices)
		// 	return
		// }
		if minute > 10000 {
			fmt.Println(minute)
			return
		}
	}
}

func Solve2(path string) {
	predata := util.GetFileLines(path)
	data := extendInput(predata)
	var start = util.Point{X: 1, Y: 0}
	var goal = util.Point{X: len(data[0]) - 2, Y: len(data) - 1}

	for _, s := range data {
		fmt.Println(s)
	}
	fmt.Println(goal)

	blizzers := []blizzard{}
	walls := make(map[util.Point]bool)
	for i, line := range data {
		for j, char := range line {
			if char == '^' || char == 'v' || char == '<' || char == '>' {
				newBliz := blizzard{
					positon:   util.Point{X: j, Y: i},
					direction: char,
				}
				blizzers = append(blizzers, newBliz)
			} else if char == '#' {
				walls[util.Point{X: j, Y: i}] = true
			}
		}
	}
	height := len(data)
	length := len(data[0])

	minutes, _ := findBestTime(start, goal, blizzers, walls, height, length)
	fmt.Println(minutes)
}

func Solve3(path string) {
	data := util.GetFileLines(path)
	var start = util.Point{X: 1, Y: 0}
	var goal = util.Point{X: len(data[0]) - 2, Y: len(data) - 1}

	// for _, s := range data {
	// 	fmt.Println(s)
	// }
	// fmt.Println(goal)

	blizzers := []blizzard{}
	walls := make(map[util.Point]bool)
	for i, line := range data {
		for j, char := range line {
			if char == '^' || char == 'v' || char == '<' || char == '>' {
				newBliz := blizzard{
					positon:   util.Point{X: j, Y: i},
					direction: char,
				}
				blizzers = append(blizzers, newBliz)
			} else if char == '#' {
				walls[util.Point{X: j, Y: i}] = true
			}
		}
	}
	height := len(data)
	length := len(data[0])

	minutes, blizzers := findBestTime(start, goal, blizzers, walls, height, length)
	minutes1, blizzers := findBestTime(goal, start, blizzers, walls, height, length)
	minutes2, blizzers := findBestTime(start, goal, blizzers, walls, height, length)
	fmt.Println(minutes + minutes1 + minutes2)
}

func findBestTime(start, goal util.Point, blizzers []blizzard, walls map[util.Point]bool, height, length int) (int, []blizzard) {
	minute := 0
	choices := []util.Point{start}
	for {
		minute++
		blizLocations := moveBlizzers(blizzers, height, length)
		// add new traversal options
		for _, choice := range choices {
			choices = choices[1:]
			choices = append(choices, getChoices2(choice, blizLocations, walls, goal, length-1, height-1)...)
		}
		// remove duplicates
		choices = removeDuplicates(choices)
		// check for end state
		if len(choices) == 0 {
			break
		}
		for _, choice := range choices {
			if choice.Equals(goal) {
				return minute, blizzers
			}
		}
		if minute > 10000 {
			return -1, blizzers
		}
	}
	return 0, blizzers
}

func moveBlizzers(blizzards []blizzard, height, length int) map[util.Point]bool {
	blizPos := make(map[util.Point]bool, len(blizzards))
	for i, b := range blizzards {
		switch b.direction {
		case '^':
			b.positon.Y--
			if b.positon.Y == 0 {
				b.positon.Y = height - 2
			}
		case 'v':
			b.positon.Y++
			if b.positon.Y == height-1 {
				b.positon.Y = 1
			}
		case '<':
			b.positon.X--
			if b.positon.X == 0 {
				b.positon.X = length - 2
			}
		case '>':
			b.positon.X++
			if b.positon.X == length-1 {
				b.positon.X = 1
			}
		}
		blizzards[i] = b
		blizPos[b.positon] = true
	}
	return blizPos
}

func getChoices(location util.Point, blizzards map[util.Point]bool, goal util.Point, xLimit, yLimit int) []util.Point {
	result := []util.Point{}
	n := util.Point{X: location.X, Y: location.Y - 1}
	s := util.Point{X: location.X, Y: location.Y + 1}
	w := util.Point{X: location.X - 1, Y: location.Y}
	e := util.Point{X: location.X + 1, Y: location.Y}

	if goal.Equals(n) || goal.Equals(s) || goal.Equals(w) || goal.Equals(e) {
		result = append(result, goal) // if goal is reachable, don't add other options
	} else {
		if !blizzards[location] { // wait
			result = append(result, location)
		}
		if n.Y > 0 && !blizzards[n] { // move north (up)
			result = append(result, n)
		}
		if s.Y < yLimit && !blizzards[s] { // move south (down)
			result = append(result, s)
		}
		if w.X > 0 && !blizzards[w] { // move west (left)
			result = append(result, w)
		}
		if e.X < xLimit && e.Y > 0 && !blizzards[e] { // move east (right)
			result = append(result, e)
		}
	}

	return result
}

func getChoices2(location util.Point, blizzards, walls map[util.Point]bool, goal util.Point, xLimit, yLimit int) []util.Point {
	result := []util.Point{}
	n := util.Point{X: location.X, Y: location.Y - 1}
	s := util.Point{X: location.X, Y: location.Y + 1}
	w := util.Point{X: location.X - 1, Y: location.Y}
	e := util.Point{X: location.X + 1, Y: location.Y}

	if goal.Equals(n) || goal.Equals(s) || goal.Equals(w) || goal.Equals(e) {
		result = append(result, goal) // if goal is reachable, don't add other options
	} else {
		if !blizzards[location] { // wait
			result = append(result, location)
		}
		if n.Y > 0 && !walls[n] && !blizzards[n] { // move north (up)
			result = append(result, n)
		}
		if s.Y < yLimit && !walls[s] && !blizzards[s] { // move south (down)
			result = append(result, s)
		}
		if w.X > 0 && !walls[w] && !blizzards[w] { // move west (left)
			result = append(result, w)
		}
		if e.X < xLimit && !walls[e] && !blizzards[e] { // move east (right)
			result = append(result, e)
		}
	}

	return result
}

func removeDuplicates(points []util.Point) []util.Point {
	uniques := make(map[util.Point]bool)
	for _, p := range points {
		uniques[p] = true
	}
	uniquePoints := []util.Point{}
	for k := range uniques {
		uniquePoints = append(uniquePoints, k)
	}
	return uniquePoints
}

func extendInput(data []string) []string {
	result := make([]string, 0, (3*len(data) - 2))
	for i := 0; i < len(data); i++ {
		result = append(result, data[i])
	}
	for i := 1; i < len(data); i++ {
		str := data[len(data)-(i+1)]
		newstr := []rune{}
		for _, r := range str {
			if r == 'v' {
				newstr = append(newstr, '^')
			} else if r == '^' {
				newstr = append(newstr, 'v')
			} else {
				newstr = append(newstr, r)
			}
		}
		result = append(result, string(newstr))
	}
	for i := 1; i < len(data); i++ {
		result = append(result, data[i])
	}
	return result
}
