package day15

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	sensors := make(map[util.Point]bool)
	beacons := make(map[util.Point]bool)
	sobs := make(map[util.Point]util.Point)
	for _, line := range data {
		parts := strings.Split(line, " ")
		//fmt.Println(parts[2], parts[3], parts[8], parts[9])
		// sensor
		s := parseSoB(parts[2], parts[3], false)
		// beacon
		b := parseSoB(parts[8], parts[9], true)

		sensors[s] = true
		beacons[b] = true
		sobs[s] = b
	}

	checkLine := 2000000
	cannotBeacons := make(map[int]bool)
	for s, b := range sobs {
		md := util.ManhatDistance(s, b)
		//fmt.Println(s, b, md)
		dtoline := util.Abs(s.Y - checkLine)
		if dtoline <= md {
			for i := 0; i <= md-dtoline; i++ {
				cannotBeacons[s.X+i] = true
				cannotBeacons[s.X-i] = true
			}
		}
	}

	subCount := 0
	for x := range cannotBeacons {
		if beacons[util.Point{X: x, Y: checkLine}] {
			subCount++
		}
	}

	fmt.Println("cannot:", len(cannotBeacons))
	fmt.Println("is:", subCount)
	//fmt.Println(cannotBeacons)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	sensors := make(map[util.Point]bool)
	beacons := make(map[util.Point]bool)
	sobs := make(map[util.Point]util.Point)
	for _, line := range data {
		parts := strings.Split(line, " ")
		//fmt.Println(parts[2], parts[3], parts[8], parts[9])
		// sensor
		s := parseSoB(parts[2], parts[3], false)
		// beacon
		b := parseSoB(parts[8], parts[9], true)

		sensors[s] = true
		beacons[b] = true
		sobs[s] = b
	}

	//checkLine := 2000000
	for i := 0; i <= 4000000; i++ {
		fmt.Println(i)
		cannotBeacons := make(map[int]bool)
		for s, b := range sobs {
			md := util.ManhatDistance(s, b)
			dtoline := util.Abs(s.Y - i)
			if dtoline <= md {
				for j := 0; j <= md-dtoline; j++ {
					if s.X+j <= 4000000 {
						cannotBeacons[s.X+j] = true
					}
					if s.X-j >= 0 {
						cannotBeacons[s.X-j] = true
					}
				}
			}
		}
		//fmt.Println(len(cannotBeacons))
		if len(cannotBeacons) == 4000000 {
			fmt.Println("y pos:", i)
			for j := 0; j < 4000000; j++ {
				if !cannotBeacons[j] {
					fmt.Println("x pos:", j)
					fmt.Println("answer:", (j*4000000)+i)
				}
			}
			break
		}
	}

}

func Solve3(path string) {
	data := util.GetFileLines(path)

	sensors := make(map[util.Point]bool)
	beacons := make(map[util.Point]bool)
	sobs := make(map[util.Point]util.Point)
	for _, line := range data {
		parts := strings.Split(line, " ")
		//fmt.Println(parts[2], parts[3], parts[8], parts[9])
		// sensor
		s := parseSoB(parts[2], parts[3], false)
		// beacon
		b := parseSoB(parts[8], parts[9], true)

		sensors[s] = true
		beacons[b] = true
		sobs[s] = b
	}

	limit := 4000000
	//for i := 0; i <= limit; i++ {
	for i := 4600; i <= limit; i++ {
		fmt.Println(i)
		for j := 0; j <= limit; j++ {
			outside := false
			for s, b := range sobs {
				md := util.ManhatDistance(s, b)
				td := util.ManhatDistance(s, util.Point{X: i, Y: j})
				if td <= md {
					outside = false
					break
				} else {
					outside = true
				}
			}
			if outside {
				fmt.Println("x pos:", i)
				fmt.Println("y pos:", j)
				fmt.Println("answer:", (i*4000000)+j)
				i = limit + 2
				j = limit + 2
			}
		}
	}
}

func Solve4(path string) {
	data := util.GetFileLines(path)

	sensors := make(map[util.Point]bool)
	beacons := make(map[util.Point]bool)
	sobs := make(map[util.Point]util.Point)
	for _, line := range data {
		parts := strings.Split(line, " ")
		//fmt.Println(parts[2], parts[3], parts[8], parts[9])
		// sensor
		s := parseSoB(parts[2], parts[3], false)
		// beacon
		b := parseSoB(parts[8], parts[9], true)

		sensors[s] = true
		beacons[b] = true
		sobs[s] = b
	}

	limit := 4000000
	//pointsChecked := 0
	die := false
	for s, b := range sobs {
		//fmt.Println(pointsChecked)
		md := util.ManhatDistance(s, b)
		rang := md + 1
		toScan := make([]util.Point, 0, rang*4)
		for i := 0; i < rang; i++ {
			eastNorth := util.Point{X: s.X + (rang - i), Y: s.Y - i}
			southEast := util.Point{X: s.X + i, Y: s.Y + (rang - i)}
			westSouth := util.Point{X: s.X - (rang - i), Y: s.Y + i}
			northWest := util.Point{X: s.X - i, Y: s.Y - (rang - i)}
			if eastNorth.X <= limit && eastNorth.Y >= 0 {
				toScan = append(toScan, eastNorth)
			}
			if southEast.X <= limit && southEast.Y <= limit {
				toScan = append(toScan, southEast)
			}
			if westSouth.X >= 0 && westSouth.Y <= limit {
				toScan = append(toScan, westSouth)
			}
			if northWest.X >= 0 && northWest.Y >= 0 {
				toScan = append(toScan, northWest)
			}
		}

		for _, p := range toScan {
			outside := false
			for s2, b2 := range sobs {
				md := util.ManhatDistance(s2, b2)
				td := util.ManhatDistance(s2, p)
				if td <= md {
					outside = false
					break
				} else {
					outside = true
				}
			}
			if outside {
				fmt.Println("x pos:", p.X)
				fmt.Println("y pos:", p.Y)
				fmt.Println("answer:", (p.X*4000000)+p.Y)
				die = true
				break
			}
		}

		if die {
			break
		}
		//pointsChecked += len(toScan)
	}

}

func parseSoB(sxr, syr string, beacon bool) util.Point {
	sxr = sxr[2 : len(sxr)-1]
	sx, _ := strconv.Atoi(sxr)

	if beacon {
		syr = syr[2:]
	} else {
		syr = syr[2 : len(syr)-1]
	}

	sy, _ := strconv.Atoi(syr)

	return util.Point{X: sx, Y: sy}
}
