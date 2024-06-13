package day2

import (
	"advent/util"
	"fmt"
	"strings"
)

func Solve(path string) {
	data := util.GetFileLines(path)

	score := 0
	for _, line := range data {
		if line != "" {
			pair := strings.Split(line, " ")
			gamePair := getGamePair(pair)
			result := calcGame(gamePair)
			score += result + gamePair[1]
		}
	}

	fmt.Println(score)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	score := 0
	for _, line := range data {
		if line != "" {
			pair := strings.Split(line, " ")
			result := calcGame2(pair)
			// fmt.Print(pair, ": ")
			// fmt.Println(result)
			score += result
		}
	}

	fmt.Println(score)
}

func calcGame(gamePair []int) int {
	if gamePair[0] == gamePair[1] {
		return 3
	}
	if gamePair[1]-gamePair[0] == 1 || gamePair[1]-gamePair[0] == -2 {
		return 6
	}
	return 0
}

func getGamePair(pair []string) []int {
	gamePair := make([]int, 2)
	gamePair[0] = convert(pair[0])
	gamePair[1] = convert(pair[1])
	return gamePair
}

func calcGame2(pair []string) int {
	gamePair := make([]int, 2)
	gamePair[0] = convert(pair[0])
	switch pair[1] {
	case "X": // lose
		gamePair[1] = gamePair[0] - 1
		if gamePair[1] == 0 {
			gamePair[1] = 3
		}
		return gamePair[1]
	case "Y": // tie
		gamePair[1] = gamePair[0]
		return 3 + gamePair[1]
	case "Z": // win
		gamePair[1] = gamePair[0] + 1
		if gamePair[1] > 3 {
			gamePair[1] = 1
		}
		return 6 + gamePair[1]
	}
	return 0
}

/*
ROCK = A|1
PAPER = B|2
SCISSORS = C|3

X = lose
Y = tie
Z = win

1 - 2 = 6 (win)
1 - 3 = 0 (lose)

2 - 1 = 0 (lose)
2 - 3 = 6 (win)

3 - 1 = 6 (win)
3 - 2 = 0 (lose)

A X = 3 + 0 = 3
A Y = 1 + 3 = 4
A Z = 2 + 6 = 8

B X = 1 + 0 = 1
B Y = 2 + 3 = 5
B Z = 3 + 6 = 9

C X = 2 + 0 = 2
C Y = 3 + 3 = 6
C Z = 1 + 6 = 7
*/
func convert(fist string) int {
	switch fist {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}
