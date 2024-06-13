package day19

import (
	"advent/util"
	"strconv"
	"strings"
)

type blueprint struct {
	oreRobot      cost
	clayRobot     cost
	obsidianRobot cost
	geodeRobot    cost
	maxOre        int
	maxClay       int
	maxObsidian   int
}

type cost struct {
	ore      int
	clay     int
	obsidian int
}

func parseBlueprints(path string) []blueprint {
	data := util.GetFileLines(path)

	blueprints := []blueprint{}
	for _, line := range data {
		b := blueprint{}
		blue := strings.Split(line, ": ")[1]
		parts := strings.Split(blue, ". ")

		// ore robot
		orv, _ := strconv.Atoi(parts[0][21:22])
		b.oreRobot = cost{ore: orv}

		// clay robot
		crv, _ := strconv.Atoi(parts[1][22:23])
		b.clayRobot = cost{ore: crv}

		// obsidian robot
		obov := 0
		obcv := 0
		if len(parts[2]) == 42 { // one digit clay value
			obov, _ = strconv.Atoi(parts[2][26:27])
			obcv, _ = strconv.Atoi(parts[2][36:37])
		} else { // two digit clay value
			obov, _ = strconv.Atoi(parts[2][26:27])
			obcv, _ = strconv.Atoi(parts[2][36:38])
		}
		b.obsidianRobot = cost{ore: obov, clay: obcv}

		// geode robot
		gov := 0
		gobv := 0
		if len(parts[3]) == 44 { // one digit clay value
			gov, _ = strconv.Atoi(parts[3][23:24])
			gobv, _ = strconv.Atoi(parts[3][33:34])
		} else { // two digit clay value
			gov, _ = strconv.Atoi(parts[3][23:24])
			gobv, _ = strconv.Atoi(parts[3][33:35])
		}
		b.geodeRobot = cost{ore: gov, obsidian: gobv}

		mo, mc, mob := getMaxRobots(b)
		b.maxOre = mo
		b.maxClay = mc
		b.maxObsidian = mob

		blueprints = append(blueprints, b)
	}

	return blueprints
}

func getMaxRobots(b blueprint) (int, int, int) {
	maxObsidian := b.geodeRobot.obsidian
	maxClay := b.obsidianRobot.clay
	// maxOre := b.geodeRobot.ore + 2
	// maxOre = util.Min(4, maxOre)
	// maxObsidian := 2
	// maxClay := 4
	// maxOre := 1
	maxOre := util.Max(util.Max(b.clayRobot.ore, b.obsidianRobot.ore), b.geodeRobot.ore)
	// maxOre = util.Max((maxOre - 1), 1)
	return maxOre, maxClay, maxObsidian
}
