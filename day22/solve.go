package day22

import (
	"advent/util"
	"fmt"
	"strconv"
)

// 32 = " "
// 46 = "."
// 35 = "#"
const (
	WALL = 35
	PATH = 46
	WARP = 32
)

type runner struct {
	row    int
	column int
	facing int
	// 0 = right >
	// 1 = down v
	// 2 = left <
	// 3 = up ^
}

type instruction struct {
	move          int
	turnClockwise bool
}

func Solve(path string) {
	data := util.GetFileLines(path)
	grid := make([][]rune, len(data)-2)
	player := runner{}
	heightLimit := len(grid) - 1
	widthLimit := 0
	for _, line := range data {
		if line == "" {
			break
		}
		widthLimit = util.Max(widthLimit, len(line))
	}
	fmt.Println(widthLimit)

	for i, line := range data {
		if line == "" {
			break
		}
		row := make([]rune, widthLimit)
		for i, r := range line {
			row[i] = r
		}
		for i, r := range row {
			if r == 0 {
				row[i] = 32
			}
		}
		grid[i] = row
	}

	ins := parseInstructions(data[len(data)-1])

	for i, r := range grid[0] {
		if r == PATH {
			player.column = i // starting spot
			break
		}
	}

	visited := make(map[util.Point]int)
	fmt.Println(player)
	visited[util.Point{X: player.row, Y: player.column}] = player.facing
	for _, in := range ins {
		//fmt.Println(i)
		if in.move == 0 { // turn
			player.turn(in.turnClockwise)
		} else { // move in direction
			for rep := 0; rep < in.move; rep++ {
				switch player.facing {
				case 0: // right
					if player.column == widthLimit-1 || grid[player.row][player.column+1] == WARP {
						for i, r := range grid[player.row] {
							var done bool
							switch r {
							case WARP:
								continue
							case WALL:
								done = true
							case PATH:
								player.column = i
								done = true
							}
							if done {
								break
							}
						}
					} else if grid[player.row][player.column+1] == PATH {
						player.column++
					}
				case 1: // down
					if player.row == heightLimit || grid[player.row+1][player.column] == WARP {
						for i := 0; i < len(grid); i++ {
							var done bool
							switch grid[i][player.column] {
							case WARP:
								continue
							case WALL:
								done = true
							case PATH:
								player.row = i
								done = true
							}
							if done {
								break
							}
						}
					} else if grid[player.row+1][player.column] == PATH {
						player.row++
					}
				case 2: // left
					if player.column == 0 || grid[player.row][player.column-1] == WARP {
						for i := len(grid[player.row]) - 1; i >= 0; i-- {
							var done bool
							switch grid[player.row][i] {
							case WARP:
								continue
							case WALL:
								done = true
							case PATH:
								player.column = i
								done = true
							}
							if done {
								break
							}
						}
					} else if grid[player.row][player.column-1] == PATH {
						player.column--
					}
				case 3: // up
					if player.row == 0 || grid[player.row-1][player.column] == WARP {
						for i := len(grid) - 1; i >= 0; i-- {
							var done bool
							switch grid[i][player.column] {
							case WARP:
								continue
							case WALL:
								done = true
							case PATH:
								player.row = i
								done = true
							}
							if done {
								break
							}
						}
					} else if grid[player.row-1][player.column] == PATH {
						player.row--
					}
				}
				visited[util.Point{X: player.row, Y: player.column}] = player.facing
			}
		}
		//fmt.Println(player)
	}

	// for y, row := range grid {
	// 	for x, r := range row {
	// 		if _, ok := visited[util.Point{X: y, Y: x}]; ok {
	// 			switch visited[util.Point{X: y, Y: x}] {
	// 			case 0:
	// 				fmt.Print(">")
	// 			case 1:
	// 				fmt.Print("v")
	// 			case 2:
	// 				fmt.Print("<")
	// 			case 3:
	// 				fmt.Print("^")
	// 			}
	// 		} else if r == 32 {
	// 			fmt.Print("@")
	// 		} else {
	// 			fmt.Print(string(r))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(player)
	result := (1000 * (player.row + 1)) + (4 * (player.column + 1)) + player.facing
	fmt.Println(result)
}

func Solve2(path string) {
	data := util.GetFileLines(path)
	grid := make([][]rune, len(data)-2)
	player := runner{}
	heightLimit := len(grid) - 1
	widthLimit := 0
	for _, line := range data {
		if line == "" {
			break
		}
		widthLimit = util.Max(widthLimit, len(line))
	}
	// fmt.Println(heightLimit)
	// fmt.Println(widthLimit)

	for i, line := range data {
		if line == "" {
			break
		}
		row := make([]rune, widthLimit)
		for i, r := range line {
			row[i] = r
		}
		for i, r := range row {
			if r == 0 {
				row[i] = 32
			}
		}
		grid[i] = row
	}

	// fmt.Println(len(grid))
	// fmt.Println(len(grid[0]))

	ins := parseInstructions(data[len(data)-1])

	for i, r := range grid[0] {
		if r == PATH {
			player.column = i // starting spot
			break
		}
	}

	// fmt.Println(getFaceIndex(runner{25, 111, 0}))
	// fmt.Println(getFaceIndex(runner{51, 111, 0}))
	// fmt.Println(getFaceIndex(runner{51, 80, 0}))
	// fmt.Println(getFaceIndex(runner{140, 10, 0}))
	// fmt.Println(getFaceIndex(runner{130, 75, 0}))
	// fmt.Println(getFaceIndex(runner{155, 3, 0}))
	// fmt.Println(getFaceIndex(runner{165, 3, 0}))

	visited := make(map[util.Point]int)
	// fmt.Println(player)
	// fmt.Println(getFaceIndex(player))
	visited[util.Point{X: player.row, Y: player.column}] = player.facing
	for _, in := range ins {
		//fmt.Println(i)
		if in.move == 0 { // turn
			player.turn(in.turnClockwise)
		} else { // move in direction
			for rep := 0; rep < in.move; rep++ {
				faceIndex := getFaceIndex(player)
				switch player.facing {
				case 0: // right
					if player.column == widthLimit-1 || grid[player.row][player.column+1] == WARP {
						newRow, newCol, newFacing := warpRight(faceIndex, player)
						if grid[newRow][newCol] == WALL {
							break // don't move
						} else { // else set new position
							player.column = newCol
							player.row = newRow
							player.facing = newFacing
						}
					} else if grid[player.row][player.column+1] == PATH {
						player.column++
					}
				case 1: // down
					if player.row == heightLimit || grid[player.row+1][player.column] == WARP {
						newRow, newCol, newFacing := warpDown(faceIndex, player)
						if grid[newRow][newCol] == WALL {
							break // done move
						} else {
							player.column = newCol
							player.row = newRow
							player.facing = newFacing
						}
					} else if grid[player.row+1][player.column] == PATH {
						player.row++
					}
				case 2: // left
					if player.column == 0 || grid[player.row][player.column-1] == WARP {
						newRow, newCol, newFacing := warpLeft(faceIndex, player)
						if grid[newRow][newCol] == WALL {
							break // done move
						} else {
							player.column = newCol
							player.row = newRow
							player.facing = newFacing
						}
					} else if grid[player.row][player.column-1] == PATH {
						player.column--
					}
				case 3: // up
					if player.row == 0 || grid[player.row-1][player.column] == WARP {
						newRow, newCol, newFacing := warpUp(faceIndex, player)
						if grid[newRow][newCol] == WALL {
							break // done move
						} else {
							player.column = newCol
							player.row = newRow
							player.facing = newFacing
						}
					} else if grid[player.row-1][player.column] == PATH {
						player.row--
					}
				}
				visited[util.Point{X: player.row, Y: player.column}] = player.facing
			}
		}
		// fmt.Println(player)
	}

	for y, row := range grid {
		for x, r := range row {
			if _, ok := visited[util.Point{X: y, Y: x}]; ok {
				switch visited[util.Point{X: y, Y: x}] {
				case 0:
					fmt.Print(">")
				case 1:
					fmt.Print("v")
				case 2:
					fmt.Print("<")
				case 3:
					fmt.Print("^")
				}
			} else if r == 32 {
				fmt.Print("@")
			} else {
				fmt.Print(string(r))
			}
		}
		fmt.Println()
	}

	fmt.Println(player)
	result := (1000 * (player.row + 1)) + (4 * (player.column + 1)) + player.facing
	fmt.Println(result)
}

func warpRight(fi int, r runner) (int, int, int) {
	var newRow, newCol, newFace int
	switch fi {
	case 2:
		newFace = 2
		newRow = 149 - r.row
		newCol = 99
	case 3:
		newFace = 3
		newRow = 49
		newCol = 100 + (r.row - 50)
	case 5:
		newFace = 2
		newRow = 149 - r.row
		newCol = 149
	case 6:
		newFace = 3
		newRow = 149
		newCol = r.row - 100
	}
	return newRow, newCol, newFace
}

func warpLeft(fi int, r runner) (int, int, int) {
	var newRow, newCol, newFace int
	switch fi {
	case 1:
		newFace = 0
		newRow = 149 - r.row
		newCol = 0
	case 3:
		newFace = 1
		newRow = 100
		newCol = r.row - 50
	case 4:
		newFace = 0
		newRow = 149 - r.row
		newCol = 50
	case 6:
		newFace = 1
		newRow = 0
		newCol = r.row - 100
	}
	return newRow, newCol, newFace
}

func warpUp(fi int, r runner) (int, int, int) {
	var newRow, newCol, newFace int
	switch fi {
	case 1:
		newFace = 0
		newRow = r.column + 100
		newCol = 0
	case 2:
		newFace = 3
		newRow = 199
		newCol = r.column - 100
	case 4:
		newFace = 0
		newRow = r.column + 50
		newCol = 50
	}
	return newRow, newCol, newFace
}

func warpDown(fi int, r runner) (int, int, int) {
	var newRow, newCol, newFace int
	switch fi {
	case 2:
		newFace = 2
		newRow = r.column - 50
		newCol = 99
	case 5:
		newFace = 2
		newRow = r.column + 50
		newCol = 49
	case 6:
		newFace = 1
		newRow = 0
		newCol = r.column + 100
	}
	return newRow, newCol, newFace
}

func getFaceIndex(r runner) int {
	if r.row < 50 {
		if r.column < 100 {
			return 1
		} else {
			return 2
		}
	} else if r.row < 100 {
		return 3
	} else if r.row < 150 {
		if r.column < 50 {
			return 4
		} else {
			return 5
		}
	} else if r.row < 200 {
		return 6
	}
	return 0
}

func parseInstructions(ins string) []instruction {
	temp := ""
	result := []instruction{}
	for _, r := range ins {
		if r < 76 { // number
			temp += string(r)
		} else {
			if temp != "" {
				v, _ := strconv.Atoi(temp)
				result = append(result, instruction{move: v})
				temp = ""
			}
			if r == 76 { // L
				result = append(result, instruction{turnClockwise: false})
			} else { // R
				result = append(result, instruction{turnClockwise: true})
			}
		}
	}
	if temp != "" {
		v, _ := strconv.Atoi(temp)
		result = append(result, instruction{move: v})
		temp = ""
	}
	return result
}

// func (r *runner) turnLeft(clockwise bool) {
// 	r.turn(false)
// }

// func (r *runner) turnRight(clockwise bool) {
// 	r.turn(true)
// }

func (r *runner) turn(clockwise bool) {
	if clockwise {
		r.facing++
		if r.facing == 4 {
			r.facing = 0
		}
	} else {
		r.facing--
		if r.facing == -1 {
			r.facing = 3
		}
	}
}
