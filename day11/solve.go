package day11

import (
	"advent/util"
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	items           []int
	inspectFunction func(int) int
	throwTestValue  int
	trustMonkey     *monkey
	distrustMonkey  *monkey
	inspectCount    int
}

type bmonkey struct {
	items           []int
	inspectFunction func(int) int
	throwTestValue  int
	trustMonkey     *bmonkey
	distrustMonkey  *bmonkey
	inspectCount    int
	modTrackers     []*modTracker
}

type modTracker struct {
	modChecks   []int
	currentMods []int
}

func Solve(path string) {
	monkeys := parseMonkeys(path)

	rounds := 20
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				if item == 0 {
					// do nothing
				} else {
					worryLevel := m.inspectFunction(item)
					if worryLevel%m.throwTestValue == 0 {
						m.trustMonkey.items = append(m.trustMonkey.items, worryLevel)
					} else {
						m.distrustMonkey.items = append(m.distrustMonkey.items, worryLevel)
					}
					m.inspectCount++
					m.items = make([]int, 0)
				}
			}
			//fmt.Println(m)
		}
	}

	for _, m := range monkeys {
		//fmt.Println(m)
		fmt.Println(m.inspectCount)
	}
}

func Solve2(path string) {
	monkeys := parseBMonkeys(path)

	rounds := 10000
	for i := 0; i < rounds; i++ {
		for mi, m := range monkeys {
			for _, item := range m.modTrackers {
				//fmt.Println("before", mi, item)
				for j := range item.currentMods {
					item.currentMods[j] = m.inspectFunction(item.currentMods[j]) % item.modChecks[j]
				}
				//fmt.Println("after", mi, item)
				if item.currentMods[mi] == 0 {
					m.trustMonkey.modTrackers = append(m.trustMonkey.modTrackers, item)
				} else {
					m.distrustMonkey.modTrackers = append(m.distrustMonkey.modTrackers, item)
				}
				m.inspectCount++
				m.modTrackers = make([]*modTracker, 0)

			}
		}
	}

	for _, m := range monkeys {
		fmt.Println(m.inspectCount)
	}
}

func parseMonkeys(path string) []*monkey {
	data := util.GetFileLines(path)

	var monkeyCount int = (len(data) + 1) / 7
	monkeyCounter := 0
	monkeys := make([]*monkey, monkeyCount)
	for i := range monkeys {
		monkeys[i] = &monkey{}
	}

	for i, line := range data {
		switch (i + 1) % 7 {
		case 1:
			// do nothing
		case 2:
			items := strings.Split(line[18:], " ")
			for _, it := range items {
				itv, _ := strconv.Atoi(it[:2])
				monkeys[monkeyCounter].items = append(monkeys[monkeyCounter].items, itv)
			}
		case 3:
			if strings.HasSuffix(line, "d") { // square function
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x * x
				}
			} else if strings.HasPrefix(line[23:], "+") { // addition
				value, _ := strconv.Atoi(line[25:])
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x + value
				}
			} else { // multiplication
				value, _ := strconv.Atoi(line[25:])
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x * value
				}
			}
		case 4:
			value, _ := strconv.Atoi(line[21:])
			monkeys[monkeyCounter].throwTestValue = value
		case 5:
			value, _ := strconv.Atoi(line[29:])
			monkeys[monkeyCounter].trustMonkey = monkeys[value]
		case 6:
			value, _ := strconv.Atoi(line[30:])
			monkeys[monkeyCounter].distrustMonkey = monkeys[value]
		case 0:
			monkeyCounter++
		}
	}

	return monkeys
}

func parseBMonkeys(path string) []*bmonkey {
	data := util.GetFileLines(path)

	var monkeyCount int = (len(data) + 1) / 7
	monkeyCounter := 0
	monkeys := make([]*bmonkey, monkeyCount)
	for i := range monkeys {
		monkeys[i] = &bmonkey{}
	}

	for i, line := range data {
		switch (i + 1) % 7 {
		case 1:
			// do nothing
		case 2:
			items := strings.Split(line[18:], " ")
			for _, it := range items {
				itv, _ := strconv.Atoi(it[:2])
				monkeys[monkeyCounter].items = append(monkeys[monkeyCounter].items, itv)
			}
		case 3:
			if strings.HasSuffix(line, "d") { // square function
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x * x
				}
			} else if strings.HasPrefix(line[23:], "+") { // addition
				value, _ := strconv.Atoi(line[25:])
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x + value
				}
			} else { // multiplication
				value, _ := strconv.Atoi(line[25:])
				monkeys[monkeyCounter].inspectFunction = func(x int) int {
					return x * value
				}
			}
		case 4:
			value, _ := strconv.Atoi(line[21:])
			monkeys[monkeyCounter].throwTestValue = value
		case 5:
			value, _ := strconv.Atoi(line[29:])
			monkeys[monkeyCounter].trustMonkey = monkeys[value]
		case 6:
			value, _ := strconv.Atoi(line[30:])
			monkeys[monkeyCounter].distrustMonkey = monkeys[value]
		case 0:
			monkeyCounter++
		}
	}

	mods := []int{}
	for _, m := range monkeys {
		mods = append(mods, m.throwTestValue)
	}

	for _, m := range monkeys {
		m.modTrackers = make([]*modTracker, 0)
		for _, itm := range m.items {
			cms := make([]int, len(monkeys))
			for i := range cms {
				cms[i] = itm % mods[i]
			}
			mt := modTracker{
				modChecks:   mods,
				currentMods: cms,
			}

			m.modTrackers = append(m.modTrackers, &mt)
		}

	}

	return monkeys
}
