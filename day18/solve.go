package day18

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	lava := util.ParsePoint3sFromFile(path)
	sides := make([]int, len(lava))
	for i := range sides {
		sides[i] = 6
	}

	for i := 0; i < len(lava); i++ {
		for j := i; j < len(lava); j++ {
			if util.ManhatDistance3(lava[i], lava[j]) == 1 {
				sides[i] = sides[i] - 1
				sides[j] = sides[j] - 1
				if sides[i] == 0 {
					break
				}
			}
		}
	}

	sideCount := 0
	for i := range sides {
		sideCount += sides[i]
		//fmt.Println(lava[i], sides[i])
	}

	fmt.Println(sideCount)
}

func Solve2(path string) {
	lava := util.ParsePoint3sFromFile(path)
	sides := make([]int, len(lava))
	for i := range sides {
		sides[i] = 6
	}

	maxX := 0
	maxY := 0
	maxZ := 0
	lavaMap := make(map[util.Point3]bool)
	for i := range lava {
		maxX = util.Max(lava[i].X, maxX)
		maxY = util.Max(lava[i].Y, maxY)
		maxZ = util.Max(lava[i].Z, maxZ)
		lavaMap[lava[i]] = true
	}

	airBlocks := make([]util.Point3, 0)
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			for k := 0; k <= maxZ; k++ {
				next := util.Point3{X: i, Y: j, Z: k}
				if !lavaMap[next] {
					if !isOutside(next, lavaMap, maxX, maxY, maxZ) {
						airBlocks = append(airBlocks, next)
					}
				}
			}
		}
	}

	for i := 0; i < len(lava); i++ {
		for j := i; j < len(lava); j++ {
			if util.ManhatDistance3(lava[i], lava[j]) == 1 {
				sides[i] = sides[i] - 1
				sides[j] = sides[j] - 1
				if sides[i] == 0 {
					break
				}
			}
		}
		for j := range airBlocks {
			if util.ManhatDistance3(lava[i], airBlocks[j]) == 1 {
				sides[i] = sides[i] - 1
				if sides[i] == 0 {
					break
				}
			}
		}
	}

	sideCount := 0
	for i := range sides {
		sideCount += sides[i]
	}

	fmt.Println(airBlocks)
	fmt.Println(sideCount)
}

func Solve3(path string) {
	lava := util.ParsePoint3sFromFile(path)

	maxX := 0
	maxY := 0
	maxZ := 0
	lavaCells := make(map[util.Point3]bool)
	for i := range lava {
		maxX = util.Max(lava[i].X, maxX)
		maxY = util.Max(lava[i].Y, maxY)
		maxZ = util.Max(lava[i].Z, maxZ)
		lavaCells[lava[i]] = true
	}

	contains := func(p util.Point3, points []util.Point3) bool {
		for _, i := range points {
			if p.Equals(i) {
				return true
			}
		}
		return false
	}

	steamCells := make(map[util.Point3]bool)
	toSearch := []util.Point3{{X: -1, Y: -1, Z: -1}}
	for len(toSearch) > 0 {
		current := toSearch[0]
		steamCells[current] = true
		toSearch = toSearch[1:]
		for _, p := range getNeighbours(current, maxX, maxY, maxZ) {
			if !lavaCells[p] && !steamCells[p] && !contains(p, toSearch) {
				toSearch = append(toSearch, p)
			}
		}
		//fmt.Println(len(toSearch), len(steamCells))
	}

	sideCount := 0
	for _, p := range lava {
		for _, s := range getNeighbours(p, maxX, maxY, maxZ) {
			if steamCells[s] {
				sideCount++
			}
		}
	}
	// for s := range steamCells {
	// 	fmt.Print(s, " ")
	// }

	//fmt.Println()
	fmt.Println(sideCount)

}

func getNeighbours(p util.Point3, xlimit, ylimit, zlimit int) []util.Point3 {
	result := make([]util.Point3, 0, 6)
	if p.X > -1 {
		result = append(result, util.Point3{X: p.X - 1, Y: p.Y, Z: p.Z})
	}
	if p.Y > -1 {
		result = append(result, util.Point3{X: p.X, Y: p.Y - 1, Z: p.Z})
	}
	if p.Z > -1 {
		result = append(result, util.Point3{X: p.X, Y: p.Y, Z: p.Z - 1})
	}
	if p.X <= xlimit {
		result = append(result, util.Point3{X: p.X + 1, Y: p.Y, Z: p.Z})
	}
	if p.Y <= ylimit {
		result = append(result, util.Point3{X: p.X, Y: p.Y + 1, Z: p.Z})
	}
	if p.Z <= zlimit {
		result = append(result, util.Point3{X: p.X, Y: p.Y, Z: p.Z + 1})
	}
	return result
}

func isOutside(start util.Point3, blocks map[util.Point3]bool, xlimit, ylimit, zlimit int) bool {
	if outsideNorth(start, blocks, ylimit) {
		return true
	}
	if outsideSouth(start, blocks, 0) {
		return true
	}
	if outsideEast(start, blocks, xlimit) {
		return true
	}
	if outsideWest(start, blocks, 0) {
		return true
	}
	if outsideDown(start, blocks, 0) {
		return true
	}
	if outsideUp(start, blocks, zlimit) {
		return true
	}
	return false
}

func outsideNorth(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.Y + 1; i <= limit; i++ {
		next := util.Point3{X: start.X, Y: i, Z: start.Z}
		if blocks[next] {
			return false
		}
	}
	return true
}

func outsideSouth(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.Y - 1; i >= limit; i-- {
		next := util.Point3{X: start.X, Y: i, Z: start.Z}
		if blocks[next] {
			return false
		}
	}
	return true
}

func outsideEast(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.X + 1; i <= limit; i++ {
		next := util.Point3{X: i, Y: start.Y, Z: start.Z}
		if blocks[next] {
			return false
		}
	}
	return true
}

func outsideWest(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.X - 1; i >= limit; i-- {
		next := util.Point3{X: i, Y: start.Y, Z: start.Z}
		if blocks[next] {
			return false
		}
	}
	return true
}

func outsideUp(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.Z + 1; i <= limit; i++ {
		next := util.Point3{X: start.X, Y: start.Y, Z: i}
		if blocks[next] {
			return false
		}
	}
	return true
}

func outsideDown(start util.Point3, blocks map[util.Point3]bool, limit int) bool {
	for i := start.Z - 1; i >= limit; i-- {
		next := util.Point3{X: start.X, Y: start.Y, Z: i}
		if blocks[next] {
			return false
		}
	}
	return true
}
