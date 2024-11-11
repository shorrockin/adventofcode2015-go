package day14

import (
	"adventofcode2015/pkg/utils"
	"math"
	"strings"
)

type Reindeer struct {
	speed int
	time  int
	rest  int
}

func PartOne(path string, totalTime int) int {
	reindeers := parse(path)
	distances := utils.Map(reindeers, func(reindeer Reindeer) int {
		return distance(reindeer, totalTime)
	})
	return utils.Max(distances)
}

// not efficient, but it'll work given the small data size and can
// use all the same methods from part one
func PartTwo(path string, totalTime int) int {
	reindeers := parse(path)
	scores := make(map[Reindeer]int)

	for time := 1; time < totalTime; time++ {
		distances := utils.Map(reindeers, func(reindeer Reindeer) int {
			return distance(reindeer, time)
		})
		var bestReindeers []Reindeer
		bestDistance := 0
		for index, distance := range distances {
			if distance > bestDistance {
				bestDistance = distance
				bestReindeers = []Reindeer{reindeers[index]}
			} else if distance == bestDistance {
				bestReindeers = append(bestReindeers, reindeers[index])
			}
		}

		for _, reindeer := range bestReindeers {
			scores[reindeer] += 1
		}
	}

	return utils.MaxMapValue(scores)
}

func parse(path string) []Reindeer {
	return utils.Map(utils.MustReadInput(path), func(line string) Reindeer {
		parts := strings.Split(line, " ")
		return Reindeer{
			speed: utils.MustAtoi(parts[3]),
			time:  utils.MustAtoi(parts[6]),
			rest:  utils.MustAtoi(parts[13]),
		}
	})
}

func distance(reindeer Reindeer, totalTime int) int {
	cycleTime := reindeer.time + reindeer.rest
	cycleDistance := reindeer.speed * reindeer.time
	cycles := int(math.Floor(float64(totalTime) / float64(cycleTime)))
	timeRemaining := totalTime - (cycles * cycleTime)

	return (cycleDistance * cycles) + (int(math.Min(float64(reindeer.time), float64(timeRemaining))) * reindeer.speed)
}
