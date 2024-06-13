package util

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func (p Point) CalcFloat() float32 {
	result := float32(p.X)
	radix := float32(p.Y)

	for radix > 1 {
		radix = radix / 10
	}

	result += radix
	return result
}

func (p Point) ToString() string {
	x := strconv.Itoa(p.X)
	y := strconv.Itoa(p.Y)
	return "x:" + x + " y:" + y
}

func (p Point) Equals(b Point) bool {
	if p.X == b.X && p.Y == b.Y {
		return true
	}
	return false
}

func (p Point) GetNeighbours(xLimit, yLimit int) []Point {
	result := []Point{}
	n := Point{p.X, p.Y - 1}
	s := Point{p.X, p.Y + 1}
	w := Point{p.X - 1, p.Y}
	e := Point{p.X + 1, p.Y}

	// add neighbours only if in range
	if n.Y >= 0 {
		result = append(result, n)
	}
	if s.Y < yLimit {
		result = append(result, s)
	}
	if w.X >= 0 {
		result = append(result, w)
	}
	if e.X < xLimit {
		result = append(result, e)
	}

	return result
}

func ManhatDistance(a, b Point) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}

func ParsePointsFromFile(path string) []Point {
	data := GetFileLines(path)
	result := []Point{}
	for _, line := range data {
		if line != "" {
			s := line[1 : len(line)-1]
			parts := strings.Split(s, " ")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			result = append(result, Point{X: x, Y: y})
		}
	}
	return result
}
