package day21

import (
	"advent/util"
	"fmt"
	"strconv"
)

type monkey struct {
	name     string
	result   int
	f1       string
	f2       string
	operator string
	parent   *monkey
	answer   func(m1, m2 *monkey) int
}

func Solve(path string) {
	data := util.GetFileLines(path)

	mons := make(map[string]*monkey, len(data))
	for _, line := range data {
		mon := monkey{}
		mon.name = line[0:4]
		if len(line) == 17 { // function monkey
			mon.f1 = line[6:10]
			mon.f2 = line[13:17]
			switch line[11:12] {
			case "+":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) + m2.getResult(mons)
				}
			case "-":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) - m2.getResult(mons)
				}
			case "/":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) / m2.getResult(mons)
				}
			case "*":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) * m2.getResult(mons)
				}
			}
		} else {
			v, _ := strconv.Atoi(line[6:])
			mon.result = v
		}
		mons[mon.name] = &mon
	}

	fmt.Println(mons["root"].getResult(mons))
}

// takes too long
func Solve2(path string) {
	data := util.GetFileLines(path)

	mons := make(map[string]*monkey, len(data))
	for _, line := range data {
		mon := monkey{}
		mon.name = line[0:4]
		if len(line) == 17 { // function monkey
			mon.f1 = line[6:10]
			mon.f2 = line[13:17]
			switch line[11:12] {
			case "+":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) + m2.getResult(mons)
				}
			case "-":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) - m2.getResult(mons)
				}
			case "/":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) / m2.getResult(mons)
				}
			case "*":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) * m2.getResult(mons)
				}
			}
		} else {
			v, _ := strconv.Atoi(line[6:])
			mon.result = v
		}
		mons[mon.name] = &mon
	}

	for _, m := range mons {
		if m.f1 != "" {
			child1 := mons[m.f1]
			child2 := mons[m.f2]
			child1.parent = m
			child2.parent = m
		}
	}

	humaner := mons["humn"]
	for humaner.parent.name != "root" {
		humaner = humaner.parent
	}
	fmt.Println(humaner.name)

	root := mons["root"]
	var nonhumaner *monkey
	if root.f1 == humaner.name {
		nonhumaner = mons[root.f2]
	} else {
		nonhumaner = mons[root.f1]
	}
	fmt.Println(nonhumaner.name)
	otherResult := nonhumaner.getResult(mons)
	fmt.Println(otherResult)

	human := mons["humn"]
	for i := 1; i < 1000000; i++ {
		human.result = i
		// for m := human.parent; m.parent.name != "root"; m = m.parent {
		// 	m.result = 0
		// }
		humanerResult := humaner.getResult(mons)
		// if i == 301 {
		// 	fmt.Println(humaner.getResult(mons))
		// }
		if humanerResult == otherResult {
			fmt.Println(i)
			break
		}
	}
}

func Solve3(path string) {
	data := util.GetFileLines(path)

	mons := make(map[string]*monkey, len(data))
	for _, line := range data {
		mon := monkey{}
		mon.name = line[0:4]
		if len(line) == 17 { // function monkey
			mon.f1 = line[6:10]
			mon.f2 = line[13:17]
			mon.operator = line[11:12]
			switch line[11:12] {
			case "+":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) + m2.getResult(mons)
				}
			case "-":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) - m2.getResult(mons)
				}
			case "/":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) / m2.getResult(mons)
				}
			case "*":
				mon.answer = func(m1, m2 *monkey) int {
					return m1.getResult(mons) * m2.getResult(mons)
				}
			}
		} else {
			v, _ := strconv.Atoi(line[6:])
			mon.result = v
		}
		mons[mon.name] = &mon
	}

	for _, m := range mons {
		if m.f1 != "" {
			child1 := mons[m.f1]
			child2 := mons[m.f2]
			child1.parent = m
			child2.parent = m
		}
	}

	mons["root"].operator = "="
	humanAnswer := 0
	for m := mons["root"]; m.f1 != ""; {
		humaner := mons["humn"]
		// var nonhumaner *monkey
		for humaner.parent != m {
			humaner = humaner.parent
		}
		if m.f1 == humaner.name {
			humanAnswer = m.solveForX(humanAnswer, true, mons)
		} else {
			humanAnswer = m.solveForX(humanAnswer, false, mons)
		}
		// fmt.Println(nonhumaner.name, nonhumaner.getResult(mons))
		m = mons[humaner.name]
	}
	fmt.Println(humanAnswer)
}

func (m *monkey) getResult(mons map[string]*monkey) int {
	if m.f1 != "" {
		//fmt.Println("finding result for", m.name)
		m.result = m.answer(mons[m.f1], mons[m.f2])
	}
	return m.result
}

func (m *monkey) solveForX(answer int, f1human bool, mons map[string]*monkey) int {
	switch m.operator {
	case "+":
		if f1human {
			return answer - mons[m.f2].getResult(mons)
		} else {
			return answer - mons[m.f1].getResult(mons)
		}
	case "-":
		if f1human {
			return answer + mons[m.f2].getResult(mons)
		} else {
			return mons[m.f1].getResult(mons) - answer
		}
	case "/":
		if f1human {
			return answer * mons[m.f2].getResult(mons)
		} else {
			return mons[m.f1].getResult(mons) / answer
		}
	case "*":
		if f1human {
			return answer / mons[m.f2].getResult(mons)
		} else {
			return answer / mons[m.f1].getResult(mons)
		}
	case "=":
		if f1human {
			return mons[m.f2].getResult(mons)
		} else {
			return mons[m.f1].getResult(mons)
		}
	}
	return 0
}
