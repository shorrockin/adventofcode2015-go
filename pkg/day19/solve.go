package day19

import (
	"adventofcode2015/pkg/assert"
	"adventofcode2015/pkg/utils"
	"fmt"
	"sort"
	"strings"
)

type Input struct {
	rules    []Rule
	medicine string
}

type Rule struct {
	target      string
	replacement string
}

func PartOne(path string) int {
	input := parse(path)
	combinations := utils.NewSet[string]()

	for _, rule := range input.rules {
		for _, index := range utils.Indexes(input.medicine, rule.target) {
			combinations.Add(replace(input.medicine, index, len(string(rule.target)), rule.replacement))
		}
	}

	return len(combinations)
}

func PartTwo(path string) int {
	input := parse(path)
	fmt.Printf("Input: %+v\n", input)
	combinations := utils.NewSet[string]()
	combinations.Add(input.medicine)
	iterations := 0

	for {
		next := utils.NewSet[string]()

		for current := range combinations {
			for _, rule := range input.rules {
				// because rules are sorted by length, we only try the first
				// one as a greedy algo then continue. this works by chance,
				// but certainly won't for all theoretical combinations of
				// input
				index := strings.Index(current, rule.replacement)
				if index != -1 {
					current := replace(current, index, len(string(rule.replacement)), rule.target)
					next.Add(current)
					break
				}
			}
		}

		if combinations.Contains("e") {
			break
		}

		iterations++
		combinations = next
	}

	return iterations
}

func replace(value string, from int, length int, with string) string {
	start := value[:from]
	end := value[from+length:]
	return start + with + end
}

func parse(path string) Input {
	lines := utils.MustReadInput(path)
	input := Input{rules: make([]Rule, 0, 1)}

	for _, line := range lines {
		if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			assert.True(2 == len(parts), "expected parsed rule to have two parts")
			input.rules = append(input.rules, Rule{parts[0], parts[1]})
		} else if len(line) > 0 {
			input.medicine = line
		}
	}

	sort.Slice(input.rules, func(l int, r int) bool {
		return len(input.rules[l].replacement) > len(input.rules[r].replacement)
	})

	return input
}
