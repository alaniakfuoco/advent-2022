package day3

import (
	"advent/util"
	"fmt"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	result := 0
	for _, line := range data {
		if line != "" {
			pair := splitInHalf(line)
			shared := findShared(pair[0], pair[1])
			result += getPriority(shared)
		}
	}

	fmt.Println(result)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	result := 0
	for i := 0; i < len(data); i += 3 {
		lines := []string{data[i], data[i+1], data[i+2]}
		shared := findAllShared(lines[0], lines[1])
		same := findShared(string(shared), lines[2])
		result += getPriority(same)
	}

	fmt.Println(result)
}

func findAllShared(a, b string) []rune {
	finds := []rune{}
	seen := make(map[rune]bool)
	for _, v := range a {
		seen[v] = true
	}

	for _, k := range b {
		_, ok := seen[k]
		if ok {
			finds = append(finds, k)
		}
	}

	return finds
}

func findShared(a, b string) rune {
	seen := make(map[rune]bool)
	for _, v := range a {
		seen[v] = true
	}

	for _, k := range b {
		_, ok := seen[k]
		if ok {
			return k
		}
	}

	return 0
}

func splitInHalf(str string) []string {
	half := len(str) / 2
	pair := make([]string, 2)
	pair[0] = str[:half]
	pair[1] = str[half:]
	return pair
}

func getPriority(r rune) int {
	val := int(r)
	if val > 97 {
		return val - 96
	} else {
		return val - 38
	}
}
