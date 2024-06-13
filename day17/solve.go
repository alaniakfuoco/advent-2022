package day17

import (
	"advent/util"
	"fmt"
)

const (
	LEFT  = 0
	RIGHT = 1
	DOWN  = 2
)

// byte 60 = <
// byte 62 = >
func Solve(path string, x int) {
	movePattern := util.ReadFile(path)

	rocks := make(map[util.Point]bool)
	for _, p := range getFloor() {
		rocks[p] = true
	}

	isBlocked := func(s *shape, direction int) bool {
		if direction == LEFT { // checking left
			for _, p := range s.getAllPoints() {
				check := util.Point{X: p.X - 1, Y: p.Y}
				if rocks[check] {
					return true
				}
			}
		} else if direction == RIGHT { // checking right
			for _, p := range s.getAllPoints() {
				check := util.Point{X: p.X + 1, Y: p.Y}
				if rocks[check] {
					return true
				}
			}
		} else if direction == DOWN { // checking bottom
			for _, p := range s.bots {
				check := util.Point{X: p.X, Y: p.Y - 1}
				if rocks[check] {
					return true
				}
			}
		}
		return false
	}

	var rockCount int
	var maxHeight int
	var pushCount int
	piece := nextRock(rockCount, maxHeight+4)
	//fmt.Println("simulating rock:", rockCount+1)
	for rockCount = 0; rockCount < x; pushCount++ {
		// if rockCount == 2 {
		// 	fmt.Print(piece, " | ")
		// }
		// push left or right
		direction := movePattern[pushCount%len(movePattern)]
		//fmt.Print(string(direction))
		// if rockCount == 2 {
		// 	fmt.Print(" | ", string(direction), " | ")
		// }
		if direction == 60 { // left
			//fmt.Println("pushing left")
			if !isBlocked(piece, LEFT) {
				// if rockCount == 2 {
				// 	fmt.Print(" | going left | ")
				// }
				piece.moveLeft()
			}
		} else if direction == 62 { // right
			//fmt.Println("pushing right")
			if !isBlocked(piece, RIGHT) {
				// if rockCount == 2 {
				// 	fmt.Print(" | going right | ")
				// }
				piece.moveRight()
			}
		}
		// if rockCount == 2 {
		// 	fmt.Println(piece)
		// }
		// fall down or come to rest
		if isBlocked(piece, DOWN) { // piece comes to a rest
			newRocks := piece.getAllPoints()
			for _, p := range newRocks {
				rocks[p] = true
			}
			//fmt.Println(newRocks)
			maxHeight = util.Max(maxHeight, piece.tippyTop)
			rockCount++
			piece = nextRock(rockCount, maxHeight+4)
			//fmt.Println("simulating rock:", rockCount+1, "maxheight:", maxHeight)
			//fmt.Println(piece)

		} else {
			piece.moveDown()
		}
	}

	//fmt.Println(rocks)
	fmt.Println(maxHeight)
}

func Solve2(path string) {
	movePattern := util.ReadFile(path)

	rocks := make(map[util.Point]bool)
	for _, p := range getFloor() {
		rocks[p] = true
	}

	isBlocked := func(s *shape, direction int) bool {
		if direction == LEFT { // checking left
			for _, p := range s.getAllPoints() {
				check := util.Point{X: p.X - 1, Y: p.Y}
				if rocks[check] {
					return true
				}
			}
		} else if direction == RIGHT { // checking right
			for _, p := range s.getAllPoints() {
				check := util.Point{X: p.X + 1, Y: p.Y}
				if rocks[check] {
					return true
				}
			}
		} else if direction == DOWN { // checking bottom
			for _, p := range s.bots {
				check := util.Point{X: p.X, Y: p.Y - 1}
				if rocks[check] {
					return true
				}
			}
		}
		return false
	}

	var rockCount int
	var maxHeight int
	var pushCount int
	piece := nextRock(rockCount, maxHeight+4)
	//combos := make(map[util.Point]int)
	var toCheck util.Point
	var cycle int
	// singles := util.ParsePointsFromFile("./day17/tochecks.txt")
	// fmt.Println(len(singles))
	for rockCount = 0; rockCount < 100000; pushCount++ {
		if rockCount > 631 {
			if rockCount == 632 {
				toCheck.X = rockCount % 5
				toCheck.Y = pushCount % len(movePattern)
			} else {
				a := util.Point{X: rockCount % 5, Y: pushCount % len(movePattern)}
				if a.Equals(toCheck) {
					fmt.Println(rockCount, maxHeight)
					cycle++
					if cycle == 4 {
						break
					}
				}
			}
		}
		// a := util.Point{X: rockCount % 5, Y: pushCount % len(movePattern)}
		// //combos[a] = combos[a] + 1
		// for _, p := range singles {
		// 	if p.Equals(a) {
		// 		fmt.Println(rockCount, a)
		// 	}
		// }

		// push left or right
		direction := movePattern[pushCount%len(movePattern)]
		if direction == 60 { // left
			if !isBlocked(piece, LEFT) {
				piece.moveLeft()
			}
		} else if direction == 62 { // right
			if !isBlocked(piece, RIGHT) {
				piece.moveRight()
			}
		}

		// fall down or come to rest
		if isBlocked(piece, DOWN) { // piece comes to a rest
			newRocks := piece.getAllPoints()
			for _, p := range newRocks {
				rocks[p] = true
			}
			maxHeight = util.Max(maxHeight, piece.tippyTop)
			rockCount++
			piece = nextRock(rockCount, maxHeight+4)

		} else {
			piece.moveDown()
		}
	}

	// for i, c := range combos {
	// 	if c == 1 {
	// 		fmt.Println(i)
	// 	}
	// }
	//fmt.Println(count)
	//fmt.Println(maxHeight)
}

func nextRock(count, height int) *shape {
	switch count % 5 {
	case 0:
		return newFlat(height)
	case 1:
		return newPlus(height)
	case 2:
		return newL(height)
	case 3:
		return newTall(height)
	case 4:
		return newSquare(height)
	}
	return nil
}

func getFloor() []util.Point {
	a := util.Point{X: 1, Y: 0}
	b := util.Point{X: 2, Y: 0}
	c := util.Point{X: 3, Y: 0}
	d := util.Point{X: 4, Y: 0}
	e := util.Point{X: 5, Y: 0}
	f := util.Point{X: 6, Y: 0}
	g := util.Point{X: 7, Y: 0}
	return []util.Point{a, b, c, d, e, f, g}
}

func singles() []util.Point {
	a := util.Point{X: 0, Y: 1}
	b := util.Point{X: 0, Y: 25}
	c := util.Point{X: 0, Y: 14}
	d := util.Point{X: 0, Y: 26}
	e := util.Point{X: 0, Y: 27}

	f := util.Point{X: 1, Y: 5}
	g := util.Point{X: 1, Y: 30}
	h := util.Point{X: 1, Y: 4}
	i := util.Point{X: 1, Y: 18}
	j := util.Point{X: 1, Y: 31}

	k := util.Point{X: 2, Y: 9}
	l := util.Point{X: 2, Y: 36}
	m := util.Point{X: 3, Y: 18}
	n := util.Point{X: 3, Y: 19}
	return []util.Point{a, b, c, d, e, f, g, h, i, j, k, l, m, n}
}
