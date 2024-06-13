package day16

import (
	"advent/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type pipe struct {
	code        string
	flowRate    int
	connections []string
	open        bool
}

type choice struct {
	location    string
	time        int
	benefit     int
	nextChoices []*choice
}

type choice2 struct {
	location1   string
	last1       string
	location2   string
	last2       string
	time        int
	benefit     int
	nextChoices []*choice2
	open1       bool
	open2       bool
}

type node struct {
	loc   string
	steps int
}

type actor struct {
	name        string
	location    string
	wait        int
	nextBenefit int
}

var depth int = 0

func Solve(path string) {
	data := util.GetFileLines(path)

	pipes := make(map[string]pipe, len(data))
	for _, line := range data {
		p := parsePipe(line)
		pipes[p.code] = p
	}

	startChoice := choice{
		location: "AA",
		benefit:  0,
		time:     0,
	}
	startChoice.nextChoices = getChoices(startChoice.location, startChoice.time, pipes, 30)

	// fmt.Println(choices)
	// for _, c := range startChoice.nextChoices {
	// 	fmt.Println(c.location, getMaxBenefit2(c))
	// }
	fmt.Println("best:", getMaxBenefit2(&startChoice))
	// fmt.Println()

	// correctPath := []string{"DD", "BB", "JJ", "HH", "EE", "CC"}
	// currChoice := startChoice
	// for i := range correctPath {
	// 	fmt.Println(currChoice)
	// 	for _, c := range currChoice.nextChoices {
	// 		if c.location == correctPath[i] {
	// 			currChoice = *c
	// 		}
	// 	}
	// }
	// fmt.Println(currChoice)
}

func Solve2(path string) {
	data := util.GetFileLines(path)

	pipes := make(map[string]pipe, len(data))
	for _, line := range data {
		p := parsePipe(line)
		pipes[p.code] = p
	}

	startChoice := choice2{
		location1: "AA",
		location2: "AA",
		benefit:   0,
		time:      0,
	}
	fmt.Println(printState("AA", "AA", pipes))
	startChoice.nextChoices = getChoices2(&startChoice, pipes, make(map[string]bool))
	fmt.Println("calculating max benefit")
	fmt.Println("best:", getMaxBenefit3(&startChoice))
}

func Solve3(path string) {
	data := util.GetFileLines(path)

	pipes := make(map[string]pipe, len(data))
	for _, line := range data {
		p := parsePipe(line)
		pipes[p.code] = p
	}

	closedCount := 0
	for _, p := range pipes {
		if !p.open {
			closedCount++
		}
	}

	for i := 0; i < util.Power(2, closedCount+1)-2; i++ {

	}
}

// D H E
// J B C
func Solve4(path string) {
	data := util.GetFileLines(path)

	pipes := make(map[string]pipe, len(data))
	for _, line := range data {
		p := parsePipe(line)
		pipes[p.code] = p
	}

	actor1 := actor{"y", "AA", 0, 0}
	actor2 := actor{"e", "AA", 0, 0}
	totalPressure := 0
	act := func(a *actor, time int) {
		if a.wait == 0 {
			totalPressure += a.nextBenefit
			// if a.nextBenefit != 0 {
			// 	fmt.Println("adding", actor1.nextBenefit)
			// }
			a.nextBenefit = 0
			var bestChoice choice
			bestHeuristic := -10000000
			choices := getNextChoices(a.location, 0, pipes)
			biggestFlow := 0
			for _, p := range pipes {
				if p.flowRate > biggestFlow {
					biggestFlow = p.flowRate
				}
			}
			for _, ch := range choices {
				heuristic := ch.benefit - (ch.time * biggestFlow)
				if bestHeuristic < heuristic {
					bestChoice = *ch
					bestHeuristic = heuristic
				}
			}
			//fmt.Println(len(choices))
			if bestChoice.benefit != 0 {
				//fmt.Println(a.name, ":", time, bestChoice)
				a.location = bestChoice.location
				a.nextBenefit = pipes[bestChoice.location].flowRate * (26 - (time + (bestChoice.time)))
				p := pipes[bestChoice.location]
				p.open = true
				pipes[bestChoice.location] = p
				a.wait = bestChoice.time - 1
			}
		} else {
			a.wait--
			//fmt.Println("actor1 waiting:", actor1.wait, i)
		}
	}
	for i := 0; i < 26; i++ {
		act(&actor1, i)
		act(&actor2, i)
	}

	// fmt.Println("actor1", actor1)
	// fmt.Println("actor2", actor2)
	fmt.Println(totalPressure)
}

// D H E
// J B C
func Solve5(path string) {
	data := util.GetFileLines(path)

	pipes := make(map[string]pipe, len(data))
	for _, line := range data {
		p := parsePipe(line)
		pipes[p.code] = p
	}

	closedPipes := []string{}
	for _, p := range pipes {
		if !p.open {
			closedPipes = append(closedPipes, p.code)
		}
	}

	combos := util.Combinations(closedPipes, len(closedPipes)/2)
	maxBenefit := 0
	fmt.Println("# of combos:", len(combos))
	for i, c := range combos {
		// if i == 2 {
		fmt.Println("checking combo", (i + 1))
		pipes1 := make(map[string]pipe, len(data))
		pipes2 := make(map[string]pipe, len(data))
		for k, v := range pipes {
			cv := pipe{
				v.code,
				v.flowRate,
				v.connections,
				true,
			}
			pipes1[k] = cv // treat all pipes as open
			pipes2[k] = v

		}
		for _, v := range c {
			p1 := pipes1[v]
			p1.open = false
			pipes1[v] = p1
			p2 := pipes2[v]
			p2.open = true
			pipes2[v] = p2
		}

		startChoice1 := choice{
			location: "AA",
			benefit:  0,
			time:     0,
		}
		startChoice2 := choice{
			location: "AA",
			benefit:  0,
			time:     0,
		}
		// fmt.Println("pipes1", pipes1)
		// fmt.Println("pipes2", pipes2)
		startChoice1.nextChoices = getChoices(startChoice1.location, startChoice1.time, pipes1, 26)
		startChoice2.nextChoices = getChoices(startChoice2.location, startChoice2.time, pipes2, 26)
		b1 := getMaxBenefit2(&startChoice1)
		b2 := getMaxBenefit2(&startChoice2)
		//fmt.Println(b1, b2)
		bestBenfit := b1 + b2
		if bestBenfit > maxBenefit {
			maxBenefit = bestBenfit
			//fmt.Println("current best combo:", c)
		}
		//}
	}

	fmt.Println(maxBenefit)
}

func getChoices(start string, currentSteps int, pipes map[string]pipe, totalTime int) []*choice {
	//time.Sleep(2 * time.Second)
	toSearch := []node{{start, currentSteps}}
	visited := make(map[string]int)
	choices := []*choice{}
	for len(toSearch) > 0 {
		current := toSearch[0]
		visited[current.loc] = current.steps
		toSearch = toSearch[1:]
		nextSteps := current.steps + 1
		for _, s := range pipes[current.loc].connections {
			nextPipe := pipes[s]
			_, inVisited := visited[s]
			if !inVisited && nextSteps < (totalTime-1) {
				if !nextPipe.open {
					choices = append(choices, &choice{
						location: nextPipe.code,
						time:     nextSteps + 1,
						benefit:  nextPipe.flowRate * (totalTime - (nextSteps + 1)),
					})
					toSearch = append(toSearch, node{s, nextSteps})
					visited[s] = nextSteps + 1
				} else {
					toSearch = append(toSearch, node{s, nextSteps})
					visited[s] = nextSteps
				}
			}
		}
	}
	closedCount := 0
	for _, p := range pipes {
		if !p.open {
			closedCount++
		}
	}
	if closedCount > 0 && currentSteps < 29 {
		for _, c := range choices {
			nextPipes := make(map[string]pipe, len(pipes))
			for k, v := range pipes {
				nextPipes[k] = v
			}
			p := nextPipes[c.location]
			p.open = true
			nextPipes[c.location] = p
			//fmt.Println(c.location, nextPipes)
			c.nextChoices = getChoices(c.location, c.time, nextPipes, totalTime)
		}
	}
	return choices
}

func getNextChoices(start string, currentSteps int, pipes map[string]pipe) []*choice {
	toSearch := []node{{start, currentSteps}}
	visited := make(map[string]int)
	choices := []*choice{}
	for len(toSearch) > 0 {
		current := toSearch[0]
		visited[current.loc] = current.steps
		toSearch = toSearch[1:]
		nextSteps := current.steps + 1
		for _, s := range pipes[current.loc].connections {
			nextPipe := pipes[s]
			_, inVisited := visited[s]
			if !inVisited && nextSteps < 25 {
				if !nextPipe.open {
					choices = append(choices, &choice{
						location: nextPipe.code,
						time:     nextSteps + 1,
						benefit:  nextPipe.flowRate * (26 - (nextSteps + 1)),
					})
					toSearch = append(toSearch, node{s, nextSteps})
					visited[s] = nextSteps + 1
				} else {
					toSearch = append(toSearch, node{s, nextSteps})
					visited[s] = nextSteps
				}
			}
		}
	}
	return choices
}

func getChoices2(start *choice2, pipes map[string]pipe, states map[string]bool) []*choice2 {
	depth++
	choices := []*choice2{}
	currentStep := start.time + 1
	currentState := printState(start.location1, start.location2, pipes)
	fmt.Println("depth:", depth, "current step:", currentStep, "currentState:", currentState)
	// if all pipes are open, stop searching
	closedCount := 0
	for _, p := range pipes {
		if !p.open {
			closedCount++
		}
	}
	if closedCount == 0 || currentStep > 26 || states[currentState] {
		return choices
	}
	states[currentState] = true // ensure we are not calculating the same states over and over

	// list out all movement options
	for _, conn1 := range pipes[start.location1].connections {
		for _, conn2 := range pipes[start.location2].connections {
			if start.last1 != conn1 && start.last2 != conn2 {
				choices = append(choices, &choice2{
					location1: conn1,
					location2: conn2,
					time:      currentStep,
					benefit:   0,
					open1:     false,
					open2:     false,
				})
			}
		}
	}

	// list out all open actions
	if !pipes[start.location1].open {
		for _, conn2 := range pipes[start.location2].connections {
			choices = append(choices, &choice2{
				location1: start.location1,
				location2: conn2,
				time:      currentStep,
				benefit:   pipes[start.location1].flowRate * (27 - currentStep),
				open1:     true,
				open2:     false,
			})
		}
	}
	if !pipes[start.location2].open {
		for _, conn1 := range pipes[start.location1].connections {
			choices = append(choices, &choice2{
				location1: conn1,
				location2: start.location2,
				time:      currentStep,
				benefit:   pipes[start.location2].flowRate * (27 - currentStep),
				open1:     false,
				open2:     true,
			})
		}
	}
	if !pipes[start.location1].open && !pipes[start.location2].open {
		choices = append(choices, &choice2{
			location1: start.location1,
			location2: start.location2,
			time:      currentStep,
			benefit:   (pipes[start.location1].flowRate + pipes[start.location2].flowRate) * (27 - currentStep),
			open1:     true,
			open2:     true,
		})
	}

	for _, ch := range choices {
		nextPipes := make(map[string]pipe, len(pipes))
		for k, v := range pipes {
			nextPipes[k] = v
		}
		if ch.open1 {
			p := nextPipes[ch.location1]
			p.open = true
			nextPipes[ch.location1] = p
		}
		if ch.open2 {
			p := nextPipes[ch.location2]
			p.open = true
			nextPipes[ch.location2] = p
		}
		ch.nextChoices = getChoices2(ch, nextPipes, states)
	}

	return choices
}

func printState(loc1, loc2 string, pipes map[string]pipe) string {
	state := loc1 + loc2
	pipenames := []string{}
	for k := range pipes {
		pipenames = append(pipenames, k)
	}
	sort.Slice(pipenames, func(i, j int) bool {
		return strings.Compare(pipenames[i], pipenames[j]) == -1
	})
	for _, p := range pipenames {
		if pipes[p].open {
			state += "T"
		} else {
			state += "F"
		}
	}
	return state
}

func getMaxBenefit3(c *choice2) int {
	maxBenefit := c.benefit

	bestSum := 0
	for _, ch := range c.nextChoices {
		sum := getMaxBenefit3(ch)
		if sum > bestSum {
			bestSum = sum
		}
	}

	return maxBenefit + bestSum
}

func getMaxBenefit2(c *choice) int {
	maxBenefit := c.benefit

	bestSum := 0
	for _, ch := range c.nextChoices {
		sum := getMaxBenefit2(ch)
		if sum > bestSum {
			bestSum = sum
		}
	}

	return maxBenefit + bestSum
}

func getMaxBenefit(c *choice) int {
	maxBenefit := c.benefit
	if len(c.nextChoices) == 0 {
		return maxBenefit
	}
	var nextBest *choice
	for _, ch := range c.nextChoices {
		if nextBest == nil {
			nextBest = ch
		} else {
			if ch.benefit > nextBest.benefit {
				nextBest = ch
			}
		}
	}
	maxBenefit += getMaxBenefit(nextBest)
	return maxBenefit
}

func getBenefits(c *choice) []int {
	benefits := []int{c.benefit}
	if len(c.nextChoices) == 0 {
		return benefits
	}
	var nextBest *choice
	for _, ch := range c.nextChoices {
		if nextBest == nil {
			nextBest = ch
		} else {
			if ch.benefit > nextBest.benefit {
				nextBest = ch
			}
		}
	}
	benefits = append(benefits, getBenefits(nextBest)...)
	return benefits
}

func parsePipe(s string) pipe {
	parts := strings.Split(s, " ")
	cd := parts[1]
	frr := parts[4][5 : len(parts[4])-1]
	fr, _ := strconv.Atoi(frr)
	cons := []string{}
	for i := 9; i < len(parts); i++ {
		cons = append(cons, parts[i][:2])
	}

	return pipe{
		code:        cd,
		flowRate:    fr,
		connections: cons,
		open:        fr == 0, // treat pipes with flow rate 0 as open
	}
}
