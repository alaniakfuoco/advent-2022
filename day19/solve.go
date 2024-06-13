package day19

import (
	"advent/util"
	"fmt"
)

type resouces struct {
	oreRobots      int
	clayRobots     int
	obsidianRobots int
	geodeRobots    int
	ore            int
	clay           int
	obsidian       int
	geode          int
}

type choice struct {
	res     resouces
	time    int
	choices []*choice
}

func Solve(path string) {
	blueprints := parseBlueprints(path)

	result := 0
	ch := &choice{
		res:  resouces{oreRobots: 1},
		time: 0,
	}
	for i, b := range blueprints {
		ch.choices = getShortChoices(b, ch, 24)
		fmt.Println(getChoiceCount(ch))
		best := getMaxGeode(ch, 24)
		fmt.Println(best)
		result += (i + 1) * best
	}

	fmt.Println(result)
}

func Solve2(path string) {
	blueprints := parseBlueprints(path)

	result := 0
	ch := &choice{
		res:  resouces{oreRobots: 1},
		time: 0,
	}
	for i, b := range blueprints {
		ch.choices = getShortChoices(b, ch, 32)
		//fmt.Println(getChoiceCount(ch))
		best := getMaxGeode(ch, 32)
		fmt.Println(best)
		result += (i + 1) * best
	}

	fmt.Println(result)
}

func simulateBlueprint(b blueprint) int {
	maxOre, maxClay, maxObsidian := getMaxRobots(b)

	r := resouces{oreRobots: 1}
	fmt.Println(b)

	for i := 0; i < 24; i++ {
		if r.ore >= b.geodeRobot.ore && r.obsidian >= b.geodeRobot.obsidian {
			r.geodeRobots++
			r.ore -= b.geodeRobot.ore
			r.obsidian -= b.geodeRobot.obsidian
		} else if r.obsidianRobots < maxObsidian && r.ore >= b.obsidianRobot.ore && r.clay >= b.obsidianRobot.clay {
			r.obsidianRobots++
			r.ore -= b.obsidianRobot.ore
			r.clay -= b.obsidianRobot.clay
		} else if r.oreRobots < maxOre && r.ore >= b.oreRobot.ore {
			r.oreRobots++
			r.ore -= b.oreRobot.ore
		} else if r.clayRobots < maxClay && r.ore >= b.clayRobot.ore {
			r.clayRobots++
			r.ore -= b.clayRobot.ore
		}
		//fmt.Println("Time:", i, "ORs:", r.oreRobots, "CRs:", r.clayRobots, "OBRs:", r.obsidianRobots, "GRs:", r.geodeRobots)
		fmt.Println("Time:", i, r)

		r.ore += r.oreRobots
		r.clay += r.clayRobots
		r.obsidian += r.obsidianRobots
		r.geode += r.geodeRobots
	}

	fmt.Println(r)
	return r.geode
}

func getChoices(b blueprint, current *choice) []*choice {
	if current.time == 24 {
		return nil
	} else {
		result := []*choice{}
		// if we can build a geode robot skip all other choices
		if current.res.ore >= b.geodeRobot.ore && current.res.obsidian >= b.geodeRobot.obsidian {
			next := &choice{
				res: resouces{
					oreRobots:      current.res.oreRobots,
					clayRobots:     current.res.clayRobots,
					obsidianRobots: current.res.obsidianRobots,
					geodeRobots:    current.res.geodeRobots + 1,
					ore:            (current.res.ore - b.geodeRobot.ore) + current.res.oreRobots,
					clay:           current.res.clay + current.res.clayRobots,
					obsidian:       (current.res.obsidian - b.geodeRobot.obsidian) + current.res.obsidianRobots,
					geode:          current.res.geode + current.res.geodeRobots,
				},
				time: current.time + 1,
			}
			next.choices = getChoices(b, next)
			result = append(result, next)
			return result
		} else { // otherwise add all other build options
			if current.res.obsidianRobots < b.maxObsidian && current.res.ore >= b.obsidianRobot.ore && current.res.clay >= b.obsidianRobot.clay {
				buildObsidian := &choice{
					res: resouces{
						oreRobots:      current.res.oreRobots,
						clayRobots:     current.res.clayRobots,
						obsidianRobots: current.res.obsidianRobots + 1,
						geodeRobots:    current.res.geodeRobots,
						ore:            (current.res.ore - b.obsidianRobot.ore) + current.res.oreRobots,
						clay:           (current.res.clay - b.obsidianRobot.clay) + current.res.clayRobots,
						obsidian:       current.res.obsidian + current.res.obsidianRobots,
						geode:          current.res.geode + current.res.geodeRobots,
					},
					time: current.time + 1,
				}
				buildObsidian.choices = getChoices(b, buildObsidian)
				result = append(result, buildObsidian)
			}
			if current.res.clayRobots < b.maxClay && current.res.ore >= b.clayRobot.ore {
				buildClay := &choice{
					res: resouces{
						oreRobots:      current.res.oreRobots,
						clayRobots:     current.res.clayRobots + 1,
						obsidianRobots: current.res.obsidianRobots,
						geodeRobots:    current.res.geodeRobots,
						ore:            (current.res.ore - b.clayRobot.ore) + current.res.oreRobots,
						clay:           current.res.clay + current.res.clayRobots,
						obsidian:       current.res.obsidian + current.res.obsidianRobots,
						geode:          current.res.geode + current.res.geodeRobots,
					},
					time: current.time + 1,
				}
				buildClay.choices = getChoices(b, buildClay)
				result = append(result, buildClay)
			}
			if current.res.oreRobots < b.maxOre && current.res.ore >= b.oreRobot.ore {
				buildOre := &choice{
					res: resouces{
						oreRobots:      current.res.oreRobots + 1,
						clayRobots:     current.res.clayRobots,
						obsidianRobots: current.res.obsidianRobots,
						geodeRobots:    current.res.geodeRobots,
						ore:            (current.res.ore - b.oreRobot.ore) + current.res.oreRobots,
						clay:           current.res.clay + current.res.clayRobots,
						obsidian:       current.res.obsidian + current.res.obsidianRobots,
						geode:          current.res.geode + current.res.geodeRobots,
					},
					time: current.time + 1,
				}
				buildOre.choices = getChoices(b, buildOre)
				result = append(result, buildOre)
			}
			noBuild := &choice{
				res: resouces{
					oreRobots:      current.res.oreRobots,
					clayRobots:     current.res.clayRobots,
					obsidianRobots: current.res.obsidianRobots,
					geodeRobots:    current.res.geodeRobots,
					ore:            current.res.ore + current.res.oreRobots,
					clay:           current.res.clay + current.res.clayRobots,
					obsidian:       current.res.obsidian + current.res.obsidianRobots,
					geode:          current.res.geode + current.res.geodeRobots,
				},
				time: current.time + 1,
			}
			noBuild.choices = getChoices(b, noBuild)
			result = append(result, noBuild)
			return result
		}
	}
}

func getShortChoices(b blueprint, current *choice, maxTime int) []*choice {
	if current.time == maxTime {
		return nil
	}
	choices := []*choice{}
	if current.res.obsidianRobots > 0 { // add choice for geode robo
		var orewait int
		var obswait int
		if current.res.ore < b.geodeRobot.ore {
			orewait = (b.geodeRobot.ore - current.res.ore) / current.res.oreRobots
			if orewait > -1 && (b.geodeRobot.ore-current.res.ore)%current.res.oreRobots != 0 {
				orewait++
			}
		}
		if current.res.obsidian < b.geodeRobot.obsidian {
			obswait = (b.geodeRobot.obsidian - current.res.obsidian) / current.res.obsidianRobots
			if obswait > -1 && (b.geodeRobot.obsidian-current.res.obsidian)%current.res.obsidianRobots != 0 {
				obswait++
			}
		}
		wait := util.Max(orewait, obswait) + 1
		wait = util.Max(wait, 1)
		if current.time+wait <= maxTime {
			gc := &choice{
				res: resouces{
					oreRobots:      current.res.oreRobots,
					clayRobots:     current.res.clayRobots,
					obsidianRobots: current.res.obsidianRobots,
					geodeRobots:    current.res.geodeRobots + 1,
					ore:            current.res.ore + (current.res.oreRobots * wait),
					clay:           current.res.clay + (current.res.clayRobots * wait),
					obsidian:       current.res.obsidian + (current.res.obsidianRobots * wait),
					geode:          current.res.geode + (current.res.geodeRobots * wait),
				},
				time: current.time + wait,
			}
			gc.res.ore = gc.res.ore - b.geodeRobot.ore
			gc.res.obsidian = gc.res.obsidian - b.geodeRobot.obsidian
			gc.choices = getShortChoices(b, gc, maxTime)
			choices = append(choices, gc)
		}
	}
	if current.res.clayRobots > 0 { // add choice for obsidian robo
		var orewait int
		var claywait int
		if current.res.ore < b.obsidianRobot.ore {
			orewait = (b.obsidianRobot.ore - current.res.ore) / current.res.oreRobots
			if orewait > -1 && (b.obsidianRobot.ore-current.res.ore)%current.res.oreRobots != 0 {
				orewait++
			}
		}
		if current.res.clay < b.obsidianRobot.clay {
			claywait = (b.obsidianRobot.clay - current.res.clay) / current.res.clayRobots
			if claywait > -1 && (b.obsidianRobot.clay-current.res.clay)%current.res.clayRobots != 0 {
				claywait++
			}
		}
		wait := util.Max(orewait, claywait) + 1
		wait = util.Max(wait, 1)
		if current.time+wait <= maxTime && current.res.obsidianRobots < b.maxObsidian {
			obc := &choice{
				res: resouces{
					oreRobots:      current.res.oreRobots,
					clayRobots:     current.res.clayRobots,
					obsidianRobots: current.res.obsidianRobots + 1,
					geodeRobots:    current.res.geodeRobots,
					ore:            current.res.ore + (current.res.oreRobots * wait),
					clay:           current.res.clay + (current.res.clayRobots * wait),
					obsidian:       current.res.obsidian + (current.res.obsidianRobots * wait),
					geode:          current.res.geode + (current.res.geodeRobots * wait),
				},
				time: current.time + wait,
			}
			obc.res.ore = obc.res.ore - b.obsidianRobot.ore
			obc.res.clay = obc.res.clay - b.obsidianRobot.clay
			obc.choices = getShortChoices(b, obc, maxTime)
			choices = append(choices, obc)
		}
	}
	// add choice for clay robot
	var clayRwait int
	if current.res.ore < b.clayRobot.ore {
		clayRwait = (b.clayRobot.ore - current.res.ore) / current.res.oreRobots
		if clayRwait > -1 && (b.clayRobot.ore-current.res.ore)%current.res.oreRobots != 0 {
			clayRwait++
		}
	}
	clayRwait++
	clayRwait = util.Max(clayRwait, 1)
	if current.time+clayRwait <= maxTime && current.res.clayRobots < b.maxClay {
		cc := &choice{
			res: resouces{
				oreRobots:      current.res.oreRobots,
				clayRobots:     current.res.clayRobots + 1,
				obsidianRobots: current.res.obsidianRobots,
				geodeRobots:    current.res.geodeRobots,
				ore:            current.res.ore + (current.res.oreRobots * clayRwait),
				clay:           current.res.clay + (current.res.clayRobots * clayRwait),
				obsidian:       current.res.obsidian + (current.res.obsidianRobots * clayRwait),
				geode:          current.res.geode + (current.res.geodeRobots * clayRwait),
			},
			time: current.time + clayRwait,
		}
		cc.res.ore = cc.res.ore - b.clayRobot.ore
		cc.choices = getShortChoices(b, cc, maxTime)
		choices = append(choices, cc)
	}
	// add choice for ore robot
	var oreRwait int
	if current.res.ore < b.oreRobot.ore {
		oreRwait = (b.oreRobot.ore - current.res.ore) / current.res.oreRobots
		if oreRwait > -1 && (b.oreRobot.ore-current.res.ore)%current.res.oreRobots != 0 {
			oreRwait++
		}
	}
	oreRwait++
	oreRwait = util.Max(oreRwait, 1)
	if current.time+oreRwait <= maxTime && current.res.oreRobots < b.maxOre {
		oc := &choice{
			res: resouces{
				oreRobots:      current.res.oreRobots + 1,
				clayRobots:     current.res.clayRobots,
				obsidianRobots: current.res.obsidianRobots,
				geodeRobots:    current.res.geodeRobots,
				ore:            current.res.ore + (current.res.oreRobots * oreRwait),
				clay:           current.res.clay + (current.res.clayRobots * oreRwait),
				obsidian:       current.res.obsidian + (current.res.obsidianRobots * oreRwait),
				geode:          current.res.geode + (current.res.geodeRobots * oreRwait),
			},
			time: current.time + oreRwait,
		}
		oc.res.ore = oc.res.ore - b.oreRobot.ore
		oc.choices = getShortChoices(b, oc, maxTime)
		choices = append(choices, oc)
	}

	return choices
}

func getMaxGeode(c *choice, maxTime int) int {
	if c.choices == nil || len(c.choices) == 0 {
		g := c.res.geode
		if c.time < maxTime {
			g += c.res.geodeRobots * (maxTime - c.time)
			// fmt.Println(c.time, c.res.geode, c.res.geodeRobots, g, c)
		}
		// if g == 9 {
		// 	fmt.Println(c.time, c.res.geode, c.res.geodeRobots, c)
		// }
		return g
	} else {
		var mg int
		for _, ch := range c.choices {
			nmg := getMaxGeode(ch, maxTime)
			if nmg > mg {
				mg = nmg
			}
		}
		return mg
	}
}

func getChoiceCount(c *choice) int {
	if c.choices == nil || len(c.choices) == 0 {
		return 1
	} else {
		var count int
		for _, ch := range c.choices {
			count += getChoiceCount(ch)
		}
		return count
	}
}
