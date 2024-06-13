package util

import (
	"strconv"
	"strings"
)

type Point3 struct {
	X int
	Y int
	Z int
}

func (p Point3) Equals(b Point3) bool {
	if p.X == b.X && p.Y == b.Y && p.Z == b.Z {
		return true
	}
	return false
}

func ManhatDistance3(a, b Point3) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y) + Abs(a.Z-b.Z)
}

func (p Point3) ToString() string {
	x := strconv.Itoa(p.X)
	y := strconv.Itoa(p.Y)
	z := strconv.Itoa(p.Z)
	return "x:" + x + " y:" + y + "z:" + z
}

func ParsePoint3sFromFile(path string) []Point3 {
	data := GetFileLines(path)
	result := []Point3{}
	for _, line := range data {
		if line != "" {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			z, _ := strconv.Atoi(parts[2])
			result = append(result, Point3{X: x, Y: y, Z: z})
		}
	}
	return result
}
