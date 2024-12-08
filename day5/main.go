package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Type   string // bigger smaller
	Number int
}

type ParseResult struct {
	Rules       map[int][]Rule
	PageUpdates [][]int
}

func main() {
	input := getInput()
	validUpdates := findValidPageUpdates(input)
	middleNumSum := getMiddleNumSum(validUpdates)

	fmt.Printf("middle num sum: %d", middleNumSum)
}

func getMiddleNumSum(validUpdates [][]int) int {
	sum := 0

	for _, u := range validUpdates {
		if len(u)%2 == 0 {
			log.Fatal("no middle number found")
		}

		sum += u[(len(u) / 2)]
	}

	return sum
}

func findValidPageUpdates(result ParseResult) [][]int {
	validUpdates := make([][]int, 0, len(result.PageUpdates))

	for _, update := range result.PageUpdates {
		if isPageUpdateValid(update, result.Rules) {
			validUpdates = append(validUpdates, update)
		}
	}

	return validUpdates
}

func isPageUpdateValid(update []int, rules map[int][]Rule) bool {
	beforeNums := make(map[int]struct{}, len(update))

	for _, num := range update {
		numRules, ok := rules[num]
		if !ok {
			beforeNums[num] = struct{}{}
			continue
		}

		for _, numRule := range numRules {
			if numRule.Type == "bigger" {
				if !slices.Contains(update, numRule.Number) {
					continue
				}

				if _, ok := beforeNums[numRule.Number]; !ok {
					return false
				}
			}
		}

		beforeNums[num] = struct{}{}
	}

	return true
}

func getInput() ParseResult {
	filePath := "day5/input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules := make(map[int][]Rule)
	pageUpdates := make([][]int, 0)
	isRule := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			isRule = false
			continue
		}

		if isRule {
			split := strings.Split(line, "|")
			if len(split) != 2 {
				// Faulty line
				continue
			}

			left, _ := strconv.Atoi(split[0])
			right, _ := strconv.Atoi(split[1])

			if _, ok := rules[right]; !ok {
				rules[right] = make([]Rule, 0)
			}

			rules[right] = append(rules[right], Rule{
				Type:   "bigger",
				Number: left,
			})
			continue
		}

		split := strings.Split(line, ",")
		nums := make([]int, len(split))

		for i, numstr := range split {
			num, _ := strconv.Atoi(numstr)
			nums[i] = num
		}
		pageUpdates = append(pageUpdates, nums)
	}

	return ParseResult{
		Rules:       rules,
		PageUpdates: pageUpdates,
	}
}
