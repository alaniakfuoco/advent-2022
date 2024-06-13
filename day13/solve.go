package day13

import (
	"advent/util"
	"fmt"
	"reflect"
	"sort"
	"strconv"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	var line1 []interface{}
	var line2 []interface{}
	count := 0
	score := 0
	index := 0
	for i, line := range data {
		//fmt.Println(i + 1)
		switch i % 3 {
		case 0:
			line1 = parseList(line)
		case 1:
			index++
			line2 = parseList(line)

			c := arrayCompare(line1, line2)
			if c == -1 {
				count++
				score += index
			} else if c == 1 {
			} else {
				fmt.Println("this shouldn't happen I think")
			}

		case 2:
			//fmt.Println("count:", count)
		}
	}
	fmt.Println(score)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	// decode1 := make([]interface{}, 0, 1)
	// decode2 := make([]interface{}, 0, 1)
	// decode1 = append(decode1, []int{2})
	// decode2 = append(decode1, []int{6})

	lines := make([][]interface{}, 0, len(data)+2)
	// lines = append(lines, decode1)
	// lines = append(lines, decode2)
	for _, line := range data {
		if line != "" {
			lines = append(lines, parseList(line))
		}
	}

	sort.Slice(lines, func(i, j int) bool {
		return arrayCompare(lines[i], lines[j]) == -1
	})

	// for i, v := range lines {
	// 	if len(v) == 1 {
	// 		if reflect.TypeOf(v[0]).String() == "[]interface {}" {
	// 			arr := v[0].([]interface{})
	// 			if len(arr) == 1 {
	// 				if reflect.TypeOf(arr).String() == "int" {
	// 					num := arr[0].(int)
	// 					if num == 2 || num == 6 {
	// 						fmt.Println("index:", i+1)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	for _, v := range lines {
		fmt.Println(v)
	}

}

func compare(a, b interface{}) int {
	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)

	// fmt.Println(aType.String())
	// fmt.Println(bType.String())

	if aType.Kind() == bType.Kind() {
		if aType.String() == "int" { // both are numbers
			aVal := a.(int)
			bVal := b.(int)
			return util.IntCompare(aVal, bVal)
		} else { // both are arrays
			return arrayCompare(a, b)
		}
	} else { // one is an int
		arr := make([]interface{}, 0, 1)
		if aType.String() == "int" {
			arr = append(arr, a)
			return arrayCompare(arr, b)
		} else {
			arr = append(arr, b)
			return arrayCompare(a, arr)
		}
	}
}

func arrayCompare(a, b interface{}) int {
	aVal := a.([]interface{})
	bVal := b.([]interface{})
	if len(aVal) < len(bVal) {
		for i := range aVal {
			c := compare(aVal[i], bVal[i])
			if c == -1 || c == 1 {
				return c
			}
		}
		return -1
	} else if len(bVal) < len(aVal) {
		for i := range bVal {
			c := compare(aVal[i], bVal[i])
			if c == -1 || c == 1 {
				return c
			}
		}
		return 1
	} else {
		for i := range aVal {
			c := compare(aVal[i], bVal[i])
			if c == -1 || c == 1 {
				return c
			}
		}
		return 0
	}
}

func parseList(s string) []interface{} {
	result := make([]interface{}, 0)
	line := s[1 : len(s)-1] // remove beginning and ending brackets
	nextVal := []rune{}
	chars := []rune(line)
	//fmt.Println("parsing:", line)
	for i := 0; i < len(chars); i++ {
		if chars[i] == '[' {
			backBracketIndex := findNextBacking(line[i:])
			result = append(result, parseList(line[i:i+backBracketIndex+1]))
			i += backBracketIndex
			//fmt.Println("skipping to:", i, string(chars[i]))
		} else if chars[i] == ',' || chars[i] == ']' {
			if len(nextVal) != 0 {
				num, _ := strconv.Atoi(string(nextVal))
				result = append(result, num)
				nextVal = []rune{}
			}
		} else {
			nextVal = append(nextVal, chars[i])
		}
	}
	if len(nextVal) != 0 {
		num, _ := strconv.Atoi(string(nextVal))
		result = append(result, num)
	}
	return result
}

func findNextBacking(s string) int {
	//fmt.Println("processing:", s)
	count := 0
	for i, v := range s {
		if v == '[' {
			count++
		} else if v == ']' {
			count--
		}

		if count == 0 {
			return i
		}
	}
	return -1
}
