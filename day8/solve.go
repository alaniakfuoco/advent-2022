package day8

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	data := util.GetFileLines(path)
	length := len(data[0])
	height := len(data)
	forest := make([][]int, height)
	visible := make([][]bool, height)

	for i, line := range data { // build forest
		treeLine := make([]int, length)
		vl := make([]bool, length)
		for j, r := range line {
			treeLine[j] = numRunetoInt(r)
		}
		if i == 0 || i == height-1 {
			for j, _ := range vl {
				vl[j] = true // all top and bottom edge trees are visible
			}
		} else {
			vl[0] = true
			vl[length-1] = true // all edge trees are visible
		}
		forest[i] = treeLine
		visible[i] = vl
	}

	//printForest(forest)

	for i := range forest {
		if i != 0 && i != len(forest)-1 { // only run for non-edges
			walkRow(&forest, &visible, i)
		}
	}

	for i := range forest[0] {
		if i != 0 && i != len(forest[0])-1 { // only run for non-edges
			walkColumn(&forest, &visible, i)
		}
	}

	fmt.Println(countVisible(visible))
	//printVisible(visible)
}

func Solve2(path string) {
	data := util.GetFileLines(path)
	length := len(data[0])
	height := len(data)
	forest := make([][]int, height)
	visible := make([][]bool, height)

	for i, line := range data { // build forest
		treeLine := make([]int, length)
		vl := make([]bool, length)
		for j, r := range line {
			treeLine[j] = numRunetoInt(r)
		}
		if i == 0 || i == height-1 {
			for j, _ := range vl {
				vl[j] = true // all top and bottom edge trees are visible
			}
		} else {
			vl[0] = true
			vl[length-1] = true // all edge trees are visible
		}
		forest[i] = treeLine
		visible[i] = vl
	}

	maxScenic := 0
	// maxX := 0
	// maxY := 0
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[0])-1; j++ {
			res := calcScenic(forest, i, j)
			if res > maxScenic {
				maxScenic = res
				// maxX = i
				// maxY = j
			}
		}
	}

	fmt.Println(maxScenic)
	// // fmt.Println()
	// // fmt.Println(maxX)
	// // fmt.Println(maxY)
	// // fmt.Println()
	// fmt.Println(lookNorth(forest, 3, 2))
	// fmt.Println(lookSouth(forest, 3, 2))
	// fmt.Println(lookEast(forest, 3, 2))
	// fmt.Println(lookWest(forest, 3, 2))
}

func walkRow(forest *[][]int, visible *[][]bool, index int) {
	// check from the left
	max := (*forest)[index][0]
	for i := 1; i < len((*forest)[index])-1; i++ { // skip edges
		if (*forest)[index][i] > max {
			(*visible)[index][i] = true
			max = (*forest)[index][i]
		}
		if max == 9 {
			break // stop when we see a 9
		}
	}
	// check from the right
	max = (*forest)[index][len((*forest)[index])-1]
	for i := len((*forest)[index]) - 2; i >= 1; i-- { // skip edges
		if (*forest)[index][i] > max {
			(*visible)[index][i] = true
			max = (*forest)[index][i]
		}
		if max == 9 {
			break // stop when we see a 9
		}
	}
}

func walkColumn(forest *[][]int, visible *[][]bool, index int) {
	// check from the top
	max := (*forest)[0][index]
	for i := 1; i < len((*forest))-1; i++ { // skip edges
		if (*forest)[i][index] > max {
			(*visible)[i][index] = true
			max = (*forest)[i][index]
		}
		if max == 9 {
			break // stop when we see a 9
		}
	}
	// check from the right
	max = (*forest)[len(*forest)-1][index]
	for i := len((*forest)) - 2; i >= 1; i-- { // skip edges
		if (*forest)[i][index] > max {
			(*visible)[i][index] = true
			max = (*forest)[i][index]
		}
		if max == 9 {
			break // stop when we see a 9
		}
	}
}

func calcScenic(forest [][]int, x, y int) int {
	north := lookNorth(forest, x, y)
	south := lookSouth(forest, x, y)
	east := lookEast(forest, x, y)
	west := lookWest(forest, x, y)

	return north * south * east * west
}

func lookNorth(forest [][]int, x, y int) int {
	size := forest[x][y]
	treeCount := 0
	for i := x - 1; i >= 0; i-- {
		treeCount++
		if size <= forest[i][y] {
			break // stop counting
		}
	}
	return treeCount
}

func lookSouth(forest [][]int, x, y int) int {
	size := forest[x][y]
	treeCount := 0
	for i := x + 1; i < len(forest); i++ {
		treeCount++
		if size <= forest[i][y] {
			break // stop counting
		}
	}
	return treeCount
}

func lookEast(forest [][]int, x, y int) int {
	size := forest[x][y]
	treeCount := 0
	for i := y + 1; i < len(forest[x]); i++ {
		treeCount++
		if size <= forest[x][i] {
			break // stop counting
		}
	}
	return treeCount
}

func lookWest(forest [][]int, x, y int) int {
	size := forest[x][y]
	treeCount := 0
	for i := y - 1; i >= 0; i-- {
		treeCount++
		if size <= forest[x][i] {
			break // stop counting
		}
	}
	return treeCount
}

func numRunetoInt(r rune) int {
	return int(r) - 48
}

func printForest(f [][]int) {
	for _, a := range f {
		fmt.Println(a)
	}
}

func printVisible(v [][]bool) {
	for _, a := range v {
		for _, b := range a {
			if b {
				fmt.Print("T")
			} else {
				fmt.Print("F")
			}
		}
		fmt.Println()
	}
}

func countVisible(v [][]bool) int {
	count := 0
	for _, a := range v {
		for _, b := range a {
			if b {
				count++
			}
		}
	}
	return count
}
